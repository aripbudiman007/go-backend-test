package controller

import (
	"cart/service"
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type CartControllerImp struct {
	CartService service.CartService
}

func (controller *CartControllerImp) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	decoder := json.NewDecoder(request.Body)
}

func (controller *CartControllerImp) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

}

func (controller *CartControllerImp) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

}

func (controller *CartControllerImp) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

}
