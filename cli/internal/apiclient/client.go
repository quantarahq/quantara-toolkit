package apiclient

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

// Client talks to a running quantara-core instance over its REST API —
// the same API quantara-web uses. It is a read-only subset today.
type Client struct {
	baseURL    string
	httpClient *http.Client
}

// New returns a Client pointed at baseURL (e.g. "http://localhost:8080").
func New(baseURL string) *Client {
	return &Client{
		baseURL:    strings.TrimRight(baseURL, "/"),
		httpClient: &http.Client{Timeout: 10 * time.Second},
	}
}

// StatusError is returned when quantara-core responds with a non-2xx status.
type StatusError struct {
	StatusCode int
	Message    string
}

func (e *StatusError) Error() string {
	if e.Message != "" {
		return e.Message
	}
	return fmt.Sprintf("quantara-core returned HTTP %d", e.StatusCode)
}

func (c *Client) get(ctx context.Context, path string, out any) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.baseURL+path, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("could not reach quantara-core at %s: %w", c.baseURL, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		var apiErr errorResponse
		_ = json.NewDecoder(resp.Body).Decode(&apiErr)
		return &StatusError{StatusCode: resp.StatusCode, Message: apiErr.Message}
	}

	return json.NewDecoder(resp.Body).Decode(out)
}

// ListProjects calls GET /api/projects.
func (c *Client) ListProjects(ctx context.Context) ([]Project, error) {
	var projects []Project
	if err := c.get(ctx, "/api/projects", &projects); err != nil {
		return nil, err
	}
	return projects, nil
}

// GetProject calls GET /api/projects/{id}.
func (c *Client) GetProject(ctx context.Context, id string) (*Project, error) {
	var project Project
	if err := c.get(ctx, "/api/projects/"+id, &project); err != nil {
		return nil, err
	}
	return &project, nil
}

// ListDeployments calls GET /api/projects/{id}/deployments.
func (c *Client) ListDeployments(ctx context.Context, projectID string) ([]Deployment, error) {
	var deployments []Deployment
	if err := c.get(ctx, "/api/projects/"+projectID+"/deployments", &deployments); err != nil {
		return nil, err
	}
	return deployments, nil
}

// ListContracts calls GET /api/projects/{id}/contracts.
func (c *Client) ListContracts(ctx context.Context, projectID string) ([]Contract, error) {
	var contracts []Contract
	if err := c.get(ctx, "/api/projects/"+projectID+"/contracts", &contracts); err != nil {
		return nil, err
	}
	return contracts, nil
}
