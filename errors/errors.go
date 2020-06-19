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
	DeleteExchangeFailed
)

// algorithm error
const (
	ListAlgorithmFailed ErrorCode = iota + 40001
	AddAlgorithmFailed
	UpdateAlgorithmFailed
	AlgorithmNotFound
	InvalidAlgorithmID
	DeleteAlgorithmFailed
)

// trader error
const (
	ListTraderFailed ErrorCode = iota + 50001
	AddTraderFailed
	UpdateTraderFailed
	TraderNotFound
	InvalidTraderID
	DeleteTraderFailed
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
	DeleteExchangeFailed: errors.New("delete exchange failed"),

	ListAlgorithmFailed:   errors.New("list algorithm failed"),
	AddAlgorithmFailed:    errors.New("add algorithm failed"),
	UpdateAlgorithmFailed: errors.New("update algorithm failed"),
	AlgorithmNotFound:     errors.New("algorithm not found"),
	InvalidAlgorithmID:    errors.New("algorithm id must not emtpty"),
	DeleteAlgorithmFailed: errors.New("delete algorithm failed"),

	ListTraderFailed:   errors.New("list trader failed"),
	AddTraderFailed:    errors.New("add trader failed"),
	UpdateTraderFailed: errors.New("update trader failed"),
	TraderNotFound:     errors.New("trader not found"),
	InvalidTraderID:    errors.New("trader id must not emtpty"),
	DeleteTraderFailed: errors.New("delete trader failed"),
}

func GetMsg(e ErrorCode) string {
	err, ok := Error[e]
	if !ok {
		return Error[UnknownErr].Error()
	}
	return err.Error()
}
