package todo

import (
	"encoding/json"
	"os"
)

const dataFile = "tasks.json"

// Load reads the task list from the JSON file
func Load() (*TaskList, error) {
	tl := NewTaskList()

	data, err := os.ReadFile(dataFile)
	if err != nil {
		if os.IsNotExist(err) {
			return tl, nil // Return empty list if file doesn't exist
		}
		return nil, err
	}

	if err := json.Unmarshal(data, tl); err != nil {
		return nil, err
	}

	return tl, nil
}

// Save writes the task list to the JSON file
func Save(tl *TaskList) error {
	data, err := json.MarshalIndent(tl, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(dataFile, data, 0644)
}
