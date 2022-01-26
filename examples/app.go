// Copyright (c) Author Julius Foitzik. 2022.
// Licensed under MIT License.

package examples

import (
	"context"
	"fmt"
	"time"

	"github.com/daemonfire300/gaudit"
)

type store struct {
	m map[string]gaudit.AuditTrail
}
type logger struct{}

var _ gaudit.StorageBackend = (*store)(nil)
var _ gaudit.Logger = (*logger)(nil)

func (s *store) GetByUser(ctx context.Context, user string) (gaudit.AuditTrail, error) {
	if trail, found := s.m[user]; found {
		return trail, nil
	}
	return nil, gaudit.ErrNotFound
}

func (s *store) GetByUserForDay(ctx context.Context, user string, day time.Time) (gaudit.AuditTrail, error) {
	lowerBound := day.Truncate(24 * time.Hour)
	upperBound := lowerBound.Add(24 * time.Hour)
	if trail, found := s.m[user]; found {
		out := make(gaudit.AuditTrail, 0)
		for _, e := range trail {
			if e.CreatedAt.After(lowerBound) && e.CreatedAt.Before(upperBound) {
				out = append(out, e)
			}
		}
		return out, nil
	}
	return nil, gaudit.ErrNotFound
}

func (s *store) Put(ctx context.Context, entry gaudit.LogEntry) error {
	existingTrail, found := s.m[entry.User]
	if !found {
		existingTrail = make(gaudit.AuditTrail, 0)
	}
	existingTrail = append(existingTrail, entry)
	return nil
}

func (l *logger) Debug(msg string) {
	//TODO implement me
	panic("implement me")
}

func (l *logger) Info(msg string) {
	//TODO implement me
	panic("implement me")
}

func (l *logger) Warn(msg string) {
	//TODO implement me
	panic("implement me")
}

func (l *logger) Error(msg string) {
	//TODO implement me
	panic("implement me")
}

func (l *logger) Debugf(msgfmt string, args ...interface{}) {
	//TODO implement me
	panic("implement me")
}

func (l *logger) Infof(msgfmt string, args ...interface{}) {
	//TODO implement me
	panic("implement me")
}

func (l *logger) Warnf(msgfmt string, args ...interface{}) {
	//TODO implement me
	panic("implement me")
}

func (l *logger) Errorf(msgfmt string, args ...interface{}) {
	//TODO implement me
	panic("implement me")
}

func SomeRESTCall(ctx context.Context) error {
	audit, err := gaudit.NewObserver(&store{}, &logger{})
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
