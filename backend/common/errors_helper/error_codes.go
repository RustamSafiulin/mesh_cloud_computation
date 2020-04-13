package errors_helper

import "github.com/pkg/errors"

type ErrorCode uint

var (
	ErrCreateJwtToken = errors.New("Error creation authorization token.")
	ErrParseAuthorizationHeader = errors.New("Error during parse authorization header.")
	ErrWrongPassword = errors.New("Wrong password.")
	ErrAccountNotExists = errors.New("Account not found.")
	ErrAccountAlreadyExists = errors.New("Account already exists.")
	ErrTaskAlreadyExists = errors.New("Task already exists.")
	ErrTaskNotExists = errors.New("Task not found.")
	ErrPasswordHashGeneration = errors.New("Error during generate hash from password.")
	ErrStorageError = errors.New("Storage error.")
	ErrFileCreation = errors.New("File creation error.")
	ErrParseFormFileHeader = errors.New("Form file header parse error.")
	ErrWriteFile = errors.New("Write file error.")
	ErrAccountIdNotFoundInContext = errors.New("Account ID not found in context.")
)