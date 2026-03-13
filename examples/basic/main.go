package main

import (
	"fmt"

	"github.com/mirkobrombin/go-secrets/pkg/secrets"
)

func main() {
	store := secrets.NewMemoryStore()
	if err := store.Set("api-key", []byte("secret")); err != nil {
		panic(err)
	}

	value, err := store.Get("api-key")
	if err != nil {
		panic(err)
	}

	fmt.Printf("secret length: %d\n", len(value))
}
