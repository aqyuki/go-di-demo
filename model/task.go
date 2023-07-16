package model

type (
	// Task is registered task data
	Task struct {
		// Task is title of registered task
		Title string
		// Description is description of registered task
		Description string
	}
)

// New create new instance of Task and return it
func New(title string, description string) *Task {
	return &Task{
		Title:       title,
		Description: description,
	}
}
