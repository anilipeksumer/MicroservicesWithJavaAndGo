package service

import (
	"customer-service/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

// App represents the application.
type App struct {
	db *db.PostgresDB
}

// GetApp returns a new instance of App with the provided database.
func GetApp(db *db.PostgresDB) *App {
	return &App{
		db: db,
	}
}

func (a *App) PostHandler(c *gin.Context) {
	c.JSON(http.StatusOK, "ok")
}
