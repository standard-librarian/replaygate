# Incident Spec

## Current shape

The current starter incident bundle is JSON with this high-level structure:

```json
{
  "schema_version": 1,
  "kind": "replaygate.bundle",
  "incident": {
    "request": {},
    "expected_response": {},
    "actual_response": {},
    "outgoing_http": [],
    "postgres": [],
    "metadata": {}
  }
}
```

## Field intent

- `schema_version`: allows future migration without ambiguity
- `kind`: simple type marker for tooling
- `incident.request`: incoming HTTP request data
- `incident.expected_response`: what the service should have returned
- `incident.actual_response`: what it returned during the failure
- `incident.outgoing_http`: mocked downstream HTTP interactions
- `incident.postgres`: mocked PostgreSQL queries and results
- `incident.metadata`: service or environment context

## Constraints for next iteration

- keep the bundle easy to read and edit by hand
- preserve the deterministic replay story over raw completeness
- do not couple the file format too tightly to one service framework

## Likely future additions

- fixture timestamps and ordering metadata
- replay environment variables
- richer response diff sections
- fixture references for larger payloads
