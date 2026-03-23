# Architecture

## Current architecture

The prototype is a compact single-file CLI that handles:

- argument parsing
- incident loading
- bundle writing
- replay summary output
- candidate diff output

That is acceptable for a proof of concept, but the next phase should move the
core behaviors into packages so real replay execution can be added cleanly.

## Recommended package split

- `cmd/replaygate`: command wiring only
- `internal/spec`: incident and bundle types
- `internal/io`: JSON loading and writing
- `internal/diff`: diff generation and formatting
- `internal/replay`: fixture-driven replay execution

## Design principles

- make incidents portable and inspectable
- keep replay deterministic before making it comprehensive
- prefer a sharp CLI over early framework coupling
- do not overbuild capture before replay becomes credible

## Important implementation notes

- fixture ordering will matter once real replay is added
- diff output should distinguish request, response, and dependency changes
- execution should be opt-in and local-first
