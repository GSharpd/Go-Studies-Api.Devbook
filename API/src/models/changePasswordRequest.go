package models

// Represents the request for the update password operation
type ChangePasswordRequest struct {
	New     string `json:"new"`
	Current string `json:"current"`
}
