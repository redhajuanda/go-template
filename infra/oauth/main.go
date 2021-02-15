package oauth

import (
	"context"
	"go-template/config"
	"net/http"

	"github.com/coreos/go-oidc"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	verify "google.golang.org/api/oauth2/v2"
)

func New(clientID, clientSecret, redirectURL string) *oauth2.Config {
	conf := &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  redirectURL,
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email", // You have to select your own scope from here -> https://developers.google.com/identity/protocols/googlescopes#google_sign-in
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}
	google.DefaultTokenSource(context.Background())

	return conf
}

func GetTokenInfo(idToken string) (*verify.Tokeninfo, error) {
	oauth2Service, err := verify.New(&http.Client{})
	oauth2Service.Userinfo.V2.Me.Get().Do()
	if err != nil {
		return nil, err
	}
	tokenInfoCall := oauth2Service.Tokeninfo()
	tokenInfoCall.IdToken(idToken)

	return tokenInfoCall.Do()
}

func Verify(rawIDToken string) error {
	ctx := context.Background()
	cfg := config.LoadDefault()
	provider, err := oidc.NewProvider(ctx, "https://accounts.google.com")
	if err != nil {
		return err
	}

	var verifier = provider.Verifier(&oidc.Config{ClientID: cfg.Google.ClientID})

	// Parse and verify ID Token payload.
	_, err = verifier.Verify(ctx, rawIDToken)
	if err != nil {
		// handle error
	}
	return nil
}
