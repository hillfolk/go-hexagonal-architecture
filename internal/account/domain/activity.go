package domain

import "time"

type ActivityId *int64

func NewActivityId(id int64) ActivityId {
	return &id
}

type Activity struct {
	Id              ActivityId
	OwnerAccountId  AccountId
	SourceAccountId AccountId
	TargetAccountId AccountId
	Timestamp       time.Time
	ActiveMoney     Money
}

func NewActivity(
	ownerAccountId AccountId,
	sourceAccountId AccountId,
	targetAccountId AccountId,
	timestamp time.Time,
	money Money,
) Activity {
	return Activity{
		Id:              nil,
		OwnerAccountId:  ownerAccountId,
		SourceAccountId: sourceAccountId,
		TargetAccountId: targetAccountId,
		Timestamp:       time.Now(),
		ActiveMoney:     money}

}
