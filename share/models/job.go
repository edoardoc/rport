package models

import (
	"time"
)

const (
	JobStatusSuccessful = "successful"
	JobStatusRunning    = "running"
	JobStatusFailed     = "failed"
	JobStatusUnknown    = "unknown"
)

type Job struct {
	JobSummary
	SID        string     `json:"sid"`
	Command    string     `json:"command"`
	Shell      string     `json:"shell"`
	PID        *int       `json:"pid"`
	StartedAt  time.Time  `json:"started_at"`
	CreatedBy  string     `json:"created_by"`
	TimeoutSec int        `json:"timeout_sec"`
	MultiJobID *string    `json:"multi_job_id,omitempty"`
	Error      string     `json:"error,omitempty"`
	Result     *JobResult `json:"result"`
}

// JobSummary short info about a job.
type JobSummary struct {
	JID        string     `json:"jid"`
	Status     string     `json:"status"`
	FinishedAt *time.Time `json:"finished_at"`
}

type JobResult struct {
	StdOut string `json:"stdout"`
	StdErr string `json:"stderr"`
}

type MultiJob struct {
	MultiJobSummary
	ClientIDs  []string `json:"client_ids"`
	Command    string   `json:"command"`
	Shell      string   `json:"shell"`
	TimeoutSec int      `json:"timeout_sec"`
	Concurrent bool     `json:"concurrent"`
	AbortOnErr bool     `json:"abort_on_err"`
	Jobs       []*Job   `json:"jobs"`
}

type MultiJobSummary struct {
	JID       string    `json:"jid"`
	StartedAt time.Time `json:"started_at"`
	CreatedBy string    `json:"created_by"`
}

type MultiJobResult struct {
	Status string     `json:"status"`
	StdErr string     `json:"stderr"`
	Result *JobResult `json:"result"`
}
