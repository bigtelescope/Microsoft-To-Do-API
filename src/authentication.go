package request

import (
	"context"
	"fmt"
	oauth "golang.org/x/oauth2"
	"io/ioutil"
	"net/http"
	"os"
)

func GetNewClient() (*http.Client, error) {
	token, tokenErr := GetToken()
	if tokenErr != nil {
		return nil, tokenErr
	}

	saveTokenErr := SaveToken(token)
	if saveTokenErr != nil {
		return nil, saveTokenErr
	}

	tempToken := TokenResponse{
		TokenValue: token,
		TokenErr:   nil,
	}

	clientContext := context.Background()
	webClient := oauth.NewClient(clientContext, tempToken)

	return webClient, nil
}

func GetToken() (*oauth.Token, error) {
	authEndpoint := oauth.Endpoint{
		AuthURL:  "https://login.microsoftonline.com/common/oauth2/v2.0/authorize",
		TokenURL: "https://login.microsoftonline.com/common/oauth2/v2.0/token",
	}

	authConfig := oauth.Config{
		ClientID:     "1f0be847-d5ad-4872-aa14-7e584d7cc940",
		ClientSecret: "mSq-Ser-o66Bt-ASLQ3RkQ2i~~OTB.n02a",
		Endpoint:     authEndpoint,
		RedirectURL:  "https://login.microsoftonline.com/common/oauth2/nativeclient",
		Scopes: []string{
			"offline_access",
			"Tasks.ReadWrite",
		},
	}

	resp := authConfig.AuthCodeURL("state", oauth.AccessTypeOffline)

	fmt.Println("go to the next link : ")
	fmt.Println(resp)
	fmt.Println()
	fmt.Println("Put the code below:")

	var code string
	if _, err := fmt.Scan(&code); err != nil {
		fmt.Println("Can`t get code!")
		return nil, err
	}

	fmt.Println(code)

	ctx := context.Background()
	httpClient := &http.Client{}
	ctx = context.WithValue(ctx, oauth.HTTPClient, httpClient)

	token, err := authConfig.Exchange(ctx, code)

	return token, err
}

func GetSavedToken() (*oauth.Token, error) {
	savedRefreshToken, openErr := ioutil.ReadFile("token.txt")
	if openErr != nil {
		return nil, openErr
	}

	newToken := oauth.Token{
		RefreshToken: string(savedRefreshToken),
		TokenType:    "Bearer",
	}

	return &newToken, nil
}

func SaveToken(token *oauth.Token) error {
	file, fileErr := os.Create("token.txt")
	if fileErr != nil {
		return fileErr
	}

	tokenValue := []byte(token.RefreshToken)
	_, writeErr := file.Write(tokenValue)
	if writeErr != nil {
		closeErr := file.Close()
		if closeErr != nil {
			return closeErr
		}

		return writeErr
	}

	closeErr := file.Close()
	if closeErr != nil {
		return closeErr
	}
	return nil
}
