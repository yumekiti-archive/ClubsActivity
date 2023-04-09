package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"

	"api/usecase"
)

type ClubHandler interface {
	FindAll() echo.HandlerFunc
	FindByID() echo.HandlerFunc
}

type clubHandler struct {
	clubUsecase usecase.ClubUsecase
}

func NewClubHandler(cu usecase.ClubUsecase) ClubHandler {
	return &clubHandler{clubUsecase: cu}
}

type requestClub struct{}

type responseClub struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Readme    string `json:"readme"`
	Category  string `json:"category"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// @Summary クラブ一覧を取得
// @Router /v1/clubs [get]
func (h *clubHandler) FindAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		clubs, err := h.clubUsecase.FindAll()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}

		var responseClubs []responseClub
		for _, club := range clubs {
			responseClubs = append(responseClubs, responseClub{
				ID:        club.ID,
				Name:      club.Name,
				Readme:    club.Readme,
				Category:  club.Category,
				CreatedAt: club.CreatedAt.Format("2006-01-02 15:04:05"),
				UpdatedAt: club.UpdatedAt.Format("2006-01-02 15:04:05"),
			})
		}

		return c.JSON(http.StatusOK, responseClubs)
	}
}

// @Summary クラブを取得
// @Router /v1/clubs/{id} [get]
// @Param id path int true "Club ID"
func (h *clubHandler) FindByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}

		club, err := h.clubUsecase.FindByID(id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}

		return c.JSON(http.StatusOK, responseClub{
			ID:        club.ID,
			Name:      club.Name,
			Readme:    club.Readme,
			Category:  club.Category,
			CreatedAt: club.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: club.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}
}
