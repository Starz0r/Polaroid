package routers

import (
	"encoding/json"
	"net/http"

	"github.com/Starz0r/Polaroid/src/auth"
	"github.com/coreos/go-oidc"
	echo "github.com/spidernest-go/mux"
	"golang.org/x/oauth2"
)

func userLogin(c echo.Context) error {
	return c.Redirect(http.StatusFound, auth.Authenticator.AuthCodeURL(auth.State, oidc.Nonce(auth.Nonce)))
}

func userRedirect(c echo.Context) error {
	// compare states
	if c.QueryParams().Get("state") != auth.State {
		return c.JSON(http.StatusUnauthorized, RespError{Err: "state did not match"})
	}

	// exchance the code for the token
	tkn, err := auth.Authenticator.Exchange(auth.Ctx, c.QueryParams().Get("code"))
	if err != nil {
		return c.JSON(http.StatusUnauthorized, RespError{Err: "auth token exchange failed"})
	}

	// check for the token in the response
	raw, ok := tkn.Extra("id_token").(string)
	if !ok {
		return c.JSON(http.StatusUnauthorized, RespError{Err: "no id_token in auth server exchange"})
	}

	// verify the nonce
	id, err := auth.NonceEnabledVerifier.Verify(auth.Ctx, raw)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, RespError{Err: "nonce could not be verified"})
	}
	if id.Nonce != auth.Nonce {
		return c.JSON(http.StatusUnauthorized, RespError{Err: "nonce was invalid"})
	}

	// check for the claims and return the token
	resp := struct {
		Token  *oauth2.Token
		Claims *json.RawMessage
	}{tkn, new(json.RawMessage)}

	if err := id.Claims(&resp.Claims); err != nil {
		return c.JSON(http.StatusUnauthorized, RespError{Err: "required claims were missing"})
	}
	data, err := json.MarshalIndent(resp, "", "    ")
	if err != nil {
		return c.JSON(http.StatusUnauthorized, RespError{Err: "response could not be marshalled"})
	}

	return c.JSONBlob(http.StatusAccepted, data)
}
