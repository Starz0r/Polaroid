package auth

import (
	"os"

	"github.com/Starz0r/Polaroid/src/crypto"
	"github.com/coreos/go-oidc"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
)

var (
	Authenticator        oauth2.Config
	Ctx                  context.Context
	NonceEnabledVerifier *oidc.IDTokenVerifier
	State                string
	Nonce                string
)

func GetConfiguration() error {
	Ctx = context.Background()
	provider, err := oidc.NewProvider(Ctx, os.Getenv("OIDC_URL"))
	if err != nil {
		return err
	}

	cid := os.Getenv("OIDC_CLIENT_ID")

	cfg := &oidc.Config{
		ClientID: cid,
	}
	NonceEnabledVerifier = provider.Verifier(cfg)

	Authenticator = oauth2.Config{
		ClientID:     cid,
		ClientSecret: os.Getenv("OIDC_CLIENT_SECRET"),
		Endpoint:     provider.Endpoint(),
		RedirectURL:  os.Getenv("OIDC_REDIRECT_URL"),
		Scopes:       []string{oidc.ScopeOpenID, "profile", "email"},
	}

	Nonce = crypto.String(32)
	State = crypto.String(16)

	return nil
}
