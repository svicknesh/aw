package aw

import "time"

const (
	//Endpoint - default endpoint for the cloud instance
	Endpoint = "https://cloud.appwrite.io/v1"
)

// Model - default fields for all models
type Model struct {
	ID           string    `json:"$id"`
	Tenant       string    `json:"$tenant"`
	CreatedAt    time.Time `json:"$createdAt"`
	UpdatedAt    time.Time `json:"$updatedAt"`
	Permissions  []string  `json:"$permissions"`
	DatabaseID   string    `json:"$databaseId"`
	CollectionID string    `json:"$collectionId"`
}
