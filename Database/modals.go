package database

import "time"

type User struct {
	Name      string
	Email     string
	Pass      string
	CreatedAt time.Time
}

type UserStory struct {
	Title      string
	HFVersion  string
	DueDate    string
	Comments   string
	CreatedAt  time.Time
	ModifiedAt time.Time
}

type WorkLog struct {
	Title           string
	TaskType        string
	JiraID          string
	Engineer        string
	EstimatedEffort string
	ActualEffort    string
	StartDate       string
	ResolvedData    string
	Comments        string
}
