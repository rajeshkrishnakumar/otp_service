package structs

type Otp struct {
	Otp         string `json:"otp"`
	RetryCount  int    `json:"retry_count"`
	ResentCount int    `json:"resend_count"`
	ResendLock  bool   `json:"resend_lock"`
	RetryLock   bool   `json:"retry_lock"`
}

type OtpType struct {
	OtpLength          int
	KeyPrefix          string
	OtpValidityMinutes int
	RetryLimit         int
	MobileMessage      string
	RetryLockMinutes   int
	ResendLockMinutes  int
}
