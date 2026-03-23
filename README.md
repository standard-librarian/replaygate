# replaygate

Replay real production bugs locally, deterministically, in minutes.

`replaygate` is a starter incident replay CLI for Go services. It converts one
captured HTTP incident into a stable bundle and shows diffs between the expected
and candidate behavior.

## 3-minute quickstart

```bash
go run . ingest examples/incident.json --output build/incident.bundle.json
go run . replay build/incident.bundle.json --candidate examples/fixed_incident.json
```

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

## Limitations

- no live service execution yet
- no traffic capture agent yet
- JSON fixtures only
