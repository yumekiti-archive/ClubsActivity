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
	ID        int    `json:"id"`
	UID       string `json:"uid"`
	Name      string `json:"name"`
	Class     string `json:"class"`
	Icon      string `json:"icon"`
	Readme    string `json:"readme"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type responseUserWithClub struct {
	responseUser
	AffiliatedClubs []responseClub `json:"affiliated_clubs"`
}

func (h *userHandler) FindMe() echo.HandlerFunc {
	return func(c echo.Context) error {
		user, err := h.userUsecase.FindMe()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}

		var responseClubs []responseClub
		for _, club := range user.AffiliatedClubs {
			responseClubs = append(responseClubs, responseClub{
				ID:          club.ClubID,
				Name:        club.ClubName,
				Description: club.ClubDescription,
				Category:    club.ClubCategory,
				CreatedAt:   club.CreatedAt.Format("2006-01-02 15:04:05"),
				UpdatedAt:   club.UpdatedAt.Format("2006-01-02 15:04:05"),
			})
		}

		response := responseUserWithClub{
			responseUser: responseUser{
				ID:        user.UserID,
				UID:       user.UserUID,
				Name:      user.UserName,
				Class:     user.UserClass,
				Icon:      user.UserIcon,
				Readme:    user.UserReadme,
				CreatedAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
				UpdatedAt: user.UpdatedAt.Format("2006-01-02 15:04:05"),
			},
			AffiliatedClubs: responseClubs,
		}

		return c.JSON(http.StatusOK, response)
	}
}
