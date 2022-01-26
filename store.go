// Copyright (c) Author Julius Foitzik. 2022.
// Licensed under MIT License.

package gaudit

import (
	"context"
	"time"
)

type StorageBackend interface {
	// Put should implement functionality to store a log entry, whatever that constitutes for your special case.
	Put(ctx context.Context, entry LogEntry) error

	// GetByUser should return the AuditTrail for that specific user or ErrNotFound if none can be found.
	GetByUser(ctx context.Context, user string) (AuditTrail, error)

	// GetByUserForDay should return the AuditTrail for that specific user for that
	// specific day given as time.Time object or ErrNotFound if none can be found.
	GetByUserForDay(ctx context.Context, user string, day time.Time) (AuditTrail, error)
}
