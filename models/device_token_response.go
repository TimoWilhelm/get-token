package models

type DeviceTokenResponse struct {
	UserCode        string `json:"user_code,omitempty"`
	DeviceCode      string `json:"device_code,omitempty"`
	VerificationURL string `json:"verification_url,omitempty"`
	ExpiresIn       int64  `json:"expires_in,omitempty"`
	Interval        int64  `json:"interval,omitempty"`
	Message         string `json:"message,omitempty"`
}
