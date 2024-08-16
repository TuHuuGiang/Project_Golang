package response

const (
	CodeSuccess           = 2001 // Success
	ErrorCodeParamInvalid = 2002 // Email is invalid
)

// message
var msg = map[int]string{
	CodeSuccess:           "Success",
	ErrorCodeParamInvalid: "Email is invalid",
}
