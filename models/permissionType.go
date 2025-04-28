package models

type PermissionType int

const (
	PermissionTypeNone PermissionType = iota
	PermissionTypeReadOnly
	PermissionTypeExport
	PermissionTypeReadWrite
	PermissionTypeDelete
)
