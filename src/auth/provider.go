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
	ctx                  context.Context
	NonceEnabledVerifier *oidc.IDTokenVerifier
	State                string
)

func GetConfiguration() error {
	ctx = context.Background()
	provider, err := oidc.NewProvider(ctx, os.Getenv("OIDC_URL"))
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
		RedirectURL:  "http://localhost:6000/api/v0/oidc/callback",
		Scopes:       []string{oidc.ScopeOpenID, "profile", "email"},
	}

	State = crypto.String(16)

	return nil
}
