package Entities

type ProjectEntity struct {
	Id            int    `json:"id"`
	WorkspaceId   int    `json:"workspace_id"`
	Name          string `json:"name"`
	ActualHours   int    `json:"actual_hours"`
	ActualSeconds int    `json:"actual_seconds"`
}
