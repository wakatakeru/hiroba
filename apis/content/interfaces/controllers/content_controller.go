package controllers

import (
	"net/http"
	"strconv"

	"github.com/wakatakeru/hiroba/apis/content/domain"
	"github.com/wakatakeru/hiroba/apis/content/interfaces/database"
	"github.com/wakatakeru/hiroba/apis/content/usecase"
)

type ContentController struct {
	Interactor usecase.ContentInteractor
	JWTHandler JWTHandler
}

func NewContentController(sqlHandler database.SqlHandler, jwtHandler JWTHandler) *ContentController {
	return &ContentController{
		Interactor: usecase.ContentInteractor{
			ContentRepository: &database.ContentRepository{
				SqlHandler: sqlHandler,
			},
		},
		JWTHandler: jwtHandler,
	}
}

func (controller *ContentController) Create(c Context) {
	_, err := controller.JWTHandler.Verify(c.GetHeader("Authorization"))
	if err != nil {
		c.JSON(http.StatusUnauthorized, nil)
		return
	}

	content := domain.Content{}
	err = c.Bind(&content)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	id, err := controller.Interactor.Add(content)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(http.StatusOK, id)
}

func (controller *ContentController) Show(c Context) {
	_, err := controller.JWTHandler.Verify(c.GetHeader("Authorization"))
	if err != nil {
		c.JSON(http.StatusUnauthorized, nil)
		return
	}

	id, _ := strconv.Atoi(c.Param("id"))
	content, err := controller.Interactor.Content(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(http.StatusOK, content)
}

func (controller *ContentController) SiteIndex(c Context) {
	_, err := controller.JWTHandler.Verify(c.GetHeader("Authorization"))
	if err != nil {
		c.JSON(http.StatusUnauthorized, nil)
		return
	}

	siteId, _ := strconv.Atoi(c.Param("site_id"))
	contents, err := controller.Interactor.SiteContents(siteId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(http.StatusOK, contents)
}

func (controller *ContentController) UserIndex(c Context) {
	_, err := controller.JWTHandler.Verify(c.GetHeader("Authorization"))
	if err != nil {
		c.JSON(http.StatusUnauthorized, nil)
		return
	}

	userId, _ := strconv.Atoi(c.Param("user_id"))
	contents, err := controller.Interactor.SiteContents(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(http.StatusOK, contents)
}
