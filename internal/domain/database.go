package domain

import (
	"database/sql"
	"log"

	"github.com/adrichard/siderproject/internal/model"
)

func InsertToDb(tasks []model.Task, db *sql.DB) {
	// Insert data
	for _, task := range tasks {
		// Simplified example: only inserting a few fields
		_, err := db.Exec(`INSERT INTO tasks (id, logAsId, creatorId) VALUES (?, ?, ?)`,
			task.ID, task.LogAsID, task.CreatorID)
		if err != nil {
			log.Fatal(err)
		}
	}
}
