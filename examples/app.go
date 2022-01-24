// Copyright (c) Author Julius Foitzik. 2022.
// Licensed under MIT License.

package examples

import (
	"context"
	"fmt"

	"github.com/daemonfire300/gaudit"
)

type store struct{}

type logger struct{}

var _ gaudit.StorageBackend = (*store)(nil)
var _ gaudit.Logger = (*logger)(nil)

func (s store) Put(ctx context.Context, trail gaudit.LogEntry) error {
	//TODO implement me
	panic("implement me")
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
