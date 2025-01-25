package adapter

import (
	"fmt"
	"net/http"
	"regexp"

	"github.com/ricardominze/gobootstrap/infra/util"
)

type DynamicRoute struct {
	Pattern *regexp.Regexp
	Params  []string
	Handler http.Handler
}

func CompileDynamicRoute(pattern string, handler http.Handler) DynamicRoute {
	paramRegex := regexp.MustCompile(`\{(\w+)(?::([^}]+))?\}`)
	params := []string{}
	regexPattern := paramRegex.ReplaceAllStringFunc(pattern, func(m string) string {
		matches := paramRegex.FindStringSubmatch(m)
		params = append(params, matches[1])
		if matches[2] != "" {
			return fmt.Sprintf("(%s)", matches[2])
		}
		return `([^/]+)`
	})
	regexPattern = "^" + regexPattern + "$"
	return DynamicRoute{
		Pattern: regexp.MustCompile(regexPattern),
		Params:  params,
		Handler: handler,
	}
}

type RouterClassic struct {
	routes     []DynamicRoute
	middleware []util.Middleware
}

func NewRouterClassic() util.IRouter {
	return &RouterClassic{}
}

func (rt *RouterClassic) New() util.IRouter {
	return &RouterClassic{}
}

func (rt *RouterClassic) Handle(pattern string, handler http.Handler) {
	route := CompileDynamicRoute(pattern, handler)
	rt.routes = append(rt.routes, route)
}

func (rt *RouterClassic) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	for _, route := range rt.routes {
		if matches := route.Pattern.FindStringSubmatch(path); matches != nil {
			params := r.URL.Query()
			for i, param := range matches[1:] {
				params.Set(route.Params[i], param)
			}
			r.URL.RawQuery = params.Encode()
			for _, mw := range rt.middleware {
				route.Handler = mw(route.Handler)
			}
			route.Handler.ServeHTTP(w, r)
			return
		}
	}
	http.NotFound(w, r)
}

func (rt *RouterClassic) Use(mw ...util.Middleware) {
	for _, v := range mw {
		rt.middleware = append(rt.middleware, v)
	}
}

func (rt *RouterClassic) Vars(r *http.Request) map[string]string {
	vars := make(map[string]string, len(r.URL.Query()))
	for k, v := range r.URL.Query() {
		vars[k] = v[0]
	}
	return vars
}
