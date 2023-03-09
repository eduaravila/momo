package errors

type ErrorType string

var (
	ErrorTypeUnknown       ErrorType = "unknown"
	ErrorTypeAuthorization ErrorType = "authorization"
	ErrorTypeInvalidInput  ErrorType = "invalid-input"
)

type SlugError struct {
	slug      string
	message   string
	errorType ErrorType
}

func NewSlugError(slug string, message string, errorType ErrorType) *SlugError {
	return &SlugError{slug: slug, message: message, errorType: errorType}
}

func (s *SlugError) Slug() string {
	return s.slug
}

func (s *SlugError) Message() string {
	return s.message
}

func (s *SlugError) ErrorType() ErrorType {
	return s.errorType
}

func NewAuthorizationError(slug string, message string) *SlugError {
	return NewSlugError(slug, message, ErrorTypeAuthorization)
}

func NewInvalidInputError(slug string, message string) *SlugError {
	return NewSlugError(slug, message, ErrorTypeInvalidInput)
}
