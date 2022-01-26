// Copyright (c) Author Julius Foitzik. 2022.
// Licensed under MIT License.

package gaudit

import "time"

type AuditTrail []LogEntry

type LogEntry struct {
	CreatedAt time.Time
	User      string
	Action    string
	Meta      map[string]string
}
