package branchManager

type QueryBranchesRequest struct {
	Query   string `json:"query"`
	OrderBy string `json:"order_by"`
	Xid     *int64 `json:"xid"`
}
