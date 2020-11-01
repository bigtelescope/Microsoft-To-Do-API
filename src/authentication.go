package request

import (
	"fmt"
	"context"
	"net/http"
	oauth "golang.org/x/oauth2"
)

func GetDefaultClient() *http.Client {
	token, tokenErr := GetToken()
	if tokenErr != nil {
		return nil
	}

	tempToken := TokenResponse {
		TokenValue	: token,
		TokenErr 	: nil,
	}

	clientContext := context.Background()
	webClient  	  := oauth.NewClient(clientContext, tempToken)

	return webClient
}

func GetToken() (*oauth.Token, error) {
	authEndpoint := oauth.Endpoint {
		AuthURL : "https://login.microsoftonline.com/common/oauth2/v2.0/authorize",
		TokenURL : "https://login.microsoftonline.com/common/oauth2/v2.0/token",
	}

	authConfig := oauth.Config {
		ClientID 	: "1f0be847-d5ad-4872-aa14-7e584d7cc940",
		ClientSecret 	: "mSq-Ser-o66Bt-ASLQ3RkQ2i~~OTB.n02a",
		Endpoint 	: authEndpoint,
		RedirectURL	: "https://login.microsoftonline.com/common/oauth2/nativeclient",
		Scopes 		: []string {
					"offline_access",
					"Tasks.ReadWrite",
		},
	}

	resp := authConfig.AuthCodeURL("state", oauth.AccessTypeOffline)

	fmt.Println("go to the next link : ")
	fmt.Println(resp)
	fmt.Println()
	fmt.Println()

	var code string
	if _, err := fmt.Scan(&code); err != nil {
		fmt.Println("Can`t get code!")
		return nil, err
	}

	ctx := context.Background()
	httpClient := &http.Client{}
	ctx = context.WithValue(ctx, oauth.HTTPClient, httpClient)

	token, err := authConfig.Exchange(ctx, code)

	return token, err
}