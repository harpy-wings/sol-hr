package solderManager

import "time"

type AlocateBranchRequest struct {
	SolderUid      string     `json:"solder_uid"`
	BranchID       int64      `json:"branch_id"`
	IsPrimary      bool       `json:"is_primary"`
	StartDate      time.Time  `json:"start_date"`
	DueDate        *time.Time `json:"due_date"`
	NetDurationDay int64      `json:"net_duration_day"`
	IsNational     bool       `json:"is_national"`
}
