package response

const (
	ErrCodeSuccess      = 20001 // Success
	ErrCodeParamInvalid = 20003 // Email is invalid
	ErrInvalidToken     = 30001 // Token is invalid
	ErrInvalidOTP       = 30002 // OTP is invalid
	ErrSendEmailOTP     = 30003

	// User Authentication
	ErrCodeAuthFailed = 40005

	// Register Code
	ErrCodeUserHasExists = 50001 // user has already registered
	ErrCodeUserNotExists = 50002

	// Err Login
	ErrCodeOtpNotExists     = 60009
	ErrCodeUserOtpNotExists = 60008

	// Two factor Authentication
	ErrCodeTwoFactorAuthSetupFailed = 80001
)

var msg = map[int]string{
	ErrCodeSuccess:      "Success",
	ErrCodeParamInvalid: "Email is invalid",
	ErrInvalidToken:     "Token is invalid",
	ErrInvalidOTP:       "OTP is invalid",
	ErrSendEmailOTP:     "Failed to send email otp",

	ErrCodeAuthFailed: "Auth failed",

	ErrCodeUserHasExists: "User has already registered",
	ErrCodeUserNotExists: "User not exist",

	ErrCodeOtpNotExists:     "Otp exist but not registered",
	ErrCodeUserOtpNotExists: "OTP not exists",

	ErrCodeTwoFactorAuthSetupFailed: "Setup two factor authentication failed",
}
