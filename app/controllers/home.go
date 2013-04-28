package controllers

import (
	"encoding/json"
	. "github.com/paulbellamy/mango"
)

func IndexPage(env Env) (status Status, headers Headers, body Body) {
	b, _ := json.Marshal("Index Page")
	body = Body(b)
	return
}

func AboutPage(env Env) (status Status, headers Headers, body Body) {
	b, _ := json.Marshal("About Page")
	body = Body(b)
	return
}
