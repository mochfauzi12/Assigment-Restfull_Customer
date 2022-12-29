package router

import (
	"Assigment-Restfull_Customer/controller"
	"Assigment-Restfull_Customer/exception"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(customerController controller.CustomerController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/customers", customerController.FindAll)

	router.GET("/api/customers/:customerId", customerController.FindById)

	router.POST("/api/customers", customerController.Create)

	router.PUT("/api/customers/:customerId", customerController.Update)

	router.DELETE("/api/customers/:customerId", customerController.Delete)

	router.PanicHandler = exception.ErrorHandler

	return router
}
