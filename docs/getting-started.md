# Getting Started

`go-secrets` provides a small abstraction for packages that want to depend on a secret store without coupling themselves to a single backend.

## Included stores

- `secrets.MemoryStore` for in-memory use and tests.
- `secrets.EnvStore` for read-only access to environment variables.
- `secrets.VaultStore` as a placeholder for future external secret manager integrations.

## Error behavior

The package exposes explicit sentinel errors for not-found, read-only, and not-implemented scenarios so callers can branch on them with `errors.Is`.
