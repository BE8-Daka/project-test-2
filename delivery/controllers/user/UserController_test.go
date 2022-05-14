package user

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"project-test/delivery/views/response"
	"project-test/entity"
	"strings"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {
	t.Run("Status OK", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"name":     "testing",
			"username": "testing",
			"no_hp":    "098765433212",
			"email":    "testing@gmail.com",
			"password": "testing",
		})
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/users")
		controller := NewUserController(&mockUser{}, validator.New())
		controller.Insert()(context)

		type Response struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
			Data    interface{}
		}

		var resp Response
		json.Unmarshal([]byte(res.Body.Bytes()), &resp)

		assert.Equal(t, 201, resp.Code)
		assert.Equal(t, "successfully created", resp.Message)
		assert.Equal(t, map[string]interface{}(map[string]interface{}{"created_at":"0001-01-01T00:00:00Z", "email":"testing@gmail.com", "name":"testing", "no_hp":"098765433212", "password":"testing", "username":"testing"}), resp.Data)
	})

	t.Run("Status BadRequest Bind", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"name":     1,
			"username": "testing",
			"no_hp":    "098765433212",
			"email":    "testing@gmail.com",
			"password": "testing",
		})

		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/users")
		controller := NewUserController(&mockUser{}, validator.New())
		controller.Insert()(context)

		type Response struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
			Data    interface{}
		}

		var resp Response
		json.Unmarshal([]byte(res.Body.Bytes()), &resp)

		assert.Equal(t, 400, resp.Code)
		assert.Equal(t, "field=name, expected=string", resp.Message)
		assert.Nil(t, resp.Data)
	})

	t.Run("Status BadRequest Required", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"name":     "",
			"username": "testing",
			"no_hp":    "098765433212",
			"email":    "testing@gmail.com",
			"password": "testing",
		})

		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/users")
		controller := NewUserController(&mockUser{}, validator.New())
		controller.Insert()(context)

		type Response struct {
			Code    int    `json:"code"`
			Message []string `json:"message"`
			Data    interface{}
		}

		var resp Response
		json.Unmarshal([]byte(res.Body.Bytes()), &resp)

		assert.Equal(t, 400, resp.Code)
		assert.Equal(t, []string{"field Name : required"}, resp.Message)
		assert.Nil(t, resp.Data)
	})

	t.Run("Status BadRequest Duplicate", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"name":     "testing",
			"username": "testing",
			"no_hp":    "098765433212",
			"email":    "testing@gmail.com",
			"password": "testing",
		})

		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/users")
		controller := NewUserController(&mockError{}, validator.New())
		controller.Insert()(context)

		type Response struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
			Data    interface{}
		}

		var resp Response
		json.Unmarshal([]byte(res.Body.Bytes()), &resp)

		assert.Equal(t, 400, resp.Code)
		assert.Equal(t, "field  : duplicate!", resp.Message)
		assert.Nil(t, resp.Data)
	})
}

type mockUser struct {}

func (m *mockUser) Insert(user *entity.User) (response.InsertUser, error) {
	return response.InsertUser{
		Name: user.Name,
		Username: user.Username,
		Email: user.Email,
		NoHp: user.NoHp,
		Password: user.Password,
		CreatedAt: user.CreatedAt,
	}, nil
}

type mockError struct {}

func (m *mockError) Insert(user *entity.User) (response.InsertUser, error) {
	return response.InsertUser{}, errors.New("field '...' : duplicate!")
}