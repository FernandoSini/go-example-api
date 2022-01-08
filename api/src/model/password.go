package model

type Password struct {
	NewPassword     string `json:"newPassword"`
	CurrentPassword string `json:"currentPassword"`
}
