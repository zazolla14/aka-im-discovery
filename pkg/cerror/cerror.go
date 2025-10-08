package cerror

import (
	"errors"
)

var (
	ErrEntryTypeInvalid       error = errors.New("invalid type of entry type")
	ErrTransactionTypeInvalid error = errors.New("invalid type of transaction type")
	ErrDiscoverIDRequired     error = errors.New("discover id is required")
	ErrDescriptionEnRequired  error = errors.New("description of transaction english is required")
	ErrDescriptionZhRequired  error = errors.New("description of transaction zh is required")
	ErrReferenceCodeRequired  error = errors.New("reference code of transaction is required")
	ErrImpactedItemRequired   error = errors.New("impacted item is required")
	ErrInsufficientBalance    error = errors.New("insufficient balance")
	ErrDiscoverNotFound       error = errors.New("discover not found")
	ErrInvalidFormatStartDate error = errors.New("invalid format start date (format: YYYY-MM-DD)")
	ErrInvalidFormatEndDate   error = errors.New("invalid format end date (format: YYYY-MM-DD)")
	ErrUserIDNotFoundCtx      error = errors.New("user id not found in context")
	ErrDiscoverExisted        error = errors.New("discover already existed")
	ErrInvalidStatusRequest   error = errors.New("invalid status request deposit")
)
