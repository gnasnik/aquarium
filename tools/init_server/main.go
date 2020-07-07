package main

import (
	"context"

	"github.com/frankffenn/aquarium/config"
	"github.com/frankffenn/aquarium/sdk"
	"github.com/frankffenn/aquarium/sdk/mod"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	// TODO:
	config.InitConfig()

	createSuperUser()
}

func createSuperUser() error {
	found, _ := sdk.GetUser(context.Background(), "admin")
	if found != nil {
		return nil
	}

	passHash, err := bcrypt.GenerateFromPassword([]byte("123456"), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user := &mod.User{
		Username: "admin",
		Password: string(passHash),
	}

	sdk.CreateUser(context.Background(), user)
	return nil
}
