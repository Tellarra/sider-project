package model

import "time"

type Query struct {
	UserId       string `form:"userId"`
	FilterStatus bool   `form:"filterStatus"`
}

type Task struct {
	ID                   string          `json:"_id"`
	LogAsID              string          `json:"logAsId"`
	CreatorID            string          `json:"creatorId"`
	OrganisationID       string          `json:"organisationId"`
	ShiftIDs             []string        `json:"shiftIds"`
	CityID               string          `json:"cityId"`
	ManagerID            string          `json:"managerId"`
	RequestedSiderIds    []string        `json:"requestedSiderIds"`
	AssigneeID           string          `json:"assigneeId"`
	Alias                string          `json:"alias"`
	Status               string          `json:"status"`
	Address              Address         `json:"address"`
	LocationOptions      LocationOptions `json:"locationOptions"`
	WorkLegalStatus      string          `json:"workLegalStatus"`
	SelectionStatus      string          `json:"selectionStatus"`
	RequestedSidersOnly  bool            `json:"requestedSidersOnly"`
	IsPreSelection       bool            `json:"isPreSelection"`
	Visible              bool            `json:"visible"`
	UsersAlreadyNotified bool            `json:"usersAlreadyNotified"`
	CompanyNotified      bool            `json:"companyNotified"`
	Type                 string          `json:"type"`
	SubtaskIds           []string        `json:"subtaskIds"`
	Purpose              string          `json:"purpose"`
	DressCode            string          `json:"dressCode"`
	WorkConditions       string          `json:"workConditions"`
	Experiences          string          `json:"experiences"`
	Motive               Motive          `json:"motive"`
	MissionInformation   string          `json:"missionInformation"`
	SideNote             string          `json:"sideNote"`
	PricingID            string          `json:"pricingId"`
	HourlyRate           float64         `json:"hourlyRate"`
	SubmittedAt          time.Time       `json:"submittedAt"`
	LiveAt               time.Time       `json:"liveAt"`
	PostedAt             time.Time       `json:"postedAt"`
	Applicants           []Applicant     `json:"applicants"`
}

type TaskToFeed struct {
	ID                   string          `json:"id"`
	LogAsID              string          `json:"logAsId"`
	CreatorID            string          `json:"creatorId"`
	OrganisationID       string          `json:"organisationId"`
	ShiftIDs             []string        `json:"shiftIds"`
	CityID               string          `json:"cityId"`
	ManagerID            string          `json:"managerId"`
	RequestedSiderIds    []string        `json:"requestedSiderIds"`
	AssigneeID           string          `json:"assigneeId"`
	Alias                string          `json:"alias"`
	Status               string          `json:"status"`
	Address              Address         `json:"address"`
	LocationOptions      LocationOptions `json:"locationOptions"`
	WorkLegalStatus      string          `json:"workLegalStatus"`
	SelectionStatus      string          `json:"selectionStatus"`
	RequestedSidersOnly  bool            `json:"requestedSidersOnly"`
	IsPreSelection       bool            `json:"isPreSelection"`
	Visible              bool            `json:"visible"`
	UsersAlreadyNotified bool            `json:"usersAlreadyNotified"`
	CompanyNotified      bool            `json:"companyNotified"`
	Type                 string          `json:"type"`
	SubtaskIds           []string        `json:"subtaskIds"`
	Purpose              string          `json:"purpose"`
	DressCode            string          `json:"dressCode"`
	WorkConditions       string          `json:"workConditions"`
	Experiences          string          `json:"experiences"`
	Motive               Motive          `json:"motive"`
	MissionInformation   string          `json:"missionInformation"`
	SideNote             string          `json:"sideNote"`
	PricingID            string          `json:"pricingId"`
	HourlyRate           float64         `json:"hourlyRate"`
	SubmittedAt          time.Time       `json:"submittedAt"`
	LiveAt               time.Time       `json:"liveAt"`
	PostedAt             time.Time       `json:"postedAt"`
	Applicants           []Applicant     `json:"applicants"`
}

type Address struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	ZipCode string `json:"zipCode"`
	Country string `json:"country"`
}

type LocationOptions struct {
	Remote    RemoteOptions `json:"remote"`
	Motorized bool          `json:"motorized"`
}

type RemoteOptions struct {
	Available bool `json:"available"`
	Mandatory bool `json:"mandatory"`
}

type Motive struct {
	Reason       string        `json:"reason"`
	Replacements []Replacement `json:"replacements"`
}

type Replacement struct {
	Name          string `json:"name"`
	Position      string `json:"position"`
	Justification string `json:"justification"`
}

type Applicant struct {
	ApplicantId            string          `json:"applicantId"`
	Status                 string          `json:"status"`
	Answers                []string        `json:"answers"`
	StatusHistory          []StatusHistory `json:"statusHistory"`
	When                   time.Time       `json:"when"`
	Notified               bool            `json:"notified"`
	LastStatusNotified     string          `json:"lastStatusNotified"`
	LastStatusNotifiedDate time.Time       `json:"lastStatusNotifiedDate"`
	Reason                 string          `json:"reason"`
}

type StatusHistory struct {
	Status string    `json:"status"`
	Date   time.Time `json:"date"`
}

type DocumentToIndex struct {
	Data []byte
	Name string
}

func (t Task) ToTaskToFeed() TaskToFeed {
	return TaskToFeed(t)
}
