package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"

	"api/usecase"
)

type UserHandler interface {
	FindMe() echo.HandlerFunc
}

type userHandler struct {
	userUsecase usecase.UserUsecase
}

func NewUserHandler(userUsecase usecase.UserUsecase) UserHandler {
	return &userHandler{userUsecase: userUsecase}
}

type requestUser struct{}

type responseUser struct {
	ID               int    `json:"id"`
	UID              string `json:"uid"`
	Name             string `json:"name"`
	Class            string `json:"class"`
	Icon             string `json:"icon"`
	CreatedAt        string `json:"created_at"`
	UpdatedAt        string `json:"updated_at"`
	AffiliatedClubID []int  `json:"affiliated_club_id"`
}

func (h *userHandler) FindMe() echo.HandlerFunc {
	return func(c echo.Context) error {
		user, err := h.userUsecase.FindMe()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}

		return c.JSON(http.StatusOK, responseUser{
			ID:               user.UserID,
			UID:              user.UserUID,
			Name:             user.UserName,
			Class:            user.UserClass,
			Icon:             user.UserIcon,
			CreatedAt:        user.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:        user.UpdatedAt.Format("2006-01-02 15:04:05"),
			AffiliatedClubID: user.AffiliatedClubID,
		})
	}
}
