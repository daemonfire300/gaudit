// Copyright (c) Author Julius Foitzik. 2022.
// Licensed under MIT License.

package gaudit

type Logger interface {
	Debug(msg string)
	Info(msg string)
	Warn(msg string)
	Error(msg string)

	Debugf(msgfmt string, args ...interface{})
	Infof(msgfmt string, args ...interface{})
	Warnf(msgfmt string, args ...interface{})
	Errorf(msgfmt string, args ...interface{})
}
