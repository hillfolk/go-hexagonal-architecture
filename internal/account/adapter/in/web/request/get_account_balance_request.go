package request

type GetAccountBalanceRequest struct {
	AccountId int64 `json:"account_id" binding:"required"`
}
