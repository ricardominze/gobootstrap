package api

import (
	"fmt"
	"net/http"
	"strconv"

	account_service "github.com/ricardominze/gobootstrap/core/domain/account/service"
	"github.com/ricardominze/gobootstrap/infra/restapi"
	"github.com/ricardominze/gobootstrap/infra/telemetry"
	"github.com/ricardominze/gobootstrap/infra/util"
)

type AccountController struct {
	util.HandlerMap
	AccountService *account_service.AccountService
}

func (h *AccountController) MakeHandlers(router util.IRouter) {
	h.MapHandlers(h, router)
}

func NewAccountController(AccountService *account_service.AccountService) *AccountController {
	handle := &AccountController{AccountService: AccountService, HandlerMap: util.HandlerMap{}}
	return handle
}

func (h *AccountController) BalanceRwp() string {
	return "{id}/balance"
}

func (h *AccountController) BalanceAction() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		fmt.Println(h.Router.Vars(r))

		// Restricao de metodo aceito na requisicao

		if err := restapi.RestrictMethod(r, "GET"); err != nil {
			restapi.ErrorResponse(w, err.Error(), 405, nil)
			return
		}

		// Recuperando o contexto para propagacao

		ctx := r.Context()
		ctx, span := telemetry.MakeTraceCall(ctx)
		defer span.End()

		// Chamada do Serviço

		querystring := h.Router.Vars(r)

		idAccount, err := strconv.ParseInt(querystring["id"], 10, 32)

		if err != nil {
			restapi.ErrorResponse(w, err.Error(), 405, nil)
			return
		}

		balance, err := h.AccountService.Balance(ctx, int(idAccount))

		if err != nil {
			restapi.ErrorResponse(w, err.Error(), 405, nil)
			return
		}

		restapi.SuccessResponse(w, fmt.Sprintf("Account Balance: %v", balance), "Success")
	})
}

func (h *AccountController) DepositAction() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Restricao de metodo aceito na requisicao

		if err := restapi.RestrictMethod(r, "POST"); err != nil {
			restapi.ErrorResponse(w, err.Error(), 405, nil)
			return
		}

		// Recuperando o contexto para propagacao

		ctx := r.Context()
		ctx, span := telemetry.MakeTraceCall(ctx)
		defer span.End()

		// Chamada do Serviço

		stmap := util.NewStructMap(r)
		genstr := stmap.GenericStruct(stmap.GetLoadedData())

		amount := genstr["amount"].(float64)
		account, err := h.AccountService.Get(ctx, int(genstr["id"].(float64)))

		if err != nil {
			restapi.ErrorResponse(w, err.Error(), 405, nil)
			return
		}

		err = h.AccountService.Deposit(ctx, account, amount)

		if err != nil {
			restapi.ErrorResponse(w, err.Error(), 405, nil)
			return
		}

		account, err = h.AccountService.Get(ctx, account.Id)

		if err != nil {
			restapi.ErrorResponse(w, err.Error(), 405, nil)
			return
		}

		restapi.SuccessResponse(w, account, "Success")
	})
}

func (h *AccountController) TransferAction() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Restricao de metodo aceito na requisicao

		if err := restapi.RestrictMethod(r, "POST"); err != nil {
			restapi.ErrorResponse(w, err.Error(), 405, nil)
			return
		}

		// Recuperando o contexto para propagacao

		ctx := r.Context()
		ctx, span := telemetry.MakeTraceCall(ctx)
		defer span.End()

		// Chamada do Serviço

		stmap := util.NewStructMap(r)
		genstr := stmap.GenericStruct(stmap.GetLoadedData())

		idAccountFrom := int(genstr["id_account_from"].(float64))
		idAccountTo := int(genstr["id_account_to"].(float64))
		amount := genstr["amount"].(float64)

		accountFrom, err := h.AccountService.Get(ctx, idAccountFrom)

		if err != nil {
			restapi.ErrorResponse(w, err.Error(), 405, nil)
			return
		}

		accountTo, err := h.AccountService.Get(ctx, idAccountTo)

		if err != nil {
			restapi.ErrorResponse(w, err.Error(), 405, nil)
			return
		}

		err = h.AccountService.Transfer(ctx, accountFrom, accountTo, amount)

		if err != nil {
			restapi.ErrorResponse(w, err.Error(), 405, nil)
			return
		}

		balanceFrom, err := h.AccountService.Balance(ctx, accountFrom.Id)

		if err != nil {
			restapi.ErrorResponse(w, err.Error(), 405, nil)
			return
		}

		balanceTo, err := h.AccountService.Balance(ctx, accountTo.Id)

		if err != nil {
			restapi.ErrorResponse(w, err.Error(), 405, nil)
			return
		}

		restapi.SuccessResponse(w, fmt.Sprintf("Account Balance From: %v\nAccount Balance To: %v", balanceFrom, balanceTo), "Success")
	})
}

func (h *AccountController) WithdrawAction() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Restricao de metodo aceito na requisicao

		if err := restapi.RestrictMethod(r, "POST"); err != nil {
			restapi.ErrorResponse(w, err.Error(), 405, nil)
			return
		}

		// Recuperando o contexto para propagacao

		ctx := r.Context()
		ctx, span := telemetry.MakeTraceCall(ctx)
		defer span.End()

		// Chamada do Serviço

		stmap := util.NewStructMap(r)
		genstr := stmap.GenericStruct(stmap.GetLoadedData())

		idAccount := int(genstr["id_account"].(float64))
		amount := genstr["amount"].(float64)

		account, err := h.AccountService.Get(ctx, idAccount)

		if err != nil {
			restapi.ErrorResponse(w, err.Error(), 405, nil)
			return
		}

		balance := account.Balance
		err = h.AccountService.Withdraw(ctx, account, amount)

		if err != nil {
			restapi.ErrorResponse(w, err.Error(), 405, nil)
			return
		}

		restapi.SuccessResponse(w, fmt.Sprintf("Account Balance before Withdraw: %v, Account Balance after Withdraw: %v", balance, account.Balance), "Success")
	})
}
