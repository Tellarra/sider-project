package model

type Response struct {
	Task TaskResponse `json:"task"`
}

type TaskResponse struct {
	Id    string        `json:"id"`
	Name  string        `json:"name"`
	Orga  OrgaResponse  `json:"organisation"`
	Shift ShiftResponse `json:"shift"`
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
}

type SlotsResponse struct {
	Filled int `json:"filled"`
	Total  int `json:"total"`
}

func GetOrga(orgaTaskID string, orga []Orga) OrgaResponse {
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

func GetShift(shiftTaskIDs []string, shift []Shift) ShiftResponse {
	for _, s := range shift {
		for _, shiftTaskID := range shiftTaskIDs {
			if s.ID == shiftTaskID {
				return ShiftResponse{
					Id:        s.ID,
					StartDate: s.Time.StartDate.String(),
					EndDate:   s.Time.EndDate.String(),
					Slots: SlotsResponse{
						Filled: s.Slots,
						Total:  s.Slots,
					},
				}
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
