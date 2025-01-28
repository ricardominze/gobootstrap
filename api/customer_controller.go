package api

import (
	"net/http"
	"strconv"

	account_entity "github.com/ricardominze/gobootstrap/core/domain/account/entity"
	account_service "github.com/ricardominze/gobootstrap/core/domain/account/service"
	customer_entity "github.com/ricardominze/gobootstrap/core/domain/customer/entity"
	customer_service "github.com/ricardominze/gobootstrap/core/domain/customer/service"
	"github.com/ricardominze/gobootstrap/core/valueobject"
	"github.com/ricardominze/gobootstrap/infra/restapi"
	"github.com/ricardominze/gobootstrap/infra/telemetry"
	"github.com/ricardominze/gobootstrap/infra/util"
)

type CustomerController struct {
	util.HandlerMap
	CustomerService *customer_service.CustomerService
	AccountService  *account_service.AccountService
}

func (h *CustomerController) MakeHandlers(router util.IRouter) {
	h.MapHandlers(h, router)
}

func NewCustomerController(CustomerService *customer_service.CustomerService, AccountService *account_service.AccountService) *CustomerController {
	handle := &CustomerController{CustomerService: CustomerService, AccountService: AccountService, HandlerMap: util.HandlerMap{}}
	return handle
}

func (h *CustomerController) GetRwp() string {
	return "{id}"
}

func (h *CustomerController) GetAction() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if err := restapi.RestrictMethod(r, "GET"); err != nil {
			restapi.ErrorResponse(w, err.Error(), 405, nil)
			return
		}

		stmap := util.NewStructMap(r)
		id, _ := strconv.Atoi(stmap.Vars(r)["id"])
		customer, err := h.CustomerService.Get(r.Context(), id)

		if err != nil {
			restapi.ErrorResponse(w, err.Error(), 405, nil)
			return
		}

		restapi.SuccessResponse(w, customer, "Success")
	})
}

func (h *CustomerController) CreateAction() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		//*****
		//Trace
		//*****

		ctx := r.Context()
		ctx, span := telemetry.MakeTraceCall(ctx)
		defer span.End()

		//**********
		//Requisição
		//**********

		if err := restapi.RestrictMethod(r, "POST"); err != nil {
			restapi.ErrorResponse(w, err.Error(), 405, nil)
			return
		}

		//*****
		//Dados
		//*****

		stmap := util.NewStructMap(r)

		address := &valueobject.Address{}
		stmap.BindData(&address)

		customer := &customer_entity.Customer{}
		stmap.BindData(&customer)

		customer.Address = address

		//*******
		//Servico
		//*******

		customer, err := h.CustomerService.Create(ctx, customer)

		account := &account_entity.Account{}
		account.IdCustomer = customer.Id
		account, err = h.AccountService.Open(ctx, account)

		if err != nil {
			restapi.ErrorResponse(w, err.Error(), 405, nil)
			return
		}

		//********
		//Resposta
		//********

		restapi.SuccessResponse(w, customer, "Success")
	})
}

func (h *CustomerController) UpdateRwp() string {
	return "{id}/update"
}

func (h *CustomerController) UpdateAction() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		//*****
		//Trace
		//*****

		ctx := r.Context()
		ctx, span := telemetry.MakeTraceCall(ctx)
		defer span.End()

		//**********
		//Requisição
		//**********

		if err := restapi.RestrictMethod(r, "POST"); err != nil {
			restapi.ErrorResponse(w, err.Error(), 405, nil)
			return
		}

		//*****
		//Dados
		//*****

		stmap := util.NewStructMap(r)

		address := &valueobject.Address{}
		stmap.BindData(&address)

		customer := &customer_entity.Customer{}
		stmap.BindData(&customer)

		customer.Address = address

		//*******
		//Servico
		//*******

		customer, err := h.CustomerService.Save(ctx, customer)

		if err != nil {
			restapi.ErrorResponse(w, err.Error(), 405, nil)
			return
		}

		//********
		//Resposta
		//********

		restapi.SuccessResponse(w, customer, "Success")
	})
}
