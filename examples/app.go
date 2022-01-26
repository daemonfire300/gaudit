// Copyright (c) Author Julius Foitzik. 2022.
// Licensed under MIT License.

package main

import (
	"context"
	"fmt"
	"log"

	"github.com/daemonfire300/gaudit"
)

func main() {
	if err := SomeRESTCall(context.Background()); err != nil {
		log.Fatal(err)
	}
}

func SomeRESTCall(ctx context.Context) error {
	audit, err := gaudit.NewObserver(&gaudit.DummyStorageBackend{make(map[string]gaudit.AuditTrail)}, &gaudit.DummyLogger{})
	if err != nil {
		panic(err)
	}
	// Return copy of observation
	observation, _ := audit.ObserveAndReturn(ctx, "peter", "didSomeRestCall", map[string]string{"metadataKey": "goobedigoog"})
	fmt.Println("observation", observation)
	// Just observe do not return any data
	audit.Observe(ctx, "peter", "didSomeRestCall", map[string]string{"metadataKey": "goobedigoog"})
	return nil
}
