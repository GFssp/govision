package domain

type JobMessage struct {
	JobID    string `json:"job_id"`
	ImageURL string `json:"image_url"`
}