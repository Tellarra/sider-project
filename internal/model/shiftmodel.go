package model

import "time"

// Shift represents a shift from your JSON structure
type Shift struct {
	ID                string    `json:"_id"`
	AvailableSiderIds []string  `json:"availableSiderIds"`
	HiredSiderIds     []string  `json:"hiredSiderIds"`
	Time              ShiftTime `json:"time"`
	Type              string    `json:"type"`
	Slots             int       `json:"slots"`
	TaskID            string    `json:"taskId"`
	Break             int       `json:"break"`
	Location          string    `json:"location"`
	CreatedAt         time.Time `json:"createdAt"`
	UpdatedAt         time.Time `json:"updatedAt"`
	Status            string    `json:"status"`
}

type ShiftToFeed struct {
	ID                string    `json:"id"`
	AvailableSiderIds []string  `json:"availableSiderIds"`
	HiredSiderIds     []string  `json:"hiredSiderIds"`
	Time              ShiftTime `json:"time"`
	Type              string    `json:"type"`
	Slots             int       `json:"slots"`
	TaskID            string    `json:"taskId"`
	Break             int       `json:"break"`
	Location          string    `json:"location"`
	CreatedAt         time.Time `json:"createdAt"`
	UpdatedAt         time.Time `json:"updatedAt"`
	Status            string    `json:"status"`
}

// ShiftTime represents the time structure within a Shift
type ShiftTime struct {
	StartDate time.Time `json:"startDate"`
	EndDate   time.Time `json:"endDate"`
}

func (s Shift) ToShiftToFeed() ShiftToFeed {
	return ShiftToFeed(s)
}
