package errors

import (
	"github.com/pkg/errors"
)

type ErrorCode uint64

const (
	UnknownErr ErrorCode = iota + 90001
)

// request error
const (
	InvalidRequestParams ErrorCode = iota + 10001
	MissingRequestParams
)

// user error
const (
	GeneratePasswordFail ErrorCode = iota + 20001
	CreateNewUserFail
)

var Error = map[ErrorCode]error{
	InvalidRequestParams: errors.New("invalid request params"),
	MissingRequestParams: errors.New("missing request params"),
	UnknownErr:           errors.New("unkown error"),

	// user
	GeneratePasswordFail: errors.New("generate password fail"),
	CreateNewUserFail:    errors.New("create new user fail"),
}

func GetMsg(e ErrorCode) string {
	err, ok := Error[e]
	if !ok {
		return Error[UnknownErr].Error()
	}
	return err.Error()
}
