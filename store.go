// Copyright (c) Author Julius Foitzik. 2022.
// Licensed under MIT License.

package gaudit

import "context"

type StorageBackend interface {
	Put(ctx context.Context, trail LogEntry) error
}
