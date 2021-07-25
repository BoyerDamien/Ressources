package errors

// ErrResponse
//
// swagger:model
type ErrResponse struct {

	// Message d'erreur
	// example: error
	// required: true
	Message string `json:"message"`
}
