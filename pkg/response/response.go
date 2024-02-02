package response

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Response struct {
	Success bool        `json:"success" default:"true"`
	Message string      `json:"message" default:"OK"`
	Error   interface{} `json:"error"`
	Data    interface{} `json:"data"`
}

type ErrorConstant struct {
	Response     Response
	Code         int
	ErrorMessage error
}

type Msger struct {
	Kolom string `json:"kolom"`
	Pesan string `json:"pesan"`
}

func (r *ErrorConstant) Error() string {
	return fmt.Sprintf("error code %d,", r.Code)
}

func (r *ErrorConstant) Builder() *ErrorConstant {
	return r
}

const (
	E_DUPLICATE        = "Duplicate data entry"
	E_NOT_FOUND        = "Data not found"
	E_ENPROCESS_ENTITY = "Invalid parameter or payload"
	E_UNAUTHORIZED     = "Unaouthorized access"
	E_BAD_REQUEST      = "Bad request"
	E_SERVER_ERROR     = "Internal server error"
)

var (
	ErrDuplicate = ErrorConstant{
		Response: Response{
			Success: false,
			Message: E_DUPLICATE,
		},
	}

	ErrNotFound = ErrorConstant{
		Response: Response{
			Success: false,
			Message: E_NOT_FOUND,
		},
	}

	ErrEnprocessableEntry = ErrorConstant{
		Response: Response{
			Success: false,
			Message: E_ENPROCESS_ENTITY,
		},
	}

	ErrUnauthorized = ErrorConstant{
		Response: Response{
			Success: false,
			Message: E_UNAUTHORIZED,
		},
		Code: http.StatusUnauthorized,
	}

	ErrBadRequest = ErrorConstant{
		Response: Response{
			Success: false,
			Message: E_BAD_REQUEST,
		},
		Code: http.StatusBadRequest,
	}

	ErrServerError = ErrorConstant{
		Response: Response{
			Success: false,
			Message: E_SERVER_ERROR,
		},
		Code: http.StatusInternalServerError,
	}
)

func BuildError(err ErrorConstant, msg error) error {
	err.ErrorMessage = msg
	return &err
}

func BuildCustomError(code int, message string) error {
	return &ErrorConstant{
		Response: Response{
			Success: false,
			Message: message,
		},
		Code:         code,
		ErrorMessage: fmt.Errorf(message),
	}
}

func BuildCustomErrorWithData(code int, message string, data interface{}) error {
	return &ErrorConstant{
		Response: Response{
			Success: false,
			Message: message,
			Data:    data,
		},
		Code:         code,
		ErrorMessage: fmt.Errorf(message),
	}
}

func SuccessResponse(c echo.Context, code int, msg string, data interface{}) error {
	response := Response{
		Success: true,
		Message: msg,
		Data:    data,
	}
	return c.JSON(code, response)
}

func InvalidResponse(c echo.Context, code int, msg string, m interface{}, data interface{}) error {
	response := Response{
		Success: false,
		Message: msg,
		Error:   m,
		Data:    data,
	}
	return c.JSON(code, response)
}

func ErrorResponse(c echo.Context, err error) error {
	re, ok := err.(*ErrorConstant)
	if ok {
		return c.JSON(re.Builder().Code, re.Builder().Response)
	} else {
		return c.JSON(re.Builder().Code, re.Builder().Response)
	}
}