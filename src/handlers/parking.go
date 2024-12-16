package handlers

import (
	"back/models"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func getByID(c echo.Context) (*models.Parking, error) {
	id := c.Param("id")
	col, err := models.FindParking(&models.Parking{ID: id})
	if err != nil {
		return nil, err
	}
	return col, nil
}

func GetParkings(c echo.Context) error {
	parking := models.GetParkings()
	return c.JSON(http.StatusOK, parking)
}

func GetParking(c echo.Context) error {
	parking, err := getByID(c)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, parking)
}

// add parking
func AddParking(c echo.Context) error {
	parking := new(models.Parking)
	if err := c.Bind(parking); err != nil {
		return err
	}

	if parking.ID == "" || parking.Pubname == "" || parking.Status == "" {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "invalid id, pubname, status",
		}
	}

	_, err := models.FindParking(&models.Parking{ID: parking.ID})
	if err == nil {
		return &echo.HTTPError{
			Code:    http.StatusConflict,
			Message: "this pubname is already exist",
		}
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return &echo.HTTPError{
			Code: http.StatusInternalServerError,
			Message: err,
		}
	}

	err = models.AddParking(parking)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, Feedback{Request: "post", Message: "successed post record"})
}

// update parking status
func UpdateParking(c echo.Context) error {
	_, err := getByID(c)
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "invalid id",
		}
	}

	parking := new(models.Parking)
	if err := c.Bind(parking); err != nil {
		return err
	}

	if parking.Status == "" {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "invalid status",
		}
	}

	err = models.UpdateParking(parking)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, Feedback{Request: "put", Message: "successed put record"})
}

// delete parking record
func DeleteParking(c echo.Context) error {
	parking, err := getByID(c)
	if err != nil {
		return &echo.HTTPError{
			Code: http.StatusBadRequest,
			Message: "invalid id",
		}
	}

	err = models.DeleteParking(parking)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, Feedback{Request: "delete", Message: "successed delete record"})
}