package handlers_auth_users

import (
	"net/http"

	"github.com/labstack/echo"
)

// FetchAuthUsers は `auth_users` テーブルからすべてのユーザー情報を取得する。
func (h *AuthUsersHandler) FetchAuthUsers(c echo.Context) error {
	authUsers, err := h.AuthUsersService.FetchAuthUsers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Error fetching auth_users",
		})
	}
	return c.JSON(http.StatusOK, authUsers)
}
