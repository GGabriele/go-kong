package kong

import (
	"context"
	"fmt"
	"net/http"
)

// AbstractEntity handles entities in Kong.
type AbstractEntity interface {
	// Validate validates an entity against its schema
	Validate(ctx context.Context, entity string) (bool, error)
}

// Entity handles entities in Kong.
type Entity service

// Validate validates an arbitrary entity against its schema
func (e *Entity) Validate(ctx context.Context, entity string) (bool, error) {
	endpoint := fmt.Sprintf("/schemas/%s/validate", entity)
	req, err := e.client.NewRequest("POST", endpoint, nil, nil)
	if err != nil {
		return false, err
	}
	resp, err := e.client.Do(ctx, req, nil)
	if err != nil {
		return false, err
	}
	return resp.StatusCode == http.StatusOK, nil
}
