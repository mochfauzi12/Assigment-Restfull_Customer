package controller

import (
	"Assigment-Restfull_Customer/controller"
	"Assigment-Restfull_Customer/domain/usecase"
	"Assigment-Restfull_Customer/domain/web"
	"Assigment-Restfull_Customer/internal/helper"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type CustomerControllers struct {
	CustomerService usecase.CustomerUseCase
}

func NewCustomerController(customerService usecase.CustomerUseCase) controller.CustomerController {
	return &CustomerControllers{
		CustomerService: customerService,
	}
}

func (controller *CustomerControllers) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	customerCreateRequest := web.CustomerCreateRequest{}
	helper.ReadFromRequestBody(request, &customerCreateRequest)

	customerResponse := controller.CustomerService.Create(
		request.Context(), customerCreateRequest)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   customerResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)

}

func (controller *CustomerControllers) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	customerUpdateRequest := web.CustomerUpdateRequest{}
	helper.ReadFromRequestBody(request, &customerUpdateRequest)

	customerId := params.ByName("customerId")
	id, err := strconv.Atoi(customerId)
	helper.PanicIfError(err)

	customerUpdateRequest.Id = id

	customerResponse := controller.CustomerService.Update(request.Context(), customerUpdateRequest)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   customerResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *CustomerControllers) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	customerId := params.ByName("customerId")
	id, err := strconv.Atoi(customerId)
	helper.PanicIfError(err)

	controller.CustomerService.Delete(request.Context(), id)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *CustomerControllers) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	customerId := params.ByName("customerId")
	id, err := strconv.Atoi(customerId)
	helper.PanicIfError(err)
	customerResponse := controller.CustomerService.FindById(request.Context(), id)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   customerResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)

}

func (controller *CustomerControllers) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	customerResponse := controller.CustomerService.FindAll(request.Context())

	webresponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   customerResponse,
	}

	helper.WriteToResponseBody(writer, webresponse)

}
