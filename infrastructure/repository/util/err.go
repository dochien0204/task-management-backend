package util

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	ginI18n "github.com/gin-contrib/i18n"

	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgconn"
	"github.com/lib/pq"
)

func IsErrorCode(err error, errcode pq.ErrorCode) bool {
	if pgerr, ok := err.(*pgconn.PgError); ok {
		return pq.ErrorCode(pgerr.Code) == errcode
	}
	return false
}

func ParseError(err error) string {
	out := ""

	switch typedError := any(err).(type) {
	case validator.ValidationErrors:
		// if the type is validator.ValidationErrors then it's actually an array of validator.FieldError so we'll
		// loop through each of those and convert them one by one
		for _, e := range typedError {
			out = parseFieldError(e)
		}

	case *json.UnmarshalTypeError:
		// similarly, if the error is an unmarshalling error we'll parse it into another, more readable string format
		out = parseMarshallingError(*typedError)

	default:
		out = ginI18n.MustGetMessage(err.Error())
	}

	return out
}

func parseFieldError(e validator.FieldError) string {
	// workaround to the fact that the `gt|gtfield=Start` gets passed as an entire tag for some reason
	// https://github.com/go-playground/validator/issues/926
	fieldPrefix := fmt.Sprintf("The field %s", e.Field())
	tag := strings.Split(e.Tag(), "|")[0]
	switch tag {
	case "required_without":
		return fmt.Sprintf("%s is required if %s is not supplied", fieldPrefix, e.Param())
	case "lt", "ltfield":
		param := e.Param()
		if param == "" {
			param = time.Now().Format(time.RFC3339)
		}
		return fmt.Sprintf("%s must be less than %s", fieldPrefix, param)

	case "gt", "gtfield":
		param := e.Param()
		if param == "" {
			param = time.Now().Format(time.RFC3339)
		}
		return fmt.Sprintf("%s must be greater than %s", fieldPrefix, param)

	default:
		return e.Error()
	}
}
func parseMarshallingError(e json.UnmarshalTypeError) string {
	return fmt.Sprintf("Field %s must have type %s", e.Field, e.Type.String())
}
