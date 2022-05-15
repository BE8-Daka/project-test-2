package user

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"project-test/delivery/middlewares"
	"project-test/delivery/views/response"
	"project-test/entity"
	"strings"
	"testing"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/stretchr/testify/assert"
)

var (
	token string
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
		context.SetPath("/users/register")
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
		context.SetPath("/users/register")
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
		context.SetPath("/users/register")
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
		context.SetPath("/users/register")
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

func TestCreateToken(t *testing.T) {
	t.Run("Create Token", func(t *testing.T) {
		token, _ = middlewares.CreateToken(1)
	})
}

func TestLogin(t *testing.T) {
	t.Run("Status OK", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"username": "testing",
			"password": "testing",
		})

		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/users/login")
		controller := NewUserController(&mockUser{}, validator.New())
		controller.Login()(context)

		type Response struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
			Data    interface{}
		}

		var resp Response
		json.Unmarshal([]byte(res.Body.Bytes()), &resp)

		assert.Equal(t, 200, resp.Code)
		assert.Equal(t, "successfully", resp.Message)
		assert.NotNil(t, resp.Data)
	})

	t.Run("Status BadRequest Bind", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"username": 1,
			"password": "testing",
		})

		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/users/login")
		controller := NewUserController(&mockUser{}, validator.New())
		controller.Login()(context)

		type Response struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
			Data    interface{}
		}

		var resp Response
		json.Unmarshal([]byte(res.Body.Bytes()), &resp)

		assert.Equal(t, 400, resp.Code)
		assert.Equal(t, "field=username, expected=string", resp.Message)
		assert.Nil(t, resp.Data)
	})

	t.Run("Status BadRequest Required", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"username": "",
			"password": "testing",
		})

		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/users/login")
		controller := NewUserController(&mockUser{}, validator.New())
		controller.Login()(context)

		type Response struct {
			Code    int    `json:"code"`
			Message []string `json:"message"`
			Data    interface{}
		}

		var resp Response
		json.Unmarshal([]byte(res.Body.Bytes()), &resp)

		assert.Equal(t, 400, resp.Code)
		assert.Equal(t, []string{"field Username : required"}, resp.Message)
		assert.Nil(t, resp.Data)
	})

	t.Run("Status Unauthorized", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"username": "testing",
			"password": "testing",
		})

		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/users/login")
		controller := NewUserController(&mockError{}, validator.New())
		controller.Login()(context)

		type Response struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
			Data    interface{}
		}

		var resp Response
		json.Unmarshal([]byte(res.Body.Bytes()), &resp)

		assert.Equal(t, 401, resp.Code)
		assert.Equal(t, "username or password is wrong", resp.Message)
		assert.Nil(t, resp.Data)
	})
}

func TestGetbyID(t *testing.T) {
	t.Run("Status OK", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set("Authorization", "Bearer "+token)

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/users/profile")
		controller := NewUserController(&mockUser{}, validator.New())
		middleware.JWT([]byte("$4dm!n$"))(controller.GetbyID())(context)

		type Response struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
			Data    interface{}
		}

		var resp Response
		json.Unmarshal([]byte(res.Body.Bytes()), &resp)
		
		assert.Equal(t, 200, resp.Code)
		assert.Equal(t, "successfully", resp.Message)
		assert.Equal(t, map[string]interface {}(map[string]interface {}{"email":"dakasakti.id@gmail.com", "name":"Mahmuda Karima", "no_hp":"082278268513", "username":"dakasakti"}), resp.Data)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("Status OK", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"name":     "testing updated",
		})
		req := httptest.NewRequest(http.MethodPut, "/", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set("Authorization", "Bearer "+token)

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/users/profile")
		controller := NewUserController(&mockUser{}, validator.New())
		middleware.JWT([]byte("$4dm!n$"))(controller.Update())(context)

		type Response struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
			Data    interface{}
		}

		var resp Response
		json.Unmarshal([]byte(res.Body.Bytes()), &resp)

		assert.Equal(t, 200, resp.Code)
		assert.Equal(t, "successfully", resp.Message)
		assert.Equal(t, map[string]interface {}{"email":"", "name":"testing updated", "no_hp":"", "password":"", "updated_at":"0001-01-01T00:00:00Z"}, resp.Data)
	})

	t.Run("Status BadRequest Bind", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"name":     1,
		})

		req := httptest.NewRequest(http.MethodPut, "/", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set("Authorization", "Bearer "+token)

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/users/profile")
		controller := NewUserController(&mockUser{}, validator.New())
		middleware.JWT([]byte("$4dm!n$"))(controller.Update())(context)

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
		})

		req := httptest.NewRequest(http.MethodPut, "/", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set("Authorization", "Bearer "+token)

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/users/profile")
		controller := NewUserController(&mockUser{}, validator.New())
		middleware.JWT([]byte("$4dm!n$"))(controller.Update())(context)

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
			"name":     "testing updated",
			"no_hp":    "098765433212",
		})

		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set("Authorization", "Bearer "+token)

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/users/profile")
		controller := NewUserController(&mockError{}, validator.New())
		middleware.JWT([]byte("$4dm!n$"))(controller.Update())(context)

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

func TestDelete(t *testing.T) {
	t.Run("Status OK", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodDelete, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set("Authorization", "Bearer "+token)

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/users/profile")
		controller := NewUserController(&mockUser{}, validator.New())
		middleware.JWT([]byte("$4dm!n$"))(controller.Delete())(context)

		type Response struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
			Data    interface{}
		}

		var resp Response
		json.Unmarshal([]byte(res.Body.Bytes()), &resp)
		
		assert.Equal(t, 200, resp.Code)
		assert.Equal(t, "successfully", resp.Message)
		assert.Equal(t, map[string]interface {}{"deleted_at": "0001-01-01T00:00:00Z", "name":"testing"}, resp.Data)
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

func (m *mockUser) Login(username, password string) (response.InsertLogin, error) {
	return response.InsertLogin{
		ID: 1,
		Name: "testing",
		Token: "",
	}, nil
}

func (m *mockUser) GetbyID(id uint) response.GetUser {
	return response.GetUser{
		Name: "Mahmuda Karima",
    	Username: "dakasakti",
    	NoHp: "082278268513",
    	Email: "dakasakti.id@gmail.com",
	}
}

func (m *mockUser) Update(user_id uint, user *entity.User) (response.UpdateUser, error) {
	return response.UpdateUser{
		Name: user.Name,
		Email: user.Email,
		NoHp: user.NoHp,
		Password: user.Password,
		UpdatedAt: user.UpdatedAt,
	}, nil
}

func (m *mockUser) Delete(user_id uint) response.DeleteUser {
	return response.DeleteUser{
		Name: "testing",
		DeletedAt: time.Time{},
	}
}

type mockError struct {}

func (m *mockError) Insert(user *entity.User) (response.InsertUser, error) {
	return response.InsertUser{}, errors.New("field '...' : duplicate!")
}

func (m *mockError) Login(username, password string) (response.InsertLogin, error) {
	return response.InsertLogin{}, errors.New("username or password is wrong")
}

func (m *mockError) GetbyID(id uint) response.GetUser {
	return response.GetUser{}
}

func (m *mockError) Update(user_id uint, user *entity.User) (response.UpdateUser, error) {
	return response.UpdateUser{}, errors.New("field '...' : duplicate!")
}

func (m *mockError) Delete(user_id uint) response.DeleteUser {
	return response.DeleteUser{}
}