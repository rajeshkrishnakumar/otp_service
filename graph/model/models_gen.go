// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type OtpType struct {
	OtpLength          int    `json:"otpLength"`
	KeyPrefix          string `json:"keyPrefix"`
	OtpValidityMinutes int    `json:"otp_validity_minutes"`
	RetryLimit         int    `json:"retry_limit"`
	ResentLimit        int    `json:"resent_limit"`
	MobileMessage      string `json:"mobile_message"`
	RetryLockMinutes   int    `json:"retry_lock_minutes"`
	ResendLockMinutes  int    `json:"resend_lock_minutes"`
	OtpType            string `json:"otpType"`
}

type Status struct {
	Status       bool   `json:"status"`
	Message      string `json:"message"`
	OtherMessage string `json:"other_message"`
}

type AddOtpType struct {
	OtpLength          int    `json:"otpLength"`
	KeyPrefix          string `json:"keyPrefix"`
	OtpValidityMinutes int    `json:"otp_validity_minutes"`
	RetryLimit         int    `json:"retry_limit"`
	ResentLimit        int    `json:"resent_limit"`
	MobileMessage      string `json:"mobile_message"`
	RetryLockMinutes   int    `json:"retry_lock_minutes"`
	ResendLockMinutes  int    `json:"resend_lock_minutes"`
	OtpType            string `json:"otpType"`
}

type RemoveOtpType struct {
	OtpType string `json:"otpType"`
}

type SendOtp struct {
	Channel string `json:"channel"`
	OtpType string `json:"otp_type"`
}

type VerifyOtp struct {
	Channel string `json:"channel"`
	Otp     string `json:"otp"`
	OtpType string `json:"otp_type"`
}
