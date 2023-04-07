package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"

	"api/domain"
	"api/usecase"
)

type UserHandler interface {
	FindMe() echo.HandlerFunc
	Store() echo.HandlerFunc
}

type userHandler struct {
	userUsecase usecase.UserUsecase
}

func NewUserHandler(userUsecase usecase.UserUsecase) UserHandler {
	return &userHandler{userUsecase: userUsecase}
}

type requestUser struct {
	Name   string `json:"name"`
	Class  string `json:"class"`
	Icon   string `json:"icon"`
	Readme string `json:"readme"`
}

type responseUser struct {
	ID        int    `json:"id"`
	UID       string `json:"uid"`
	Name      string `json:"name"`
	Class     string `json:"class"`
	Icon      string `json:"icon"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type responseUserWithReadme struct {
	responseUser
	Readme string `json:"readme"`
}

type responseUserWithClub struct {
	responseUser
	Clubs []responseClub `json:"affiliated_clubs"`
}

func (h *userHandler) Store() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req requestUser
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		user, err := h.userUsecase.Store(&domain.User{
			Name:  req.Name,
			Class: req.Class,
			Icon:  req.Icon,
		})
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}

		response := responseUser{
			ID:        user.ID,
			UID:       user.UID,
			Name:      user.Name,
			Class:     user.Class,
			Icon:      user.Icon,
			CreatedAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: user.UpdatedAt.Format("2006-01-02 15:04:05"),
		}

		return c.JSON(http.StatusCreated, response)
	}
}

func (h *userHandler) FindMe() echo.HandlerFunc {
	return func(c echo.Context) error {
		user, err := h.userUsecase.FindMe(1)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}

		var responseClubs []responseClub
		for _, club := range user.Clubs {
			responseClubs = append(responseClubs, responseClub{
				ID:          club.ID,
				Name:        club.Name,
				Description: club.Description,
				Category:    club.Category,
				CreatedAt:   club.CreatedAt.Format("2006-01-02 15:04:05"),
				UpdatedAt:   club.UpdatedAt.Format("2006-01-02 15:04:05"),
			})
		}

		response := responseUserWithClub{
			responseUser: responseUser{
				ID:        user.ID,
				UID:       user.UID,
				Name:      user.Name,
				Class:     user.Class,
				Icon:      user.Icon,
				CreatedAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
				UpdatedAt: user.UpdatedAt.Format("2006-01-02 15:04:05"),
			},
			Clubs: responseClubs,
		}

		return c.JSON(http.StatusOK, response)
	}
}
