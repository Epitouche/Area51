package services

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"area51/schemas"
	"area51/toolbox"
)

type MicrosoftService interface {
	GetUserInfosByToken(accessToken string, serviceName schemas.ServiceName) func(*schemas.ServicesUserInfos)
	FindActionByName(name string) func(channel chan string, option string, workflowId uint64, actionOption string)
	FindReactionByName(name string) func(channel chan string, workflowId uint64, accessToken []schemas.ServiceToken, reactionOption string)
	AuthGetServiceAccessToken(code string, path string) (schemas.MicrosoftResponseToken, error)
}

type microsoftService struct {
}

func NewMicrosoftService() MicrosoftService {
	return &microsoftService{}
}

func (service *microsoftService) GetUserInfosByToken(accessToken string, serviceName schemas.ServiceName) func(*schemas.ServicesUserInfos) {
	return func(userInfos *schemas.ServicesUserInfos) {
		ctx := context.Background()

		url := "https://graph.microsoft.com/v1.0/me"
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
		if err != nil {
			return
		}

		req.Header.Set("Authorization", "Bearer "+accessToken)

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			return
		}
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			// READ THE BODY
			bodyBytes, _ := io.ReadAll(resp.Body)
			fmt.Println("response body: ", string(bodyBytes))
			return
		}
		err = json.NewDecoder(resp.Body).Decode(&userInfos.MicrosoftUserInfos)
		if err != nil {
			return
		}
	}
}

func (service *microsoftService) AuthGetServiceAccessToken(code string, path string) (schemas.MicrosoftResponseToken, error) {
	clientId := toolbox.GetInEnv("MICROSOFT_CLIENT_ID")
	tenantId := toolbox.GetInEnv("MICROSOFT_TENANT_ID")
	appPort := toolbox.GetInEnv("FRONTEND_PORT")
	appAdressHost := toolbox.GetInEnv("APP_HOST_ADDRESS")

	redirectUri := appAdressHost + appPort + path
	apiUrl := fmt.Sprintf("https://login.microsoftonline.com/%s/oauth2/v2.0/token", tenantId)

	data := url.Values{}
	data.Set("client_id", clientId)
	data.Set("code", code)
	data.Set("redirect_uri", redirectUri)
	data.Set("grant_type", "authorization_code")

	req, err := http.NewRequest("POST", apiUrl, strings.NewReader(data.Encode()))
	if err != nil {
		return schemas.MicrosoftResponseToken{}, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	client := &http.Client{
		Timeout: time.Second * 45,
	}
	response, err := client.Do(req)
	fmt.Printf("response: %++v\n", response)
	if err != nil {
		return schemas.MicrosoftResponseToken{}, err
	}
	bodyBytes, _ := io.ReadAll(response.Body)
	fmt.Println("response body: ", string(bodyBytes))

	var result schemas.MicrosoftResponseToken
	err = json.Unmarshal(bodyBytes, &result)
	if err != nil {
		return schemas.MicrosoftResponseToken{}, fmt.Errorf("unable to decode response because %w", err)
	}
	response.Body.Close()
	return result, nil
}

func (service *microsoftService) FindActionByName(name string) func(channel chan string, option string, workflowId uint64, actionOption string) {
	return nil
}

func (service *microsoftService) FindReactionByName(name string) func(channel chan string, workflowId uint64, accessToken []schemas.ServiceToken, reactionOption string) {
	return nil
}
