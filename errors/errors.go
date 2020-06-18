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
	TokenCreateFailed
)

// user error
const (
	GeneratePasswordFailed ErrorCode = iota + 20001
	CreateNewUserFailed
	UserNotFound
	InvalidPassword
	ListUserFailed
)

// exchange error
const (
	ListExchangeFailed ErrorCode = iota + 30001
	AddExchangeFailed
	UpdateExchangeFailed
	ExchangeNotFound
	InvalidExchangeID
)

var Error = map[ErrorCode]error{
	InvalidRequestParams: errors.New("invalid request params"),
	MissingRequestParams: errors.New("missing request params"),
	UnknownLoginType:     errors.New("unkown login type"),
	UnknownErr:           errors.New("unkown error"),
	TokenCreateFailed:    errors.New("token create failed"),

	// user
	GeneratePasswordFailed: errors.New("generate password fail"),
	CreateNewUserFailed:    errors.New("create new user failed"),
	UserNotFound:           errors.New("user not found"),
	InvalidPassword:        errors.New("invalid password"),
	ListUserFailed:         errors.New("list user failed"),

	ListExchangeFailed:   errors.New("list exchange failed"),
	AddExchangeFailed:    errors.New("add exchange failed"),
	UpdateExchangeFailed: errors.New("update exchange failed"),
	ExchangeNotFound:     errors.New("exchange not found"),
	InvalidExchangeID:    errors.New("exchange id must not emtpty"),
}

func GetMsg(e ErrorCode) string {
	err, ok := Error[e]
	if !ok {
		return Error[UnknownErr].Error()
	}
	return err.Error()
}
