# Agent Handoff

## Mission

Continue `replaygate` as a focused incident replay tool for backend debugging.
Do not turn it into another generic Go web framework or observability suite.

## What already exists

- CLI command to ingest an incident into a bundle
- CLI command to summarize and diff a bundled incident against a candidate
- example failing and fixed incidents
- end-to-end smoke coverage

## What matters most next

- refactor the CLI into packages
- make the incident schema explicit and validated
- add local replay execution against a Go handler
- improve fixture modeling and diff readability

## Constraints

- stay Go-first for now
- keep the tool CLI-first
- prioritize deterministic local replay over broad capture coverage
- avoid adding too many integrations at once
- optimize for a crisp demo and a 5-minute first success

## Suggested work order

1. Split `main.go` into packages and keep CLI behavior stable.
2. Add schema validation and unit tests.
3. Build local replay against a sample handler using HTTP and PostgreSQL
   fixtures.
4. Improve diff formatting and example coverage.
5. Refresh README around the new replay path.

## Done means

- current quickstart still works or is replaced with a better handler-based demo
- tests cover schema validation and replay internals
- docs explain the execution model without requiring code archaeology
