package domain

import "time"

type ApprovalRecord struct {
	ID         int64
	ResourceID int64
	Status     ResourceStatus
	Notes      string
	CreatedAt  time.Time
}
