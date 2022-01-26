// Copyright (c) Author Julius Foitzik. 2022.
// Licensed under MIT License.

package gaudit

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestObserver_Observe(t *testing.T) {
	backend := &DummyStorageBackend{map[string]AuditTrail{}}
	ob, err := NewObserver(backend, &DummyLogger{})
	require.NoError(t, err)
	ctx := context.Background()
	user := "pete"
	action := "Moved an object"
	ob.Observe(ctx, user, action, nil)
	require.Len(t, backend.M, 1)
	require.Len(t, backend.M[user], 1)
	assert.Equal(t, user, backend.M[user][0].User)
	assert.Equal(t, action, backend.M[user][0].Action)
	assert.False(t, backend.M[user][0].CreatedAt.IsZero())
}

func TestObserver_ObserveAndReturn(t *testing.T) {
	backend := &DummyStorageBackend{map[string]AuditTrail{}}
	ob, err := NewObserver(backend, &DummyLogger{})
	require.NoError(t, err)
	ctx := context.Background()
	user := "pete"
	action := "Moved an object"
	entry, err := ob.ObserveAndReturn(ctx, user, action, nil)
	require.NoError(t, err)
	require.Len(t, backend.M, 1)
	require.Len(t, backend.M[user], 1)
	assert.Equal(t, user, entry.User)
	assert.Equal(t, action, entry.Action)
	assert.False(t, entry.CreatedAt.IsZero())
}

func TestNewObserver_NilBackend(t *testing.T) {
	_, err := NewObserver(nil, &DummyLogger{})
	require.Error(t, err)
	require.Truef(t, errors.Is(err, ErrSetupFailed), "err must be or wrap ErrSetupFailed")
}

func TestNewObserver_NilLogger(t *testing.T) {
	_, err := NewObserver(&DummyStorageBackend{}, nil)
	require.Error(t, err)
	require.Truef(t, errors.Is(err, ErrSetupFailed), "err must be or wrap ErrSetupFailed")
}
