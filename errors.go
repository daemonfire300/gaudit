// Copyright (c) Author Julius Foitzik. 2022.
// Licensed under MIT License.

package gaudit

import "errors"

var (
	ErrNotFound    = errors.New("no results")
	ErrSetupFailed = errors.New("setup failed")
)
