package util

import (
	"fmt"
	"net/http"
	"reflect"
	"regexp"
	"strings"
)

//Handler

type IRouter interface {
	New() IRouter
	Use(mwf ...Middleware)
	Handle(string, http.Handler)
	Vars(r *http.Request) map[string]string
	ServeHTTP(http.ResponseWriter, *http.Request)
}

type Middleware func(http.Handler) http.Handler

type HandlerMap struct {
	SufixHandler string
	SufixMethod  string
	SufixParams  string
	Router       IRouter
	Routes       []string
}

type HandlerToMap interface {
	MakeHandlers(IRouter)
}

func (h *HandlerMap) addPath(path string) {
	h.Routes = append(h.Routes, path)
}

func (h *HandlerMap) GetPaths() []string {
	return h.Routes
}

func (h *HandlerMap) configSufixs() {

	if h.SufixHandler == "" {
		h.SufixHandler = "Controller"
	}

	if h.SufixMethod == "" {
		h.SufixMethod = "Action"
	}

	if h.SufixParams == "" {
		h.SufixParams = "Rwp"
	}
}

func (h *HandlerMap) toLowerPath(input string) string {

	re := regexp.MustCompile(`([a-z])([A-Z])`)
	result := re.ReplaceAllString(input, `$1-$2`)
	return strings.ToLower(result)
}

func (h *HandlerMap) MapHandlers(hm HandlerToMap, router IRouter) {

	h.configSufixs()
	h.Router = router

	handlerType := reflect.TypeOf(hm)
	handlerValue := reflect.ValueOf(hm)
	handlerName := handlerType.Elem().Name()
	basePath := h.toLowerPath(strings.Replace(handlerName, h.SufixHandler, "", 1))

	for i := 0; i < handlerType.NumMethod(); i++ {
		methodName := handlerType.Method(i).Name
		if strings.Contains(methodName, h.SufixMethod) {
			method := handlerValue.MethodByName(methodName)
			methodNoSufix := strings.Replace(methodName, h.SufixMethod, "", 1)
			methodPath := h.toLowerPath(methodNoSufix)
			methodParams := handlerValue.MethodByName(methodNoSufix + h.SufixParams)
			if method.Type().NumOut() == 1 && method.Type().Out(0) == reflect.TypeOf((*http.Handler)(nil)).Elem() {
				if handleFunction, ok := method.Call(nil)[0].Interface().(http.Handler); ok {
					path := "/" + basePath + "/" + methodPath
					if methodParams.IsValid() {
						path = "/" + basePath + "/" + methodParams.Call(nil)[0].String()
					}
					h.addPath(path)
					h.Router.Handle(path, handleFunction)
				} else {
					fmt.Println("The return type is not compatible with http.Handler")
				}
			} else {
				fmt.Println("The method (" + handlerType.Method(i).Name + ") does not a http.Handler")
			}
		}
	}
}
