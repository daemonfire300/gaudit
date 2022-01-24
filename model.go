// Copyright (c) Author Julius Foitzik. 2022.
// Licensed under MIT License.

package gaudit

type AuditTrail struct {
}

type LogEntry struct {
	User   string
	Action string
	Meta   map[string]string
}
