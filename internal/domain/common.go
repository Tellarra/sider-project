package domain

import (
	"embed"
	"fmt"

	"github.com/adrichard/siderproject/internal/model"
)

//go:embed dumpRes/*.json
var filesToStore embed.FS

func GetAllFilesToRightTypes() ([]model.DocumentToIndex, error) {
	taskFile, err := filesToStore.ReadFile("dumpRes/tasks.json")
	if err != nil {
		return nil, fmt.Errorf("error reading file tasks.json: %v", err)
	}
	orgaFile, err := filesToStore.ReadFile("dumpRes/orgas.json")
	if err != nil {
		return nil, fmt.Errorf("error reading file organisations.json: %v", err)
	}
	shiftFile, err := filesToStore.ReadFile("dumpRes/shifts.json")
	if err != nil {
		return nil, fmt.Errorf("error reading file shifts.json: %v", err)
	}
	userFile, err := filesToStore.ReadFile("dumpRes/users.json")
	if err != nil {
		return nil, fmt.Errorf("error reading file users.json: %v", err)
	}
	slotsFile, err := filesToStore.ReadFile("dumpRes/slots.json")
	if err != nil {
		return nil, fmt.Errorf("error reading file slots.json: %v", err)
	}

	documents := []model.DocumentToIndex{
		{Data: taskFile, Name: "task"},
		{Data: orgaFile, Name: "orga"},
		{Data: shiftFile, Name: "shift"},
		{Data: userFile, Name: "user"},
		{Data: slotsFile, Name: "slots"},
	}

	return documents, nil
}
