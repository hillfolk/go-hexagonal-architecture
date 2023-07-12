package domain

import (
	"time"
)

type ActiveWindow struct {
	activities []*Activity
}

func (w *ActiveWindow) GetStartTimestamp() time.Time {
	startTimestamp := w.activities[0].Timestamp
	for _, activity := range w.activities {
		if activity.Timestamp.Before(startTimestamp) {
			startTimestamp = activity.Timestamp
		}
	}
	return startTimestamp
}

func (w *ActiveWindow) GetEndTimestamp() time.Time {
	endTimestamp := w.activities[0].Timestamp
	for _, activity := range w.activities {
		if activity.Timestamp.After(endTimestamp) {
			endTimestamp = activity.Timestamp
		}
	}
	return endTimestamp
}

func (w *ActiveWindow) CalculateBalance(accountId AccountId) Money {

	depositBalance := Money{
		Amount: 0,
	}

	withdrawalBalance := Money{
		Amount: 0,
	}
	for _, activity := range w.activities {
		if activity.TargetAccountId == accountId {
			depositBalance = Money{}.Add(depositBalance, activity.ActiveMoney)
		}
		if activity.SourceAccountId == accountId {
			withdrawalBalance = Money{}.Add(withdrawalBalance, activity.ActiveMoney)
		}
	}
	return Money{}.Add(depositBalance, withdrawalBalance)
}

func (w *ActiveWindow) CanWithdraw(amount Money) bool {
	return false
}

func (w *ActiveWindow) AddActivity(activity *Activity) {
	w.activities = append(w.activities, activity)
}

func (w *ActiveWindow) GetActivities() []*Activity {
	return w.activities
}
