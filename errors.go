package api

import (
	"github.com/pkg/errors"
)

var (
	ErrCouldNotReadBody = errors.New("could not read body")
	ErrInvalidParam     = errors.New("invalid parameter")
	ErrBadRequest       = errors.New("bad request")
)
