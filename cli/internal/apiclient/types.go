// Package apiclient is a thin HTTP client for quantara-core's REST API.
package apiclient

import "time"

// Project mirrors quantara-core's ProjectResponse DTO.
type Project struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
}

// Deployment mirrors quantara-core's DeploymentResponse DTO.
type Deployment struct {
	ID           int64     `json:"id"`
	DeploymentID string    `json:"deploymentId"`
	ProjectID    int64     `json:"projectId"`
	ContractName string    `json:"contractName"`
	Status       string    `json:"status"`
	CreatedAt    time.Time `json:"createdAt"`
}

// Contract mirrors quantara-core's ContractResponse DTO.
type Contract struct {
	ID              int64     `json:"id"`
	ProjectID       int64     `json:"projectId"`
	DeploymentID    string    `json:"deploymentId"`
	ContractAddress string    `json:"contractAddress"`
	DeploymentHash  string    `json:"deploymentHash"`
	Timestamp       time.Time `json:"timestamp"`
}

// errorResponse mirrors quantara-core's ErrorResponse DTO, returned on 4xx/5xx.
type errorResponse struct {
	Status  int    `json:"status"`
	Error   string `json:"error"`
	Message string `json:"message"`
}
