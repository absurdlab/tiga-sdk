// Package commonv1 provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.8.3 DO NOT EDIT.
package commonv1

// ClientInfo contains public client information that can be used to visually identity the client.
type ClientInfo struct {
	// Name of the registered client.
	ClientName string `json:"client_name"`

	// Optional URL for the client's home page.
	ClientUri string `json:"client_uri"`

	// Optional list of emails that a user may contact for questions.
	Contacts []string `json:"contacts"`

	// Optional URL for the client's logo image.
	LogoUri string `json:"logo_uri"`

	// Optional URL for the client's privacy policy page.
	PolicyUri string `json:"policy_uri"`

	// Optional URL for the client's terms of service page.
	TosUri string `json:"tos_uri"`
}

// ErrorCallback defines model for ErrorCallback.
type ErrorCallback struct {
	// Error code.
	Error string `json:"error"`

	// Optional human-readable explanation of the error.
	ErrorDescription string `json:"error_description"`

	// Unix timestamp of the time the error occurred.
	Timestamp int64 `json:"timestamp"`

	// Optional trace id at the client site to help debugging.
	TraceId string `json:"trace_id"`
}

// ErrorResponse defines model for ErrorResponse.
type ErrorResponse struct {
	// Error code.
	Error string `json:"error"`

	// Optional human-readable explanation to the error.
	ErrorDescription string `json:"error_description"`

	// URI path of the request.
	Path string `json:"path"`

	// Optional request id.
	RequestId string `json:"request_id"`

	// HTTP status.
	Status int `json:"status"`

	// Unix timestamp of the time the error occurred.
	Timestamp int64 `json:"timestamp"`

	// Optional trace id.
	TraceId string `json:"trace_id"`
}

// Links defines model for Links.
type Links struct {
	// URL to callback an error response.
	ErrorCallbackUrl string `json:"error_callback_url"`

	// URL to redirect user to in order to resume processing.
	ResumeUrl string `json:"resume_url"`

	// URL to callback a successful response.
	SuccessCallbackUrl string `json:"success_callback_url"`
}