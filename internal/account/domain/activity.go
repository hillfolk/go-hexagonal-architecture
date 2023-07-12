package domain

import "time"

type ActivityId *int64
type Activity struct {
	Id              ActivityId
	OnwerAccountId  AccountId
	SourceAccountId AccountId
	TargetAccountId AccountId
	Timestamp       time.Time
	ActiveMoney     Money
}

func NewActivity(
	onwerAccountId AccountId,
	sourceAccountId AccountId,
	targetAccountId AccountId,
	timestamp time.Time,
	money Money,
) Activity {
	return Activity{
		Id:              nil,
		OnwerAccountId:  onwerAccountId,
		SourceAccountId: sourceAccountId,
		TargetAccountId: targetAccountId,
		Timestamp:       time.Now(),
		ActiveMoney:     money}

}
