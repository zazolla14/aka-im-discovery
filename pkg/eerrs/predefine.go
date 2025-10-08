package eerrs

import "github.com/1nterdigital/aka-im-tools/errs"

var (
	ErrPassword                 = errs.NewCodeError(ErrorCodePasswordError, "PasswordError")
	ErrAccountNotFound          = errs.NewCodeError(ErrorCodeAccountNotFound, "AccountNotFound")
	ErrPhoneAlreadyRegister     = errs.NewCodeError(ErrorCodePhoneAlreadyRegister, "PhoneAlreadyRegister")
	ErrAccountAlreadyRegister   = errs.NewCodeError(ErrorCodeAccountAlreadyRegister, "AccountAlreadyRegister")
	ErrVerifyCodeSendFrequently = errs.NewCodeError(ErrorCodeVerifyCodeSendFrequently, "VerifyCodeSendFrequently")
	ErrVerifyCodeNotMatch       = errs.NewCodeError(ErrorCodeVerifyCodeNotMatch, "VerifyCodeNotMatch")
	ErrVerifyCodeExpired        = errs.NewCodeError(ErrorCodeVerifyCodeExpired, "VerifyCodeExpired")
	ErrVerifyCodeMaxCount       = errs.NewCodeError(ErrorCodeVerifyCodeMaxCount, "VerifyCodeMaxCount")
	ErrVerifyCodeUsed           = errs.NewCodeError(ErrorCodeVerifyCodeUsed, "VerifyCodeUsed")
	ErrInvitationCodeUsed       = errs.NewCodeError(ErrorCodeInvitationCodeUsed, "InvitationCodeUsed")
	ErrInvitationNotFound       = errs.NewCodeError(ErrorCodeInvitationNotFound, "InvitationNotFound")
	ErrForbidden                = errs.NewCodeError(ErrorCodeForbidden, "Forbidden")
	ErrRefuseFriend             = errs.NewCodeError(ErrorCodeRefuseFriend, "RefuseFriend")
	ErrEmailAlreadyRegister     = errs.NewCodeError(ErrorCodeEmailAlreadyRegister, "EmailAlreadyRegister")

	ErrTokenNotExist = errs.NewCodeError(ErrorTokenNotExist, "ErrTokenNotExist")
)
