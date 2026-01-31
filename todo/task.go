package todo

import "time"

// Task represents a single to-do item
type Task struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Completed   bool      `json:"completed"`
	CreatedAt   time.Time `json:"created_at"`
}

// TaskList holds all tasks
type TaskList struct {
	Tasks   []Task `json:"tasks"`
	NextID  int    `json:"next_id"`
}

// NewTaskList creates a new empty task list
func NewTaskList() *TaskList {
	return &TaskList{
		Tasks:  []Task{},
		NextID: 1,
	}
}

// Add creates a new task with the given description
func (tl *TaskList) Add(description string) Task {
	task := Task{
		ID:          tl.NextID,
		Description: description,
		Completed:   false,
		CreatedAt:   time.Now(),
	}
	tl.Tasks = append(tl.Tasks, task)
	tl.NextID++
	return task
}

// Complete marks a task as completed by ID
func (tl *TaskList) Complete(id int) bool {
	for i := range tl.Tasks {
		if tl.Tasks[i].ID == id {
			tl.Tasks[i].Completed = true
			return true
		}
	}
	return false
}

// Delete removes a task by ID
func (tl *TaskList) Delete(id int) bool {
	for i := range tl.Tasks {
		if tl.Tasks[i].ID == id {
			tl.Tasks = append(tl.Tasks[:i], tl.Tasks[i+1:]...)
			return true
		}
	}
	return false
}
