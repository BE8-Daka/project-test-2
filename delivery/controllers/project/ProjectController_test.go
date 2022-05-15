package project

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
		context.SetPath("/projects")
		controller := NewProjectController(&mockProject{}, validator.New())
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
		context.SetPath("/projects")
		controller := NewProjectController(&mockProject{}, validator.New())
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
		context.SetPath("/projects")
		controller := NewProjectController(&mockProject{}, validator.New())
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

	t.Run("Status BadRequest Duplicate", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"name":     "testing",
		})

		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set("Authorization", "Bearer "+token)

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/projects")
		controller := NewProjectController(&mockError{}, validator.New())
		middleware.JWT([]byte("$4dm!n$"))(controller.Insert())(context)

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

func TestGetAll(t *testing.T) {
	t.Run("Status OK", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set("Authorization", "Bearer "+token)

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/projects")
		controller := NewProjectController(&mockProject{}, validator.New())
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
		assert.Equal(t, []interface {}([]interface {}{map[string]interface {}{"id":float64(1), "name":"testing"}, map[string]interface {}{"id":float64(2), "name":"testing 2"}}), resp.Data)
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
		context.SetPath("/projects")
		controller := NewProjectController(&mockProject{}, validator.New())
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
		assert.Equal(t, map[string]interface {}{"name":"project", "updated_at":"0001-01-01T00:00:00Z"}, resp.Data)
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
		context.SetPath("/projects")
		controller := NewProjectController(&mockError{}, validator.New())
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
		context.SetPath("/projects")
		controller := NewProjectController(&mockProject{}, validator.New())
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
		context.SetPath("/projects")
		controller := NewProjectController(&mockProject{}, validator.New())
		middleware.JWT([]byte("$4dm!n$"))(controller.Update())(context)

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

func TestDelete(t *testing.T) {
	t.Run("Status OK", func(t *testing.T) {
		e := echo.New()
		
		req := httptest.NewRequest(http.MethodDelete, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set("Authorization", "Bearer "+token)

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/projects")
		controller := NewProjectController(&mockProject{}, validator.New())
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
		assert.Equal(t,  map[string]interface {}{"deleted_at":interface {}(nil), "name":"project"}, resp.Data)
	})

	t.Run("Status Forbidden", func(t *testing.T) {
		e := echo.New()
		
		req := httptest.NewRequest(http.MethodDelete, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set("Authorization", "Bearer "+token)

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/projects")
		controller := NewProjectController(&mockError{}, validator.New())
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

type mockProject struct {}

func (m *mockProject) Insert(project *entity.Project) (response.InsertProject, error) {
	return response.InsertProject{
		Name: project.Name,
		CreatedAt: project.CreatedAt,
	}, nil
}

func (m *mockProject) GetAll(user_id uint) []response.Project {
	return []response.Project{
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

func (m *mockProject) Update(id uint, project *entity.Project) response.UpdateProject {
	return response.UpdateProject{
		Name: project.Name,
		UpdatedAt: project.UpdatedAt,
	}
}

func (m *mockProject) CheckExist(id, user_id uint) bool {
	return true
}

func (m *mockProject) Delete(id uint) response.DeleteProject {
	return response.DeleteProject{
		Name: "project",
		DeletedAt: gorm.DeletedAt{},
	}
}

type mockError struct {}

func (m *mockError) Insert(project *entity.Project) (response.InsertProject, error) {
	return response.InsertProject{}, errors.New("field '...' : duplicate!")
}

func (m *mockError) GetAll(user_id uint) []response.Project {
	return []response.Project{}
}

func (m *mockError) Update(id uint, project *entity.Project) response.UpdateProject {
	return response.UpdateProject{
		Name: project.Name,
		UpdatedAt: project.UpdatedAt,
	}
}

func (m *mockError) CheckExist(id, user_id uint) bool {
	return false
}

func (m *mockError) Delete(id uint) response.DeleteProject {
	return response.DeleteProject{}
}