package model

import "time"

// Slot represents the data structure for a slot.
type Slot struct {
	ID             string    `json:"_id"`
	TaskID         string    `json:"taskId"`
	ShiftID        string    `json:"shiftId"`
	OrganisationID string    `json:"organisationId"`
	IsOverbooking  bool      `json:"isOverbooking"`
	Status         string    `json:"status"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
	SiderStatus    string    `json:"siderStatus"`
}

type SlotToFeed struct {
	ID             string    `json:"id"`
	TaskID         string    `json:"taskId"`
	ShiftID        string    `json:"shiftId"`
	OrganisationID string    `json:"organisationId"`
	IsOverbooking  bool      `json:"isOverbooking"`
	Status         string    `json:"status"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
	SiderStatus    string    `json:"siderStatus"`
}

func (s Slot) ToSlotToFeed() SlotToFeed {
	return SlotToFeed(s)
}
