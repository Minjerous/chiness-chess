package statuscode

const (
	_ int = iota

	CodeUnacceptedParam
	CodeTimeout
	CodeInternalErr
	AccessDeniedErr
	EmptyMetadataErr
)
