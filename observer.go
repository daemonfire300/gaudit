// Copyright (c) Author Julius Foitzik. 2022.
// Licensed under MIT License.

package gaudit

import (
	"context"
	"fmt"
	"time"
)

type Observer struct {
	store  StorageBackend
	logger Logger
}

// NewObserver creates a new observer and returns an error if the passed arguments do not match.
func NewObserver(store StorageBackend, logger Logger) (*Observer, error) {
	if store == nil {
		return nil, fmt.Errorf("%w: storage backend cannot be nil", ErrSetupFailed)
	}
	if logger == nil {
		return nil, fmt.Errorf("%w: logger cannot be nil", ErrSetupFailed)
	}
	return &Observer{
		store:  store,
		logger: logger,
	}, nil
}

// Observe simply observes an action by a user and adds metadata if given.
// The date of observation is equal to the time of calling Observe, i.e., time.Now.
func (ob *Observer) Observe(ctx context.Context, user, action string, meta map[string]string) {
	_ = ob.observe(ctx, &LogEntry{
		CreatedAt: time.Now(),
		User:      user,
		Action:    action,
		Meta:      meta,
	})
}

// ObserveAndReturn works exactly as Observe and additionally returns a copy of the created LogEntry and an
// error in case the StorageBackend returned an error.
func (ob *Observer) ObserveAndReturn(ctx context.Context, user, action string, meta map[string]string) (LogEntry, error) {
	entry := &LogEntry{
		CreatedAt: time.Now(),
		User:      user,
		Action:    action,
		Meta:      meta,
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
