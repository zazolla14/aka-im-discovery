package eerrs

const (
	ErrorCodePasswordError = 20001 + iota
	ErrorCodeAccountNotFound
	ErrorCodePhoneAlreadyRegister
	ErrorCodeAccountAlreadyRegister
	ErrorCodeVerifyCodeSendFrequently
	ErrorCodeVerifyCodeNotMatch
	ErrorCodeVerifyCodeExpired
	ErrorCodeVerifyCodeMaxCount
	ErrorCodeVerifyCodeUsed
	ErrorCodeInvitationCodeUsed
	ErrorCodeInvitationNotFound
	ErrorCodeForbidden
	ErrorCodeRefuseFriend
	ErrorCodeEmailAlreadyRegister
)

const (
	ErrorTokenNotExist = 20101 + iota
)
