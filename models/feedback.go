package models

type Feedback struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Message string `json:"message"`
}
