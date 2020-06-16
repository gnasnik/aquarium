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
	UnknownLoginType
	TokenCreateFail
)

// user error
const (
	GeneratePasswordFail ErrorCode = iota + 20001
	CreateNewUserFail
	UserNotFound
	InvalidPassword
)

var Error = map[ErrorCode]error{
	InvalidRequestParams: errors.New("invalid request params"),
	MissingRequestParams: errors.New("missing request params"),
	UnknownLoginType:     errors.New("unkown login type"),
	UnknownErr:           errors.New("unkown error"),
	TokenCreateFail:      errors.New("token create failed"),

	// user
	GeneratePasswordFail: errors.New("generate password fail"),
	CreateNewUserFail:    errors.New("create new user fail"),
	UserNotFound:         errors.New("user not found"),
	InvalidPassword:      errors.New("invalid password"),
}

func GetMsg(e ErrorCode) string {
	err, ok := Error[e]
	if !ok {
		return Error[UnknownErr].Error()
	}
	return err.Error()
}
