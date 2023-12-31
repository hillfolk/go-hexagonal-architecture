package request

type SendMoneyRequest struct {
	SourceAccountId int64 `json:"source_account_id" form:"source_account_id" binding:"required"`
	TargetAccountId int64 `json:"target_account_id" form:"target_account_id" binding:"required"`
	Amount          int64 `json:"amount" form:"amount" binding:"required"`
}
