package errors_helper

type ErrorCode uint

const (
	Success ErrorCode = iota
	ErrCreateJwtToken
	ErrParseAuthorizationHeader
	ErrWrongPassword
	ErrAccountNotExists
	ErrAccountAlreadyExists
	ErrTaskAlreadyExists
	ErrTaskNotExists
	ErrPasswordHashGeneration
	ErrStorageError
)

func (e ErrorCode) Ordinal() uint {
	return uint(e)
}

type insertionString struct {
	ErrorCode ErrorCode
	msgString string
}

var insertionStrings = []insertionString{
	{ ErrorCode: Success, msgString: "Success" },
	{ ErrorCode: ErrCreateJwtToken, msgString: "Error creation authorization token" },
	{ ErrorCode: ErrParseAuthorizationHeader, msgString: "Error during parse authorization header" },
	{ ErrorCode: ErrWrongPassword, msgString: "Wrong password" },
	{ ErrorCode: ErrAccountNotExists, msgString: "Account not found. ID: %v"},
	{ ErrorCode: ErrTaskNotExists, msgString: "Task not found. ID: %v" },
	{ ErrorCode: ErrAccountAlreadyExists, msgString: "Account already exists"},
	{ ErrorCode: ErrTaskAlreadyExists, msgString: "Task already exists" },
	{ ErrorCode: ErrPasswordHashGeneration, msgString: "Error during generate hash from password" },
	{ ErrorCode: ErrStorageError, msgString: "Storage error. Reason: %v" },
}