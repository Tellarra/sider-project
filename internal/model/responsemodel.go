package model

import "slices"

type Response struct {
	Task TaskResponse `json:"task"`
}

type TaskResponse struct {
	Id     string        `json:"id"`
	Name   string        `json:"name"`
	Orga   OrgaResponse  `json:"organisation"`
	Shift  ShiftResponse `json:"shift"`
	Status string        `json:"status,omitempty"`
}

type OrgaResponse struct {
	Name       string `json:"name"`
	Address    string `json:"address"`
	PictureUrl string `json:"pictureUrl"`
}

type ShiftResponse struct {
	Id        string        `json:"id"`
	StartDate string        `json:"startDate"`
	EndDate   string        `json:"endDate"`
	Slots     SlotsResponse `json:"slots"`
	Status    string        `json:"status"`
}

type SlotsResponse struct {
	Filled int `json:"filled"`
	Total  int `json:"total"`
}

func GetOrga(orgaTaskID string, orga []OrgaToFeed) OrgaResponse {
	for _, o := range orga {
		if o.ID == orgaTaskID {
			return OrgaResponse{
				Name:       o.Name,
				Address:    o.Address,
				PictureUrl: o.LogoUrl,
			}
		}
	}
	return OrgaResponse{}
}

func GetShift(taskId string, shiftsIds []string, shifts []ShiftToFeed) ShiftResponse {
	for _, shift := range shiftsIds {

		index := slices.IndexFunc(shifts, func(s ShiftToFeed) bool {
			return s.ID == shift
		})
		if index != -1 {
			shift := shifts[index]
			return ShiftResponse{
				Id:        shift.ID,
				StartDate: shift.Time.StartDate.String(),
				EndDate:   shift.Time.EndDate.String(),
				Slots: SlotsResponse{
					Filled: shift.Slots,
					Total:  shift.Slots,
				},
				Status: shift.Status,
			}
		}
	}
	return ShiftResponse{}
}

func GetUser(userTaskID string, user []User) User {
	for _, u := range user {
		if u.ID == userTaskID {
			return u
		}
	}
	return User{}
}

func GetSlots(taskId string, slots []Slot) []Slot {
	var response []Slot
	for _, s := range slots {
		if s.TaskID == taskId {
			response = append(response, s)
		}
	}

	return response
}
