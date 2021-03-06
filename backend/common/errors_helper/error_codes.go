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
	ErrCreateDirectory = errors.New("Error was caused during directory creation.")
	ErrAccountIdNotFoundInContext = errors.New("Account ID not found in context.")
	ErrStartComputationTask = errors.New("Error was caused during start computation task.")
	ErrTaskDataFileNotExists = errors.New("Task data file not found.")
)