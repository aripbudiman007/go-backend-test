package controller

import (
	"cart/helper"
	"cart/model/web"
	"cart/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type CartControllerImp struct {
	CartService service.CartService
}

func NewCartController(cartService service.CartService) CartController {
	return &CartControllerImp{
		CartService: cartService,
	}
}

func (controller *CartControllerImp) Get(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {
	cartResponses := controller.CartService.Get(request.Context())

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   cartResponses,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *CartControllerImp) Save(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	cartCreateRequest := web.CartCreateRequest{}
	helper.ReadFromRequestBody(request, &cartCreateRequest)

	cartResponse := controller.CartService.Save(request.Context(), cartCreateRequest)

	webResponse := web.WebResponse{
		Code:   201,
		Status: "Created",
		Data:   cartResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *CartControllerImp) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	cartId := params.ByName("cartId")
	id, err := strconv.Atoi(cartId)
	helper.PanicIfError(err)

	cartResponse := controller.CartService.FindById(request.Context(), id)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   cartResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *CartControllerImp) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	cartId := params.ByName("cartId")
	id, err := strconv.Atoi(cartId)
	helper.PanicIfError(err)

	controller.CartService.Delete(request.Context(), id)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, webResponse)
}
