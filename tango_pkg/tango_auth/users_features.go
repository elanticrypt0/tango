package tango_auth

import (
	"net/http"
	"strconv"

	"tango_pkg/tango_errors"
	"tango_pkg/tango_validator"
	"tango_pkg/tangoapp"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

const usersPagination = false
const usersPaginationItemsPerPage = 15

func FindOneUser(ctx echo.Context, tapp *tangoapp.TangoApp) error {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	u := NewUser()
	user, err := u.FindOneByUID(tapp.DBAuth, id)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, err)
	}
	return ctx.JSON(http.StatusOK, user.ConvertToDTO())
}

func FindAllUsers(ctx echo.Context, tapp *tangoapp.TangoApp) error {
	var uBuf *[]User
	u := NewUser()

	if usersPagination == true {
		queryPage := ctx.Param("page")
		currentPage := 0
		if queryPage != "" {
			currentPage, _ = strconv.Atoi(queryPage)
		}

		// total de registros en la db
		// counter, _ := c.Count(tapp.App.DB.Auth)
		// pagination := pagination.NewPagination(currentPage,categoriesPaginationItemsPerPage,counter)

		uBuf, _ = u.FindAllPagination(tapp.DBAuth, usersPaginationItemsPerPage, currentPage)
	} else {
		uBuf, _ = u.FindAll(tapp.DBAuth)
	}

	return ctx.JSON(http.StatusOK, uBuf)

}

func ActivateUser(ctx echo.Context, tapp *tangoapp.TangoApp) error {
	// TODO
	code := ctx.Param("code")
	return ctx.String(http.StatusOK, code)
}

func CreateUser(ctx echo.Context, tapp *tangoapp.TangoApp) error {
	// get the incoming values
	uDTO := UserDTOCreate{}
	if err := ctx.Bind(&uDTO); err != nil {
		return ctx.JSON(http.StatusBadRequest, "")
	}

	u := NewUser()
	// checkquear the email is not used by other user
	email, _ := tango_validator.ValidateEmail(uDTO.Email)
	if u.IsEmailAvailable(tapp.DBAuth, email) {
		uBuf, err := u.Create(tapp.DBAuth, uDTO)

		if err != nil {
			return ctx.JSON(http.StatusBadRequest, err)
		}

		return ctx.JSON(http.StatusCreated, uBuf.ConvertToDTO())
	} else {
		return ctx.JSON(http.StatusCreated, tango_errors.ReturnDefault("Email", "Email already used", 0))
	}
}

func UpdateUser(ctx echo.Context, tapp *tangoapp.TangoApp) error {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	// get the incoming values
	uDTO := UserDTO{}
	if err := ctx.Bind(&uDTO); err != nil {
		return ctx.JSON(http.StatusBadRequest, "")
	}

	u := NewUser()
	uBuf, err := u.Update(tapp.DBAuth, id, uDTO)

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	return ctx.JSON(http.StatusOK, uBuf.ConvertToDTO())
}

func ChangePasswordUser(ctx echo.Context, tapp *tangoapp.TangoApp) error {
	updateDTO := &AuthDTOChangePassword{}
	ctx.Bind(updateDTO)

	auth := NewAuth()

	changed, err := auth.UpdatePassword(tapp.DBAuth, updateDTO)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	if !changed {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	return ctx.JSON(http.StatusOK, "Changed")

}

func DeleteUser(ctx echo.Context, tapp *tangoapp.TangoApp) error {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	u := NewUser()
	uBuf, err := u.Delete(tapp.DBAuth, id)

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	return ctx.JSON(http.StatusOK, uBuf.ConvertToDTO())
}
