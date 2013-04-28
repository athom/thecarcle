package thecarcle

import (
	"github.com/athom/thecarcle/app/controllers"
	"github.com/bmizerany/pat"
	. "github.com/paulbellamy/mango"
	"github.com/sunfmin/mangolog"
	"net/http"
)

func Mux() (mux *http.ServeMux) {
	sessionMiddleware := Sessions(
		SITE_KEY,
		APP_NAME,
		&CookieOptions{Path: COOKIE_PATH, MaxAge: COOKIE_AGE},
	)

	publicStack := new(Stack)
	publicStack.Middleware(mangolog.Logger, Jsons)

	mainStack := new(Stack)
	mainStack.Middleware(mangolog.Logger, sessionMiddleware, Jsons)

	p := pat.New()
	p.Get("/about", publicStack.HandlerFunc(controllers.AboutPage))
	p.Get("/", publicStack.HandlerFunc(controllers.IndexPage))

	mux = http.NewServeMux()
	mux.Handle("/", p)
	return
}

func Jsons(env Env, app App) (status Status, headers Headers, body Body) {
	AppHost = env.Request().Host
	status, headers, body = app(env)
	if headers == nil {
		headers = Headers{}
	}

	headers.Set("Content-Type", "application/json; charset=utf-8")

	if status == 0 {
		status = 200
	}

	return status, headers, body
}
