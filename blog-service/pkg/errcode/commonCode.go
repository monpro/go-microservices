package errcode


var (
	Success 					= NewError(0, "Success")
	ServerError 				= NewError(1000, "Server Error")
	InvalidParams 				= NewError(1001, "Invalid Params")
	NotFound 					= NewError(1002, "Not Found")
	UnauthorizedAuthNotExist 	= NewError(1003, "Unauthorized, Cannot find App Key and App Secret")
	UnauthorizedTokenError 		= NewError(1004, "Unauthorized token")
	UnauthorizedTokenTimeOut 	= NewError(1005, "Unauthorized token timeout")
	UnauthorizedTokenGenerate 	= NewError(1006, "Unauthorized token generation failure")
	TooManyRequests 			= NewError(1007, "Too man requests")
)