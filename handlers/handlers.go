package handlers

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/indmind/simpletodo/models"

	echo "gopkg.in/echo.v3"
)

// H arbitary JSON for returning data :/
type H map[string]interface{}

// GetTasks get all stored task from database model
func GetTasks(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, models.GetTasks(db))
	}
}

// PutTask input task data into database model
func PutTask(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var task models.Task

		c.Bind(&task)

		id, err := models.PutTask(db, task.Name)

		if err != nil {
			return err
		}

		return c.JSON(http.StatusCreated, H{
			"created": id,
		})
	}
}

// DeleteTask delete data using database model
func DeleteTask(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))

		_, err := models.DeleteTask(db, id)

		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, H{
			"deleted": id,
		})
	}
}
