package main

import src "github.com/Microsoft-To-Do-API/src"

func main() {
	src.Execute()
	/*
		webClient, _ := req.GetNewClient()
		err := req.CreateTaskList(webClient, "a new list called shit")
		fmt.Println("err = ", err)
	*/

	/*
		authEndpoint := oauth.Endpoint {
			AuthURL : "https://login.microsoftonline.com/common/oauth2/v2.0/authorize",
			TokenURL : "https://login.microsoftonline.com/common/oauth2/v2.0/token",
		}

		authConfig := oauth.Config {
			ClientID 		: "1f0be847-d5ad-4872-aa14-7e584d7cc940",
			ClientSecret 	: "mSq-Ser-o66Bt-ASLQ3RkQ2i~~OTB.n02a",
			Endpoint 		: authEndpoint,
			RedirectURL		: "https://login.microsoftonline.com/common/oauth2/nativeclient",
			Scopes 			: []string {
				"offline_access",
				"Tasks.ReadWrite",
			},
		}

		ctx := context.Background()
		httpClient := &http.Client{}
		ctx = context.WithValue(ctx, oauth.HTTPClient, httpClient)

		newToken := new(oauth.Token)
		newToken.RefreshToken = "0.AAAAdaWbQ4NMKUqo2_GVu2rGd0foCx-t1XJIqhR-WE18yUBvAP0.AgABAAAAAAB2UyzwtQEKR7-rWbgdcBZIAQDs_wIA9P_n4I_5u3LtRQbPr4V1KexKCYpq0c0SuSz5ryhSLw0BQiXFPir_0jR7BcA3SZUCWwZVTcAUD8tyrbzP11wEkPBNt4fnotY3YpYXX_Y3cY6_1rkQTm5Qcao1_GMOMsRojrClf5eU0SnBEW8wNY2KmAcUAWibW41TL-AxfO5licoJ_yhZpYqK93VoE8b9z2XgKyTf8QhlZDX94Rghtqe2bvUft_AncEcleNUB8ZwvZjggYYehf8g57EEPcgfXuJgEukzAu7PA5ynLMpWg3-EZLQ8cRzypk7d_C21Tm9qM66JKCxRugOWteJ4gO_MPJuWJf8d5RT9OZvV-SBx8HbWnbboY8UNaPa1liA6XYgrw-BR5xt2ZeXmCC5jhQPS6WxplLTcyH2YlmSWrvYW3w9d_4Z-EpG9HOd1zILkWdHp-PtLhDhUx0QD2kH3U9g1MXzhf5E0uA40scI8iOm1GgcUytNnYLnGHV0gVXOu5dhbMAB-Vvlvg2HqPvY48fVt5GXK6cvRqAomPl4FBrh2w9JYBFkj4ixS3kDm8Ug5rSlv2uald_3qrMBMBGA4tkuNDCO3VyqbdHQLjtZauT-I02hVSRq9ciLEO8vDcuMFV5sWNxjow5grdnyfe4_Jx9yxud82j02WSWctuasVAT_dL184HepSMz5x1nXG6rSY06ziBvEd7ehW42Cxir0SgP44gGBMdozOAK9Ar73G8AWch0VTdWuJemyQ7PTaUQyHUSMIxiAajz9GHjOSHq__AuUBDJ0rylDhdTuwfBTmxSelcKvOqkz3O2IjJAnS4wHclRdsSprbvkDqOHZmzpM8hR5bIaTnyzzXs_PoBY8npjwR4EmuLb7eWwnckWAfgzANbs4WJe4Tq9sZcq3RDJe8RVkiXF46urd30zpImRLm1jxElfQvW-H9h-FKLOperCDrMhSH7_hvTNb8cMqnDzK5J1-vC7KzHvWt8Dkmx2P89zg"
		newToken.TokenType = "Bearer"

		newToken, tokenErr := req.GetSavedToken()
		if tokenErr != nil {
			fmt.Println(":(")
		}

		webClient := authConfig.Client(ctx, newToken)

		err := req.CreateTaskList(webClient, "pew pew pew")
		fmt.Println("err_2 = ", err)
	*/
}
