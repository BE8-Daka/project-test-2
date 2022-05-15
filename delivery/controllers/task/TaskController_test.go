package task

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
	"gorm.io/gorm"
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

func TestGetAll(t *testing.T) {
	t.Run("Status OK", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set("Authorization", "Bearer "+token)

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/tasks")
		controller := NewTaskController(&mockTask{}, validator.New())
		middleware.JWT([]byte("$4dm!n$"))(controller.GetAll())(context)

		type Response struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
			Data    interface{}
		}

		var resp Response
		json.Unmarshal([]byte(res.Body.Bytes()), &resp)
		
		assert.Equal(t, 200, resp.Code)
		assert.Equal(t, "successfully get all data", resp.Message)
		assert.Equal(t, []interface {}([]interface {}{map[string]interface {}{"id":float64(1), "name":"testing", "project_id":float64(1)}, map[string]interface {}{"id":float64(2), "name":"testing 2",  "project_id":float64(1)}}), resp.Data)
	})

	t.Run("Status NotFound", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set("Authorization", "Bearer "+token)

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/tasks")
		controller := NewTaskController(&mockError{}, validator.New())
		middleware.JWT([]byte("$4dm!n$"))(controller.GetAll())(context)

		type Response struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
			Data    interface{}
		}

		var resp Response
		json.Unmarshal([]byte(res.Body.Bytes()), &resp)
		
		assert.Equal(t, 404, resp.Code)
		assert.Equal(t, "tasks not found", resp.Message)
		assert.Nil(t, resp.Data)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("Status OK", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"name": "project",
		})
		req := httptest.NewRequest(http.MethodPut, "/", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set("Authorization", "Bearer "+token)

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/task")
		controller := NewTaskController(&mockTask{}, validator.New())
		middleware.JWT([]byte("$4dm!n$"))(controller.Update())(context)

		type Response struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
			Data    interface{}
		}

		var resp Response
		json.Unmarshal([]byte(res.Body.Bytes()), &resp)

		assert.Equal(t, 200, resp.Code)
		assert.Equal(t, "successfully updated", resp.Message)
		assert.Equal(t, map[string]interface {}{"name":"testing", "updated_at":"0001-01-01T00:00:00Z"}, resp.Data)
	})

	t.Run("Status Forbidden", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"name": "project",
		})
		req := httptest.NewRequest(http.MethodPut, "/", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set("Authorization", "Bearer "+token)

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/tasks")
		controller := NewTaskController(&mockError{}, validator.New())
		middleware.JWT([]byte("$4dm!n$"))(controller.Update())(context)

		type Response struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
			Data    interface{}
		}

		var resp Response
		json.Unmarshal([]byte(res.Body.Bytes()), &resp)

		assert.Equal(t, 403, resp.Code)
		assert.Equal(t, "you are not allowed to access this resource", resp.Message)
		assert.Nil(t, resp.Data)
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
		context.SetPath("/tasks")
		controller := NewTaskController(&mockTask{}, validator.New())
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
			"name": "",
		})

		req := httptest.NewRequest(http.MethodPut, "/", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set("Authorization", "Bearer "+token)

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/tasks")
		controller := NewTaskController(&mockErrorRequired{}, validator.New())
		middleware.JWT([]byte("$4dm!n$"))(controller.Update())(context)

		type Response struct {
			Code    int      `json:"code"`
			Message string `json:"message"`
			Data    interface{}
		}

		var resp Response
		json.Unmarshal([]byte(res.Body.Bytes()), &resp)

		assert.Equal(t, 400, resp.Code)
		assert.Equal(t, "name or project_id is required", resp.Message)
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
		context.SetPath("/tasks")
		controller := NewTaskController(&mockTask{}, validator.New())
		middleware.JWT([]byte("$4dm!n$"))(controller.Delete())(context)

		type Response struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
			Data    interface{}
		}

		var resp Response
		json.Unmarshal([]byte(res.Body.Bytes()), &resp)

		assert.Equal(t, 200, resp.Code)
		assert.Equal(t, "successfully deleted", resp.Message)
		assert.Equal(t,   map[string]interface {}{"deleted_at":interface {}(nil), "name":"testing"}, resp.Data)
	})

	t.Run("Status Forbidden", func(t *testing.T) {
		e := echo.New()
		
		req := httptest.NewRequest(http.MethodDelete, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set("Authorization", "Bearer "+token)

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/tasks")
		controller := NewTaskController(&mockError{}, validator.New())
		middleware.JWT([]byte("$4dm!n$"))(controller.Delete())(context)

		type Response struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
			Data    interface{}
		}

		var resp Response
		json.Unmarshal([]byte(res.Body.Bytes()), &resp)

		assert.Equal(t, 403, resp.Code)
		assert.Equal(t, "you are not allowed to access this resource", resp.Message)
		assert.Nil(t, resp.Data)
	})
}

