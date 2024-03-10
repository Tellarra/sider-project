package domain

import (
	"errors"
	"slices"
	"time"

	"github.com/adrichard/siderproject/internal/model"
)

func FilterTasksByStatus(shifts []model.Shift) []model.Shift {
	var filteredTasks []model.Shift
	for _, shift := range shifts {
		if shift.Status == "cancelled" {
			continue
		}
		// end date is before the time.Now()
		// means it's in the past
		if shift.Time.EndDate.String() < time.Now().String() {
			shift.Status = "done"
		} else if shift.Time.StartDate.String() < time.Now().String() {
			shift.Status = "ongoing"
		} else {
			shift.Status = "upcoming"
		}
		filteredTasks = append(filteredTasks, shift)
	}
	return filteredTasks
}

func BuildResponse(tasks []model.Task, orga []model.Orga, shift []model.Shift, user []model.User, slots []model.Slot) []model.TaskResponse {
	var response []model.TaskResponse
	for _, task := range tasks {
		orga := model.GetOrga(task.OrganisationID, orga)
		shift := model.GetShift(task.ID, task.ShiftIDs, shift)
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
				Status: shift.Status,
			},
		})
	}
	return response
}

func UpdateTask(assigneeId string, taskId string, tasks []model.Task) error {
	index := slices.IndexFunc(tasks, func(t model.Task) bool {
		return t.ID == taskId
	})
	if index != -1 {
		tasks[index].AssigneeID = assigneeId
		return nil
	}
	return errors.New("task not found")
}
