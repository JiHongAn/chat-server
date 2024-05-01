package errors

type AppError struct {
	Code    int
	Message string
	Status  int
}

func (e AppError) Error() string {
	return e.Message
}

var (
	InvalidRequest = AppError{
		Code:    -2,
		Message: "요청을 확인해주세요.",
		Status:  400,
	}
	InvalidAccessToken = AppError{
		Code:    -3,
		Message: "액세스 토큰이 유효하지 않습니다.",
		Status:  401,
	}
	ApiNotFound = AppError{
		Code:    -4,
		Message: "API를 찾을 수 없습니다.",
		Status:  404,
	}
	NoPermission = AppError{
		Code:    -5,
		Message: "권한이 없습니다.",
		Status:  403,
	}
	DBError = AppError{
		Code:    -6,
		Message: "데이터베이스 오류가 발생했습니다.",
		Status:  500,
	}
)
