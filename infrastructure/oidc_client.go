package infrastructure

import (
	"context"
	"fmt"
	"message/types"

	"github.com/coreos/go-oidc/v3/oidc"
)

type OidcClient struct {
	client   *oidc.Provider
	verifier *oidc.IDTokenVerifier
}

func NewOidcClient(oidcUrl string, realm string) (*OidcClient, error) {
	providerURL := fmt.Sprintf("%s/realms/%s", oidcUrl, realm)
	
	oidcProvider, err := oidc.NewProvider(context.Background(), providerURL)
	if err != nil {
		return nil, err
	}
	verifier := oidcProvider.Verifier(&oidc.Config{SkipClientIDCheck: true})
	return &OidcClient{
		oidcProvider,
		verifier,
	}, nil
}

func (client *OidcClient) VerifyToken(token string) (*types.User, error) {
	idToken, err := client.verifier.Verify(context.Background(), token)
	if err != nil {
		return nil, err
	}
	user := types.User{}
	err = idToken.Claims(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
