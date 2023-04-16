package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"

	"api/usecase"
)

type ActivityHandler interface {
	FindByClubID() echo.HandlerFunc
}

type activityHandler struct {
	activityUsecase usecase.ActivityUsecase
}

func NewActivityHandler(au usecase.ActivityUsecase) ActivityHandler {
	return &activityHandler{activityUsecase: au}
}

type responseActivity struct {
	ID        uint           `json:"id"`
	Place     string         `json:"place"`
	Detail    string         `json:"detail"`
	Users     []responseUser `json:"users"`
	CreatedAt string         `json:"created_at"`
	UpdatedAt string         `json:"updated_at"`
}

// @Summary クラブの活動を取得
// @Router /v1/clubs/{id}/activities [get]
// @Param id path int true "Club ID"
func (h *activityHandler) FindByClubID() echo.HandlerFunc {
	return func(c echo.Context) error {
		clubID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		activities, err := h.activityUsecase.FindByClubID(clubID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}

		var responseActivities []responseActivity
		for _, activity := range activities {
			var responseUsers []responseUser
			for _, user := range activity.Users {
				responseUsers = append(responseUsers, responseUser{
					ID:        user.ID,
					UID:       user.UID,
					Name:      user.Name,
					Class:     user.Class,
					Icon:      user.Icon,
					CreatedAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
					UpdatedAt: user.UpdatedAt.Format("2006-01-02 15:04:05"),
				})
			}

			responseActivities = append(responseActivities, responseActivity{
				ID:        activity.ID,
				Place:     activity.Place,
				Detail:    activity.Detail,
				Users:     responseUsers,
				CreatedAt: activity.CreatedAt.Format("2006-01-02 15:04:05"),
				UpdatedAt: activity.UpdatedAt.Format("2006-01-02 15:04:05"),
			})
		}

		return c.JSON(http.StatusOK, responseActivities)
	}
}
