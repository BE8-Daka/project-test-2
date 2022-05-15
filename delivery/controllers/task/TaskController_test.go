package task

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"project-test/delivery/middlewares"
	"project-test/delivery/views/response"
	"project-test/entity"
	"strings"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/stretchr/testify/assert"
)

var (
	token string
)

func TestCreateToken(t *testing.T) {
	t.Run("Create Token", func(t *testing.T) {
		token, _ = middlewares.CreateToken(1)
	})
}

func TestInsert(t *testing.T) {
	t.Run("Status OK", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"name": "project",
		})
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set("Authorization", "Bearer "+token)

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/tasks")
		controller := NewTaskController(&mockTask{}, validator.New())
		middleware.JWT([]byte("$4dm!n$"))(controller.Insert())(context)

		type Response struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
			Data    interface{}
		}

		var resp Response
		json.Unmarshal([]byte(res.Body.Bytes()), &resp)
		fmt.Println(res.Body.String())

		assert.Equal(t, 201, resp.Code)
		assert.Equal(t, "successfully created", resp.Message)
		assert.Equal(t, map[string]interface {}{"created_at":"0001-01-01T00:00:00Z", "name":"project"}, resp.Data)
	})

	t.Run("Status BadRequest Bind", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"name":     1,
		})

		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set("Authorization", "Bearer "+token)

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/tasks")
		controller := NewTaskController(&mockTask{}, validator.New())
		middleware.JWT([]byte("$4dm!n$"))(controller.Insert())(context)

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
			"name": "",
		})

		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set("Authorization", "Bearer "+token)

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/tasks")
		controller := NewTaskController(&mockTask{}, validator.New())
		middleware.JWT([]byte("$4dm!n$"))(controller.Insert())(context)

		type Response struct {
			Code    int      `json:"code"`
			Message []string `json:"message"`
			Data    interface{}
		}

		var resp Response
		json.Unmarshal([]byte(res.Body.Bytes()), &resp)

		assert.Equal(t, 400, resp.Code)
		assert.Equal(t, []string{"field Name : required"}, resp.Message)
		assert.Nil(t, resp.Data)
	})
}

type mockTask struct {}

func (m *mockTask) Insert(task *entity.Task) response.InsertTask {
	return response.InsertTask{
		Name: task.Name,
		CreatedAt: task.CreatedAt,
	}
}