package request

import (
	"context"
	"github.com/spf13/cobra"
	oauth "golang.org/x/oauth2"
	"log"
	"net/http"
)

var authCmd = &cobra.Command{
	Use: "auth",

	Run: func(cmd *cobra.Command, args []string) {
		token, tokenErr := GetToken()
		if tokenErr != nil {
			log.Fatal("Can't get token. Check your network and try again")
			return
		}

		saveTokenErr := SaveToken(token)
		if saveTokenErr != nil {
			log.Fatal("Can't save authenticating information. Try again")
			return
		}
	},
}

var addListCmd = &cobra.Command{
	Use: "addlst",

	Run: func(cmd *cobra.Command, args []string) {
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

		ctx := context.Background()
		httpClient := &http.Client{}
		ctx = context.WithValue(ctx, oauth.HTTPClient, httpClient)

		token, err := GetSavedToken()
		if err != nil {
			log.Fatal("Can't get saved token :", err)
			return
		}

		webClient := authConfig.Client(ctx, token)

		lstErr := CreateTaskList(webClient, "made of cobra")
		if lstErr != nil {
			log.Fatal("Can't create a new task list : ", lstErr)
			return
		}
	},
}
