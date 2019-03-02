package config

import (
	"encoding/json"
	"io"
	"net/http"
)

var (
	ErrorJsonFailed = ErrorResult{
		HttpCode: 400,
		Error: Err{
			Error:     "request body can not parse",
			ErrorCode: "001",
		},
	}

	ErrorReadBodyFailed = ErrorResult{
		HttpCode: 401,
		Error: Err{
			Error:     "read request Body err",
			ErrorCode: "002",
		},
	}

	ErrorPassword = ErrorResult{
		HttpCode: 409,
		Error: Err{
			Error:     "user register err",
			ErrorCode: "003",
		},
	}
	ErrorPasswordNotSame = ErrorResult{
		HttpCode: 409,
		Error: Err{
			Error:     "the password not same",
			ErrorCode: "004",
		},
	}

	ErrorUserName = ErrorResult{
		HttpCode: 409,
		Error: Err{
			Error:     "the English not in Username",
			ErrorCode: "005",
		},
	}

	ErrorEmail = ErrorResult{
		HttpCode: 409,
		Error: Err{
			Error:     "User register email not right",
			ErrorCode: "006",
		},
	}

	ErrorRegisterTime = ErrorResult{
		HttpCode: 409,
		Error: Err{
			Error:     "Register time not right",
			ErrorCode: "007",
		},
	}

	ErrorRpcConnFailed = ErrorResult{
		HttpCode: 409,
		Error: Err{
			Error:     "cant not conn",
			ErrorCode: "008",
		},
	}

	ErrorMethodFailed = ErrorResult{
		HttpCode: 409,
		Error: Err{
			Error:     "the method err",
			ErrorCode: "009",
		},
	}
	ErrorNotRequest = ErrorResult{
		HttpCode: 403,
		Error: Err{
			Error:     "the method err",
			ErrorCode: "010",
		},
	}

	ErrorTimeOut = ErrorResult{
		HttpCode: 408,
		Error: Err{
			Error:     "the method time out",
			ErrorCode: "011",
		},
	}

	ErrorRequestFaild = ErrorResult{
		HttpCode: 400,
		Error: Err{
			Error:     "request Failed",
			ErrorCode: "012",
		},
	}

	ErrorCall = ErrorResult{
		HttpCode: 403,
		Error: Err{
			Error:     "request Failed",
			ErrorCode: "013",
		},
	}

	ErrorRepeat = ErrorResult{
		HttpCode: 403,
		Error: Err{
			Error:     "the configuration repetition",
			ErrorCode: "014",
		},
	}

	DbError = ErrorResult{
		HttpCode: 500,
		Error: Err{
			Error:     "configuration mysql err",
			ErrorCode: "015",
		},
	}

	PermissionError = ErrorResult{
		HttpCode: 500,
		Error: Err{
			Error:     "permission denied",
			ErrorCode: "016",
		},
	}

	OperationDbErr = ErrorResult{
		HttpCode: 403,
		Error: Err{
			Error:     "operation mysql err",
			ErrorCode: "017",
		},
	}

	ErrorSrvName = ErrorResult{
		HttpCode: 403,
		Error: Err{
			Error:     "confCenter no this servicename",
			ErrorCode: "018",
		},
	}
)

type Err struct {
	Error     string `json:"error"`
	ErrorCode string `json:"error_code"`
}

type ErrorResult struct {
	Error    Err
	HttpCode int
}

func NewErrorResult() *ErrorResult {
	return &ErrorResult{}
}

func (r ErrorResult) SendErrorResponse(w http.ResponseWriter, errResponse ErrorResult) {
	w.WriteHeader(errResponse.HttpCode)
	errMessage, _ := json.Marshal(&errResponse.Error)
	io.WriteString(w, string(errMessage))
}

type NormalResult struct {
	Resp string
	Code int
}

func NewResult() *NormalResult {
	return &NormalResult{
		Resp: "ok",
		Code: 200,
	}
}

func (r *NormalResult) Response(w http.ResponseWriter) {
	w.WriteHeader(r.Code)
	io.WriteString(w, r.Resp)
}

func (r *NormalResult) NormalResponse(w http.ResponseWriter, result *NormalResult) {
	w.WriteHeader(result.Code)
	io.WriteString(w, result.Resp)
}
