package users

import (
	"net/http"

	"github.com/bhagas/go-svc-echo/config"
	"github.com/labstack/echo/v4"
)

type (
	user struct {
		ID        string `json:"id"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
	}
)

var (
	conf *config.Config
)

func Pasang(group *echo.Group, cfg config.Config) {
	conf = &cfg
	group.GET("", GetAllUsers)

}

func GetAllUsers(c echo.Context) error {
	var users []user
	config := *conf
	config.DB().Master().Raw("SELECT * FROM users").Scan(&users)

	return c.JSON(http.StatusOK, users)
}
