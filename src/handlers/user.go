package handlers

import (
	"back/models"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func GetUsers(c echo.Context) error {
	users, err := models.AllUsers()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, users)
}

func GetUser(c echo.Context) error {
	name := c.Param("name")

	/* var u *models.User
	if err := c.Bind(u); err != nil {
		return &echo.HTTPError{
			Code: http.StatusBadRequest,
			Message: "invalid id",
		}
	} */

	user, err := models.FindUser(&models.User{Name: name})
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, user)
}

func AddUser(c echo.Context) error {
	user := new(models.User)
	if err := c.Bind(user); err != nil {
		return err
	}

	if user.Name == "" || user.Email == "" {
		return &echo.HTTPError{
			Code: http.StatusBadRequest,
			Message: "invalid info",
		}
	}

	_, err := models.FindUser(&models.User{Email: user.Email})
	if err == nil{
		return &echo.HTTPError{
			Code: http.StatusBadRequest,
			Message: "this email is already exist",
		}
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return &echo.HTTPError{
			Code: http.StatusInternalServerError,
			Message: err,
		}
	}

	if err = models.CreateUser(user); err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, Feedback{Request: "post", Message: "successed to add user"})
}

func DeleteUser(c echo.Context) error {
	name := c.Param("name")

	user, err := models.FindUser(&models.User{Name: name})
	if err != nil {
		return &echo.HTTPError{
			Code: http.StatusBadRequest,
			Message: "invalid name",
		}
	}

	if err = models.DeleteUser(user); err != nil {
		return err
	}
	
	return c.JSON(http.StatusOK, Feedback{Request: "delete", Message: "successed to delete record"})
}