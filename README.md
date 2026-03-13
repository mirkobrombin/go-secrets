# Go Secrets

A compact **secret store abstraction** for Go with in-memory and environment-backed implementations.

## Features

- **Shared Store Contract:** Keep secret access behind a small interface.
- **In-Memory Store:** Useful for tests and ephemeral application state.
- **Environment Store:** Read secrets from environment variables without exposing write operations.
- **Extensible Design:** Reserve room for Vault or other external backends.

## Installation

```bash
go get github.com/mirkobrombin/go-secrets
```

## Quick Start

```go
package main

import (
    "fmt"

    "github.com/mirkobrombin/go-secrets/pkg/secrets"
)

func main() {
    store := secrets.NewMemoryStore()
    if err := store.Set("db-password", []byte("super-secret")); err != nil {
        panic(err)
    }

    value, err := store.Get("db-password")
    if err != nil {
        panic(err)
    }

    fmt.Println(len(value))
}
```

## Documentation

- [Getting Started](docs/getting-started.md)

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
