package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"
)

type router struct {
	muxRouter *mux.Router
}

func newRouter(muxRouter *mux.Router) (*router, error) {
	if muxRouter == nil {
		return nil, errors.New("missing mux router")
	}

	return &router{
		muxRouter: muxRouter,
	}, nil
}

func (rtr *router) Handle(path string, handler http.Handler) *mux.Route {
	return rtr.muxRouter.Handle(path, handler)
}

func (rtr *router) HandleFunc(path string, handlerFunc http.HandlerFunc) *mux.Route {
	return rtr.muxRouter.Handle(path, handlerFunc)
}
