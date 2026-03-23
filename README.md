# replaygate

Replay real production bugs locally, deterministically, in minutes.

`replaygate` is a starter incident replay CLI for Go services. It converts one
captured HTTP incident into a stable bundle and shows diffs between the expected
and candidate behavior.

## Who this is for

`replaygate` targets backend and platform engineers who routinely hear some
version of:

- “it only broke in prod”
- “we have logs, but we still can’t reproduce it”
- “staging never shows the real failure”

The goal is to make one production incident portable enough to replay and debug
without rebuilding half of production by hand.

## Why this matters

Most backend debugging workflows stop at logs, traces, or screenshots. Those
are useful for diagnosis, but weak for reproduction. `replaygate` is trying to
turn an incident into a deterministic artifact another engineer can run,
compare, and iterate on locally.

## 3-minute quickstart

```bash
go run . ingest examples/incident.json --output build/incident.bundle.json
go run . replay build/incident.bundle.json --candidate examples/fixed_incident.json
```

## Current prototype status

Today this repo proves three core ideas:

- an incident can be normalized into a stable JSON bundle
- a replay command can summarize the captured context clearly
- candidate behavior can be diffed against the original incident quickly

This is a wedge, not the finished product. There is no live service bootstrapping
yet and no traffic capture agent yet.

## Commands

- `ingest <incident> --output <bundle>`: normalize an incident into a replay
  bundle
- `replay <bundle> [--candidate <incident>]`: inspect the bundle and diff it
  against a second run

## Supported fixture areas in this starter version

- incoming HTTP request
- expected HTTP response
- outgoing HTTP mocks
- PostgreSQL query/result fixtures

See `docs/incident-spec.md` for the current bundle shape and extension goals.

## Project shape

- `main.go`: current CLI, bundle generation, and diff output
- `examples/`: failing and fixed incident fixtures
- `main_test.go`: end-to-end smoke coverage for ingest and replay

## Limitations

- no live service execution yet
- no traffic capture agent yet
- JSON fixtures only

## If you continue this project

Start with `docs/next-phase.md`. It defines the next useful slice:

- split the CLI into packages
- formalize the incident schema
- add service execution against fixtures
- make diffs and fixture handling more ergonomic

`docs/agent-handoff.md` is written so another agent can continue implementation
without re-deriving the product direction.