func TestUpdateStatus(t *testing.T) {
	t.Run("Status OK Completed", func(t *testing.T) {
		e := echo.New()
		
		req := httptest.NewRequest(http.MethodPost, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set("Authorization", "Bearer "+token)

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/tasks/:id/completed")
		controller := NewTaskController(&mockTask{}, validator.New())
		middleware.JWT([]byte("$4dm!n$"))(controller.UpdateStatus())(context)

		type Response struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
			Data    interface{}
		}

		var resp Response
		json.Unmarshal([]byte(res.Body.Bytes()), &resp)

		assert.Equal(t, 200, resp.Code)
		assert.Equal(t, "successfully task completed", resp.Message)
		assert.Equal(t, map[string]interface {}{"name":"testing", "updated_at":"0001-01-01T00:00:00Z"}, resp.Data)
	})

	t.Run("Status OK Reopen", func(t *testing.T) {
		e := echo.New()
		
		req := httptest.NewRequest(http.MethodPost, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set("Authorization", "Bearer "+token)

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/tasks/:id/reopen")
		controller := NewTaskController(&mockTask{}, validator.New())
		middleware.JWT([]byte("$4dm!n$"))(controller.UpdateStatus())(context)

		type Response struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
			Data    interface{}
		}

		var resp Response
		json.Unmarshal([]byte(res.Body.Bytes()), &resp)

		assert.Equal(t, 200, resp.Code)
		assert.Equal(t, "successfully task reopen", resp.Message)
		assert.Equal(t, map[string]interface {}{"name":"testing", "updated_at":"0001-01-01T00:00:00Z"}, resp.Data)
	})

	t.Run("Status Forbidden", func(t *testing.T) {
		e := echo.New()
		
		req := httptest.NewRequest(http.MethodPost, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set("Authorization", "Bearer "+token)

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/tasks/:id/completed")
		controller := NewTaskController(&mockError{}, validator.New())
		middleware.JWT([]byte("$4dm!n$"))(controller.UpdateStatus())(context)

		type Response struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
			Data    interface{}
		}

		var resp Response
		json.Unmarshal([]byte(res.Body.Bytes()), &resp)

		assert.Equal(t, 403, resp.Code)
		assert.Equal(t, "you are not allowed to access this resource", resp.Message)
		assert.Nil(t, resp.Data)
	})
}

func TestGetTaskbyID(t *testing.T) {
	t.Run("Status OK", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set("Authorization", "Bearer "+token)

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/tasks")
		controller := NewTaskController(&mockTask{}, validator.New())
		middleware.JWT([]byte("$4dm!n$"))(controller.GetTaskbyID())(context)

		type Response struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
			Data    interface{}
		}

		var resp Response
		json.Unmarshal([]byte(res.Body.Bytes()), &resp)
		
		assert.Equal(t, 200, resp.Code)
		assert.Equal(t, "successfully get all data", resp.Message)
		assert.Equal(t, []interface {}{map[string]interface {}{"id":float64(1), "name":"testing"}, map[string]interface {}{"id":float64(2), "name":"testing 2"}}, resp.Data)
	})

	t.Run("Status NotFound", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set("Authorization", "Bearer "+token)

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/tasks")
		controller := NewTaskController(&mockError{}, validator.New())
		middleware.JWT([]byte("$4dm!n$"))(controller.GetTaskbyID())(context)

		type Response struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
			Data    interface{}
		}

		var resp Response
		json.Unmarshal([]byte(res.Body.Bytes()), &resp)
		
		assert.Equal(t, 404, resp.Code)
		assert.Equal(t, "tasks not found", resp.Message)
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

func (m *mockTask) GetAll(user_id uint) []response.Task {
	return []response.Task{
		{
			ID: 1,
			Name: "testing",
			ProjectID: 1,
		},
		{
			ID: 2,
			Name: "testing 2",
			ProjectID: 1,
		},
	}
}

func (m *mockTask) Update(id uint, task *entity.Task) (response.UpdateTask, error) {
	return response.UpdateTask{
		Name: "testing",
		UpdatedAt: task.UpdatedAt,
	}, nil
}

func (m *mockTask) CheckExist(id, user_id uint) bool {
	return true
}

func (m *mockTask) Delete(id uint) response.DeleteTask {
	return response.DeleteTask{
		Name: "testing",
		DeletedAt: gorm.DeletedAt{},
	}
}

func (m *mockTask) UpdateStatus(id uint, task *map[string]interface{}) response.UpdateTask {
	return response.UpdateTask{
		Name: "testing",
		UpdatedAt: time.Time{},
	} 
}

func (m *mockTask) GetTaskbyID(project_id, user_id uint) []response.TaskID {
	return []response.TaskID{
		{
			ID: 1,
			Name: "testing",
		},
		{
			ID: 2,
			Name: "testing 2",
		},
	}
}

type mockErrorRequired struct {}

func (m *mockErrorRequired) Insert(task *entity.Task) response.InsertTask {
	return response.InsertTask{}
}

func (m *mockErrorRequired) GetAll(user_id uint) []response.Task {
	return []response.Task{}
}

func (m *mockErrorRequired) Update(id uint, task *entity.Task) (response.UpdateTask, error) {
	return response.UpdateTask{}, errors.New("name or project_id is required")
}

func (m *mockErrorRequired) CheckExist(id, user_id uint) bool {
	return true
}

func (m *mockErrorRequired) Delete(id uint) response.DeleteTask {
	return response.DeleteTask{}
}

func (m *mockErrorRequired) UpdateStatus(id uint, task *map[string]interface{}) response.UpdateTask {
	return response.UpdateTask{} 
}

func (m *mockErrorRequired) GetTaskbyID(project_id, user_id uint) []response.TaskID {
	return []response.TaskID{}
}

type mockError struct {}

func (m *mockError) Insert(task *entity.Task) response.InsertTask {
	return response.InsertTask{}
}

func (m *mockError) GetAll(user_id uint) []response.Task {
	return []response.Task{}
}

func (m *mockError) Update(id uint, task *entity.Task) (response.UpdateTask, error) {
	return response.UpdateTask{}, errors.New("name or project_id is required")
}

func (m *mockError) CheckExist(id, user_id uint) bool {
	return false
}

func (m *mockError) Delete(id uint) response.DeleteTask {
	return response.DeleteTask{}
}

func (m *mockError) UpdateStatus(id uint, task *map[string]interface{}) response.UpdateTask {
	return response.UpdateTask{} 
}

func (m *mockError) GetTaskbyID(project_id, user_id uint) []response.TaskID {
	return []response.TaskID{}
}