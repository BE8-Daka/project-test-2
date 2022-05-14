package response

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

func StatusCreated(data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"code": http.StatusCreated,
		"message" : "successfully created",
		"data" : data,
	}
}

func StatusBadRequestRequired(err error) map[string]interface{} {
	var field, tag string
	var message [] string
	
	for _, err := range err.(validator.ValidationErrors) {
		field = fmt.Sprint(err.StructField())
		tag = fmt.Sprint(err.Tag())

		message = append(message, "field "+field+" : "+tag)
	}
	
	return map[string]interface{}{
		"code":    http.StatusBadRequest,
		"message": message,
		"data":    nil,
	}
}

func StatusBadRequestBind(err error) map[string]interface{} {
	var field [] string
	var message string
	
	for i, v := range strings.Fields(string(err.Error())) {
		if i == 1 && v == "message=Syntax" {
			message = "expected=string"
		} else if i == 1 && v == "message=Unmarshal" {
			message = "expected=string"
		} else if i == 6 {
			field = append(field, v)
		}
	}

	return map[string]interface{}{
		"code":    http.StatusBadRequest,
		"message": field[0]+" "+message,
		"data":    nil,
	}
}

func StatusBadRequestDuplicate(err error) map[string]interface{} {
	var message string

	for i, v := range strings.Fields(string(err.Error())) {
		if i ==  7 {
			message = v
		}
	}

	return map[string]interface{}{
		"code":    http.StatusBadRequest,
		"message": "field "+message+" : duplicate!",
		"data":    nil,
	}
}