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
	UID    int    `json:"uid" example:"220000"`
	Name   string `json:"name" example:"ぴよ太郎"`
	Class  string `json:"class" example:"IE3A"`
	Icon   string `json:"icon" example:""`
	Readme string `json:"readme" example:""`
}

type responseUser struct {
	ID        int    `json:"id"`
	UID       int    `json:"uid"`
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
	Clubs []responseClub `json:"clubs"`
}

// @Summary ユーザーを作成する
// @Router /v1/users [post]
// @Param user body requestUser true "ユーザー情報"
func (h *userHandler) Store() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req requestUser
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		user, err := h.userUsecase.Store(&domain.User{
			UID:   req.UID,
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

// @Summary 自分のユーザー情報を取得する
// @Router /v1/me [get]
func (h *userHandler) FindMe() echo.HandlerFunc {
	return func(c echo.Context) error {
		user, err := h.userUsecase.FindMe(220001)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}

		var responseClubs []responseClub
		for _, club := range user.Clubs {
			responseClubs = append(responseClubs, responseClub{
				ID:        club.ID,
				Name:      club.Name,
				Category:  club.Category,
				CreatedAt: club.CreatedAt.Format("2006-01-02 15:04:05"),
				UpdatedAt: club.UpdatedAt.Format("2006-01-02 15:04:05"),
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
