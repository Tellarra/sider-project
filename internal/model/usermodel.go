package model

import "time"

// User represents the structure of a user data.
type User struct {
	ID                    string      `json:"_id"`
	Profile               UserProfile `json:"profile"`
	Emails                []UserEmail `json:"emails"`
	ActiveOrganisationID  string      `json:"activeOrganisationId"`
	LinkedOrganisationIds *[]string   `json:"linkedOrganisationIds"` // Assuming this can be null
}

// UserProfile represents the profile information of a user.
type UserProfile struct {
	FirstName      string       `json:"FirstName"`
	LastName       string       `json:"LastName"`
	PhoneNumber    string       `json:"PhoneNumber"`
	DateOfBirth    time.Time    `json:"DateOfBirth"`
	NationalNumber string       `json:"NationalNumber"`
	CountryCode    string       `json:"CountryCode"`
	Settings       UserSettings `json:"Settings"`
	Gender         string       `json:"Gender"`
	Formations     *[]string    `json:"Formations"`  // Assuming this can be null
	Hobbies        *[]string    `json:"Hobbies"`     // Assuming this can be null
	Experiences    *[]string    `json:"Experiences"` // Assuming this can be null
}

// UserSettings represents user-specific settings.
type UserSettings struct {
	Locale       string `json:"Locale"`
	ActiveCityID string `json:"ActiveCityID"`
}

// UserEmail represents an email object within a user.
type UserEmail struct {
	Address  string `json:"Address"`
	Verified bool   `json:"Verified"`
}

type UserToFeed struct {
	ID                    string      `json:"id"`
	Profile               UserProfile `json:"profile"`
	Emails                []UserEmail `json:"emails"`
	ActiveOrganisationID  string      `json:"activeOrganisationId"`
	LinkedOrganisationIds *[]string   `json:"linkedOrganisationIds"` // Assuming this can be null
}

func (u User) ToUserToFeed() UserToFeed {
	return UserToFeed(u)
}
