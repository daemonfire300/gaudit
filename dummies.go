// Copyright (c) Author Julius Foitzik. 2022.
// Licensed under MIT License.

package gaudit

import (
	"context"
	"time"
)

type DummyStorageBackend struct {
	M map[string]AuditTrail
}
type DummyLogger struct{}

var _ StorageBackend = (*DummyStorageBackend)(nil)
var _ Logger = (*DummyLogger)(nil)

func (s *DummyStorageBackend) GetByUser(ctx context.Context, user string) (AuditTrail, error) {
	if trail, found := s.M[user]; found {
		return trail, nil
	}
	return nil, ErrNotFound
}

func (s *DummyStorageBackend) GetByUserForDay(ctx context.Context, user string, day time.Time) (AuditTrail, error) {
	lowerBound := day.Truncate(24 * time.Hour)
	upperBound := lowerBound.Add(24 * time.Hour)
	if trail, found := s.M[user]; found {
		out := make(AuditTrail, 0)
		for _, e := range trail {
			if e.CreatedAt.After(lowerBound) && e.CreatedAt.Before(upperBound) {
				out = append(out, e)
			}
		}
		return out, nil
	}
	return nil, ErrNotFound
}

func (s *DummyStorageBackend) Put(ctx context.Context, entry LogEntry) error {
	existingTrail, found := s.M[entry.User]
	if !found {
		existingTrail = make(AuditTrail, 0)
	}
	existingTrail = append(existingTrail, entry)
	s.M[entry.User] = existingTrail
	return nil
}

func (l *DummyLogger) Debug(msg string) {
	//TODO implement me
	panic("implement me")
}

func (l *DummyLogger) Info(msg string) {
	//TODO implement me
	panic("implement me")
}

func (l *DummyLogger) Warn(msg string) {
	//TODO implement me
	panic("implement me")
}

func (l *DummyLogger) Error(msg string) {
	//TODO implement me
	panic("implement me")
}

func (l *DummyLogger) Debugf(msgfmt string, args ...interface{}) {
	//TODO implement me
	panic("implement me")
}

func (l *DummyLogger) Infof(msgfmt string, args ...interface{}) {
	//TODO implement me
	panic("implement me")
}

func (l *DummyLogger) Warnf(msgfmt string, args ...interface{}) {
	//TODO implement me
	panic("implement me")
}

func (l *DummyLogger) Errorf(msgfmt string, args ...interface{}) {
	//TODO implement me
	panic("implement me")
}
