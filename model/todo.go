package model

import "time"

type (
	// A TODO expresses a task to be done.
	TODO struct {
		ID          int       `json:"id"`
		Subject     string    `json:"subject"`
		Description string    `json:"description"`
		CreatedAt   time.Time `json:"created_at"`
		UpdatedAt   time.Time `json:"updated_at"`
	}

	// A CreateTODORequest expresses a request to create a TODO.
	CreateTODORequest struct {
		Subject     string    `json:"subject"`
		Description string    `json:"description"`
		CreatedAt   time.Time `json:"created_at"`
	}

	// A CreateTODOResponse expresses a response after creating a TODO.
	CreateTODOResponse struct {
		TODO TODO `json:"todo"`
	}

	// A ReadTODORequest expresses a request to read a TODO.
	ReadTODORequest struct {
		ID int `json:"id"`
	}

	// A ReadTODOResponse expresses a response after reading a TODO.
	ReadTODOResponse struct {
		ID          int       `json:"id"`
		Subject     string    `json:"subject"`
		Description string    `json:"description"`
		CreatedAt   time.Time `json:"created_at"`
		UpdatedAt   time.Time `json:"updated_at"`
	}

	// A UpdateTODORequest expresses a request to update a TODO.
	UpdateTODORequest struct {
		ID          int       `json:"id"`
		Subject     string    `json:"subject"`
		Description string    `json:"description"`
		UpdatedAt   time.Time `json:"updated_at"`
	}

	// A UpdateTODOResponse expresses a response after updating a TODO.
	UpdateTODOResponse struct {
		ID        int       `json:"id"`
		Subject   string    `json:"subject"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	// A DeleteTODORequest expresses a request to delete a TODO.
	DeleteTODORequest struct {
		ID int `json:"id"`
	}

	// A DeleteTODOResponse expresses a response after deleting a TODO.
	DeleteTODOResponse struct{}
)
