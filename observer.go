// Copyright (c) Author Julius Foitzik. 2022.
// Licensed under MIT License.

package gaudit

import (
	"context"
	"fmt"
)

type Observer struct {
	store  StorageBackend
	logger Logger
}

func NewObserver(store StorageBackend, logger Logger) (*Observer, error) {
	if store == nil {
		return nil, fmt.Errorf("storage backend cannot be nil")
	}
	if logger == nil {
		return nil, fmt.Errorf("logger cannot be nil")
	}
	return &Observer{
		store:  store,
		logger: logger,
	}, nil
}

func (ob *Observer) Observe(ctx context.Context, user, action string, meta map[string]string) {
	_ = ob.observe(ctx, &LogEntry{
		User:   user,
		Action: action,
		Meta:   meta,
	})
}

func (ob *Observer) ObserveAndReturn(ctx context.Context, user, action string, meta map[string]string) (LogEntry, error) {
	entry := &LogEntry{
		User:   user,
		Action: action,
		Meta:   meta,
	}
	err := ob.observe(ctx, entry)
	return *entry, err
}

func (ob *Observer) observe(ctx context.Context, entry *LogEntry) error {
	err := ob.store.Put(ctx, *entry)
	if err != nil {
		ob.logger.Errorf("could not store observed audittrail: %v", err)
	}
	return err
}
