package otp

import "time"

type Otp struct {
	// retention map comenzará la retención dado el período establecido
	rm map[string]otp
}

// One-Time Passwords (OTP)
type otp struct {
	key     string
	created time.Time
}
