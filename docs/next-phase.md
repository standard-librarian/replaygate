# Next Phase

## Goal

Turn the prototype from a pure artifact and diff tool into a minimal local
incident replay system for Go HTTP services.

## Scope for the next implementation phase

### 1. Internal refactor

- split `main.go` into packages for spec, I/O, diffing, and replay
- preserve current CLI commands and example flow
- add targeted unit tests beyond the subprocess smoke test

### 2. Stronger incident contract

- formalize the incident schema in code
- validate required request and response fields
- print useful validation errors when incidents are malformed

### 3. Local replay execution

- add a mode that replays the incoming request against a local handler
- use fixture-backed outgoing HTTP and PostgreSQL responses
- keep the initial execution model simple and deterministic

### 4. Better diff rendering

- group diffs by response, HTTP fixtures, and database fixtures
- make status/body mismatches obvious at a glance
- show a short replay summary before the diff

### 5. Fixture ergonomics

- define clearer fixture types for outgoing HTTP and PostgreSQL
- support a small realistic example service for demo purposes
- document how an engineer should author and update fixtures

## Acceptance criteria

- an engineer can ingest an incident and replay it against a local Go handler
- outgoing HTTP and PostgreSQL behavior can be mocked from the incident bundle
- malformed incidents fail fast with useful messages
- diff output is sectioned and easy to scan
- smoke and unit tests pass locally

## Non-goals

- multi-language service support
- production traffic capture agent
- container orchestration
- hosted debugging workflows
