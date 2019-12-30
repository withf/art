package router

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

type router struct {
	routes  []*route
	parties []*party
	before  []context.Handler
	done    []context.Handler
}

type route struct {
	method  string
	path    string
	handler []context.Handler
}

type party struct {
	path      string
	routes    []*route
	irisParty iris.Party
	before    []context.Handler
	done      []context.Handler
}

// 唯一路由
var Router = &router{make([]*route, 0), make([]*party, 0), make([]context.Handler, 0), make([]context.Handler, 0)}

// Bind 路由绑定
func (r *router) Bind(app *iris.Application) {
	if len(r.before) > 0 {
		app.Use(r.before...)
	}
	if len(r.done) > 0 {
		app.Done(r.done...)
	}
	for _, r := range r.routes {
		app.Handle(r.method, r.path, r.handler...)
	}
	for _, p := range r.parties {
		p.irisParty = app.Party(p.path)
		if len(p.before) > 0 {
			p.irisParty.Use(p.before...)
		}
		if len(p.done) > 0 {
			p.irisParty.Done(p.done...)
		}
		for _, r := range p.routes {
			p.irisParty.Handle(r.method, r.path, r.handler...)
		}
	}
}

func (r *router) Register(method, path string, handler ...context.Handler) {
	r.routes = append(r.routes, &route{method, path, handler})
}

func (r *router) Get(path string, handler ...context.Handler) {
	r.Register("GET", path, handler...)
}

func (r *router) Post(path string, handler ...context.Handler) {
	r.Register("POST", path, handler...)
}

func (r *router) DELETE(path string, handler ...context.Handler) {
	r.Register("DELETE", path, handler...)
}

func (r *router) PUT(path string, handler ...context.Handler) {
	r.Register("PUT", path, handler...)
}

func (r *router) PATCH(path string, handler ...context.Handler) {
	r.Register("PATCH", path, handler...)
}

func (r *router) Party(path string) *party {
	p := &party{path, make([]*route, 0), nil, make([]context.Handler, 0), make([]context.Handler, 0)}
	r.parties = append(r.parties, p)
	return p
}

func (r *router) Use(handler ...context.Handler) {
	r.before = append(r.before, handler...)
}

func (r *router) Done(handler ...context.Handler) {
	r.done = append(r.done, handler...)
}

func (p *party) Register(method, path string, handler ...context.Handler) {
	p.routes = append(p.routes, &route{method, path, handler})
}

func (p *party) Get(path string, handler ...context.Handler) {
	p.Register("GET", path, handler...)
}

func (p *party) Post(path string, handler ...context.Handler) {
	p.Register("POST", path, handler...)
}

func (p *party) DELETE(path string, handler ...context.Handler) {
	p.Register("DELETE", path, handler...)
}

func (p *party) PUT(path string, handler ...context.Handler) {
	p.Register("PUT", path, handler...)
}

func (p *party) PATCH(path string, handler ...context.Handler) {
	p.Register("PATCH", path, handler...)
}

func (p *party) Party(path string) *party {
	pa := &party{p.path + path, make([]*route, 0), nil, p.before, p.done}
	Router.parties = append(Router.parties, pa)
	return pa
}

func (p *party) Use(handler ...context.Handler) {
	p.before = append(p.before, handler...)
}

func (p *party) Done(handler ...context.Handler) {
	p.done = append(p.done, handler...)
}
