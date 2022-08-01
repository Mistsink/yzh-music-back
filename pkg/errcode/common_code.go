package errcode

var (
	Success                   = NewError(0, "Success")
	ServerError               = NewError(10000, "An internal error occurred")
	InvalidParams             = NewError(10001, "Invalid param")
	NotFound                  = NewError(10000002, "Not found")
	UnauthorizedAuthNotExist  = NewError(10000003, "Cannot find the corresponding AppKey and AppSecret")
	UnauthorizedTokenError    = NewError(10000004, "Failed to verify token")
	UnauthorizedTokenTimeout  = NewError(10000005, "Token verification timeout")
	UnauthorizedTokenGenerate = NewError(10000006, "Failed to generate token")
	TooManyRequests           = NewError(10000007, "Too many request")
)

func ServerWithMsg(oneDetail string) *Error {
	return ServerWithDetail([]string{oneDetail})
}

func ServerWithDetail(details []string) *Error {
	return ServerError.WithDetails(details)
}
