package adapter

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ricardominze/gobootstrap/infra/util"
)

type RouterGorilla struct {
	router     *mux.Router
	middleware []util.Middleware
}

func NewRouterGorilla() util.IRouter {
	return &RouterGorilla{router: mux.NewRouter()}
}

func (g *RouterGorilla) New() util.IRouter {
	return &RouterGorilla{router: g.router}
}

func (g *RouterGorilla) Handle(path string, handler http.Handler) {
	g.router.NewRoute().Path(path).Handler(handler)
}

func (g *RouterGorilla) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	g.router.ServeHTTP(w, r)
}

func (g *RouterGorilla) MiddlewareToFunc(m util.Middleware) mux.MiddlewareFunc {
	return mux.MiddlewareFunc(m)
}

func (g *RouterGorilla) Use(mw ...util.Middleware) {
	for _, v := range mw {
		g.router.Use(g.MiddlewareToFunc(v))
	}
}

func (g *RouterGorilla) Vars(r *http.Request) map[string]string {
	return mux.Vars(r)
}
