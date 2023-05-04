package delivery

import (
	"cleanarc/features/categories"
	"cleanarc/utils/helpers"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	Service categories.ServiceInterface
}

func New(s categories.ServiceInterface) *Handler {
	return &Handler{
		Service: s,
	}
}

func (h *Handler) GetAll(c echo.Context) error {
	dataCore, err := h.Service.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.ResponseFail("error read data"))
	}
	dataResponse := ListCoreToResponse(dataCore)
	return c.JSON(http.StatusOK, helpers.ResponseSuccess("-", dataResponse))
}

func (t *Handler) GetById(c echo.Context) error {
	_id, _ := strconv.Atoi(c.Param("id"))
	id := uint(_id)
	dataCore, err := t.Service.GetById(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, helpers.ResponseFail("data not found"))
	}

	dataResponse := CoreToResponse(dataCore)
	return c.JSON(http.StatusOK, helpers.ResponseSuccess("-", dataResponse))
}

func (t *Handler) Create(c echo.Context) error {
	var request Request
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, helpers.ResponseFail("error bind data"))
	}

	dataResponse, err := t.Service.Create(RequestToCore(&request))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.ResponseFail(err.Error()))
	}

	return c.JSON(http.StatusCreated, helpers.ResponseSuccess("Create Data Success", CoreToResponse(dataResponse)))
}

func (t *Handler) Update(c echo.Context) error {
	var request Request
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, helpers.ResponseFail("error bind data"))
	}

	id, _ := strconv.Atoi(c.Param("id"))
	dataCore, err := t.Service.Update(RequestToCore(&request), uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, helpers.ResponseFail(err.Error()))
	}

	return c.JSON(http.StatusOK, helpers.ResponseSuccess("Update Data Success", CoreToResponse(dataCore)))
}

func (t *Handler) Delete(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	if err := t.Service.Delete(uint(id)); err != nil {
		return c.JSON(http.StatusNotFound, helpers.ResponseFail(err.Error()))
	}

	return c.JSON(http.StatusOK, helpers.ResponseSuccess("Delete Data Success", nil))
}
