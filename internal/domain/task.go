package domain

import "github.com/adrichard/siderproject/internal/model"

func filterTasksByStatus(tasks []model.Task) []model.Task {
	var filteredTasks []model.Task
	for _, task := range tasks {
		if task.Status == "open" {
			filteredTasks = append(filteredTasks, task)
		}
	}
	return filteredTasks
}

func BuildResponse(tasks []model.Task, orga []model.Orga, shift []model.Shift, user []model.User, slots []model.Slot) []model.TaskResponse {
	var response []model.TaskResponse
	for _, task := range tasks {
		orga := model.GetOrga(task.OrganisationID, orga)
		shift := model.GetShift(task.ShiftIDs, shift)
		response = append(response, model.TaskResponse{
			Id:   task.ID,
			Name: task.Alias,
			Orga: model.OrgaResponse{
				Name:       orga.Name,
				Address:    orga.Address,
				PictureUrl: orga.PictureUrl,
			},
			Shift: model.ShiftResponse{
				Id:        shift.Id,
				StartDate: shift.StartDate,
				EndDate:   shift.EndDate,
				Slots: model.SlotsResponse{
					Filled: shift.Slots.Filled,
					Total:  shift.Slots.Total,
				},
			},
		})
	}
	return response
}
