package models

const (
	AuthorizationPending = "authorization_pending"
)

type TokenErrorResponse struct {
	Error            string  `json:"error,omitempty"`
	ErrorDescription string  `json:"error_description,omitempty"`
	ErrorCodes       []int64 `json:"error_codes,omitempty"`
	Timestamp        string  `json:"timestamp,omitempty"`
	TraceID          string  `json:"trace_id,omitempty"`
	CorrelationID    string  `json:"correlation_id,omitempty"`
	ErrorUri         string  `json:"error_uri,omitempty"`
}

type TokenSucessResponse struct {
	TokenType    string `json:"token_type,omitempty"`
	Scope        string `json:"scope,omitempty"`
	ExpiresIn    int64  `json:"expires_in,omitempty"`
	ExtExpiresIn int64  `json:"ext_expires_in,omitempty"`
	AccessToken  string `json:"access_token,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
	IDToken      string `json:"id_token,omitempty"`
}
