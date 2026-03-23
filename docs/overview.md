# Replaygate Overview

## Product thesis

`replaygate` exists to narrow the gap between “we saw the incident” and “we can
reproduce the incident locally.”

The job to be done is:

> Take one real backend failure and turn it into a deterministic replay artifact
> an engineer can run and compare quickly.

## Target user

Primary user:

- Go or backend engineer debugging HTTP service failures

Common scenarios:

- flaky handlers
- production-only timeouts
- dependency drift between staging and production
- difficult regressions involving third-party APIs or database state

## Current prototype

The current version is intentionally narrow:

- incidents are provided as local JSON
- replay is currently inspection and diffing, not real service execution
- supported fixture types are incoming request, outgoing HTTP, and PostgreSQL

That narrow scope is useful because it validates the core artifact model before
introducing heavy runtime orchestration.

## Product boundaries

In scope for the near term:

- incident bundling
- fixture-driven replay
- deterministic diffs
- local debugging workflows

Out of scope for the near term:

- complete observability platform
- traffic capture for every transport
- broad language support
- full staging replacement
