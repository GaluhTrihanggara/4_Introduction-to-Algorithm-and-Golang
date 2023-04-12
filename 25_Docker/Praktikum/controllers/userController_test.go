package controllers

import (
	"Praktikum/config"
	"Praktikum/models"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestLoginUserController(t *testing.T) {
	e := echo.New()

	userJSON := `{"email": "galuhhanggara@gmail.com", "password": "12345"}`
	req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, LoginUserController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), "success login")
	}
}

func TestGetUserDetailControllers(t *testing.T) {
	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/users/5", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("5")

	if assert.NoError(t, GetUserDetailControllers(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), "success")
	}
}

func TestGetUsersController(t *testing.T) {
	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, GetUsersController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), "success get all users")
	}
}

func TestGetUserController(t *testing.T) {
	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/users/4", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("4")

	if assert.NoError(t, GetUserController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), "success get user by id")
	}
}

func TestCreateUserController(t *testing.T) {
	e := echo.New()

	userJSON := `{"name": "kaka", "email": "kaka123@gmail.com", "password": "23145"}`
	req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, CreateUserController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), "success create new user")

		// parse response body to get created user
		var res struct {
			User struct {
				Name  string `json:"name"`
				Email string `json:"email"`
			} `json:"user"`
		}
		err := json.Unmarshal(rec.Body.Bytes(), &res)
		if assert.NoError(t, err) {
			assert.Equal(t, "kaka", res.User.Name)
			assert.Equal(t, "kaka123@gmail.com", res.User.Email)
		}
	}
}

func TestDeleteUserController(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Test case 1: Delete user with valid ID
	// Create a user to delete
	user := models.User{Name: "kaka", Email: "kaka123@gmail.com", Password: "23145"}
	config.DB.Create(&user)
	// Set the ID of the user to delete
	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(int(user.ID)))
	// Call the DeleteUserController function
	err := DeleteUserController(c)
	assert.NoError(t, err)
	// Check if the response code is 200 OK
	assert.Equal(t, http.StatusOK, rec.Code)
	// Check if the response body contains "berhasil menghapus data"
	assert.Contains(t, rec.Body.String(), "berhasil menghapus data")
	// Check if the user is deleted from the database
	var deletedUser models.User
	err = config.DB.Where("id = ?", user.ID).First(&deletedUser).Error
	assert.Equal(t, gorm.ErrRecordNotFound, err)

	// Test case 2: Delete user with invalid ID
	// Set an invalid user ID (ID of a non-existent user)
	c.SetParamValues("999")
	// Call the DeleteUserController function
	err = DeleteUserController(c)
	assert.NoError(t, err)
	// Check if the response code is 200 OK
	assert.Equal(t, http.StatusOK, rec.Code)
	// Check if the response body contains "user not available"
	assert.Contains(t, rec.Body.String(), "user not available")
}

func TestUpdateUserController(t *testing.T) {
	// setup
	config.InitDB()
	e := echo.New()
	req := httptest.NewRequest(http.MethodPut, "/users/25", strings.NewReader(`{"name":"Diky", "email":"dikyrahman@gmail.com", "password":"34567"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("25")

	// test UpdateUserController
	if assert.NoError(t, UpdateUserController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		// check response
		var response map[string]interface{}
		err := json.Unmarshal([]byte(rec.Body.String()), &response)
		if assert.NoError(t, err) {
			assert.Equal(t, "berhasil update data", response["status"])
			assert.NotNil(t, response["user"])
		}

		// check updated user
		var user models.User
		if err := config.DB.Where("id = ?", 25).First(&user).Error; err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, "Diky", user.Name)
		assert.Equal(t, "dikyrahman@gmail.com", user.Email)
		assert.Equal(t, "34567", user.Password)
	}
}
