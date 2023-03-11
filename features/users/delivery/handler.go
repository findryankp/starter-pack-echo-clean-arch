package delivery

import (
	"immersiveApp/features/users"
	"immersiveApp/utils/helpers"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	Service  users.UserServiceInterface
	validate *validator.Validate
}

func New(s users.UserServiceInterface) *UserHandler {
	return &UserHandler{
		Service: s,
	}
}

func (h *UserHandler) GetAll(c echo.Context) error {
	users, err := h.Service.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.ResponseFail(err.Error()))
	}

	return c.JSON(http.StatusOK, helpers.ResponseSuccess("-", ListUserEntityToUserResponse(users)))
}

func (h *UserHandler) GetById(c echo.Context) error {
	_id, _ := strconv.Atoi(c.Param("id"))
	id := uint(_id)

	user, err := h.Service.GetById(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.ResponseFail(err.Error()))
	}

	return c.JSON(http.StatusOK, helpers.ResponseSuccess("-", UserEntityToUserResponse(user)))
}

func (h *UserHandler) Create(c echo.Context) error {
	var formInput UserRequest
	if err := c.Bind(&formInput); err != nil {
		return c.JSON(http.StatusBadRequest, helpers.ResponseFail("error bind data"))
	}

	h.validate = validator.New()
	errValidate := h.validate.Struct(formInput)
	if errValidate != nil {
		return c.JSON(http.StatusBadRequest, helpers.ResponseFail(errValidate.Error()))
	}

	user, err := h.Service.Create(UserRequestToUserEntity(formInput))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.ResponseFail(err.Error()))
	}

	return c.JSON(http.StatusCreated, helpers.ResponseSuccess("Create Data Success", UserEntityToUserResponse(user)))
}

func (h *UserHandler) Update(c echo.Context) error {
	var formInput UserRequest
	if err := c.Bind(&formInput); err != nil {
		return c.JSON(http.StatusBadRequest, helpers.ResponseFail("error bind data"))
	}

	_id, _ := strconv.Atoi(c.Param("id"))
	id := uint(_id)

	_, err := h.Service.Update(UserRequestToUserEntity(formInput), id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.ResponseFail(err.Error()))
	}

	user, _ := h.Service.GetById(id)

	return c.JSON(http.StatusOK, helpers.ResponseSuccess("Update Data Success", UserEntityToUserResponse(user)))
}

func (h *UserHandler) Delete(c echo.Context) error {
	_id, _ := strconv.Atoi(c.Param("id"))
	id := uint(_id)

	if err := h.Service.Delete(id); err != nil {
		return c.JSON(http.StatusBadRequest, helpers.ResponseFail(err.Error()))
	}

	return c.JSON(http.StatusOK, helpers.ResponseSuccess("Delete Data Success", nil))
}
