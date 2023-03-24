package delivery

import (
	"lapakUmkm/app/middlewares"
	"lapakUmkm/features/carts"
	"lapakUmkm/utils/helpers"
	"net/http"

	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"
)

type CartHandler struct {
	srv carts.CartService
}

func New(srv carts.CartService) *CartHandler {
	return &CartHandler{
		srv: srv,
	}
}

func (ch *CartHandler) Add(c echo.Context) error {
	var formInput NewCartRequest
	if err := c.Bind(&formInput); err != nil {
		return c.JSON(http.StatusBadRequest, helpers.ResponseFail("error bind data"))
	}
	claim := middlewares.ClaimsToken(c)
	formInput.UserId = uint(claim.Id)
	newCart := carts.Core{}
	copier.Copy(&newCart, &formInput)
	data, err := ch.srv.Add(newCart)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.ResponseFail(err.Error()))
	}
	res := AddResponse{}
	copier.Copy(&res, &data)
	return c.JSON(http.StatusCreated, helpers.ResponseSuccess("success add product to cart", res))
}

func (ch *CartHandler) MyCart(c echo.Context) error {
	claim := middlewares.ClaimsToken(c)
	userId := uint(claim.Id)
	data, err := ch.srv.MyCart(userId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.ResponseFail(err.Error()))
	}
	res := ListCartResponse{}
	copier.Copy(&res, &data)
	return c.JSON(http.StatusOK, helpers.ResponseSuccess("success show your cart", res))
}