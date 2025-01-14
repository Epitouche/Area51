package services

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"area51/schemas"
	"area51/toolbox"
)

type GoogleService interface {
	AuthGetServiceAccessToken(code string, path string) (schemas.GoogleResponseToken, error)
	GetUserInfosByToken(accessToken string, serviceName schemas.ServiceName) func(*schemas.ServicesUserInfos)
	FindActionByName(name string) func(channel chan string, option string, workflowId uint64, actionOption string)
	FindReactionByName(name string) func(channel chan string, workflowId uint64, accessToken []schemas.ServiceToken, reactionOption string)
}

type googleService struct {
}

func NewGoogleService() GoogleService {
	return &googleService{}
}

func (service *googleService) AuthGetServiceAccessToken(code string, path string) (schemas.GoogleResponseToken, error) {
	clientId := toolbox.GetInEnv("GOOGLE_CLIENT_ID")
	clientSecret := toolbox.GetInEnv("GOOGLE_SECRET")
	// appPort := toolbox.GetInEnv("FRONTEND_PORT")
	// appAdressHost := toolbox.GetInEnv("APP_HOST_ADDRESS")

	// redirectUri := appAdressHost + appPort + path
	redirectUri := "http://localhost:8081/callback"
	fmt.Printf("redirectUri: %s\n", redirectUri)
	apiUrl := "https://oauth2.googleapis.com/token"
	// fmt.Printf("CODE : %s\n", code)

	data := url.Values{}
	data.Set("client_id", clientId)
	data.Set("client_secret", clientSecret)
	data.Set("code", code)
	data.Set("redirect_uri", redirectUri)
	data.Set("grant_type", "authorization_code")

	// req, err := http.NewRequest("POST", apiUrl, strings.NewReader(data.Encode()))
	// if err != nil {
	// 	return schemas.GoogleResponseToken{}, err
	// }
	ctx := context.Background()

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, apiUrl, nil)
	if err != nil {
		return schemas.GoogleResponseToken{}, fmt.Errorf("unable to create request because %w", err)
	}
	req.URL.RawQuery = data.Encode()
	req.Header.Set("Accept", "application/json")
	// req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	client := &http.Client{}
	response, err := client.Do(req)
	fmt.Printf("response: %++v\n", response)
	fmt.Printf("response BODY: %++v\n", response.Body)
	if err != nil {
		return schemas.GoogleResponseToken{}, err
	}
	defer response.Body.Close()
	// bodyBytes, _ := io.ReadAll(response.Body)
	// fmt.Println("response body: ", string(bodyBytes))

	var result schemas.GoogleResponseToken
	// err = json.Unmarshal(bodyBytes, &result)
	err = json.NewDecoder(response.Body).Decode(&result)
	fmt.Printf("result: %++v\n", result)
	if err != nil {
		return schemas.GoogleResponseToken{}, fmt.Errorf("unable to decode response because %w", err)
	}
	return result, nil
}

func (service *googleService) GetUserInfosByToken(accessToken string, serviceName schemas.ServiceName) func(*schemas.ServicesUserInfos) {
	return func(userInfos *schemas.ServicesUserInfos) {
		fmt.Printf("accessToken: %s\n", accessToken)
		request, err := http.NewRequest("GET", "https://www.googleapis.com/oauth2/v1/userinfo?alt=json", nil)
		if err != nil {
			return
		}

		request.Header.Set("Authorization", "Bearer "+accessToken)
		client := &http.Client{}

		response, err := client.Do(request)
		if err != nil {
			return
		}
		defer response.Body.Close()

		fmt.Printf("response: %++v\n", response)
		if response.StatusCode != http.StatusOK {
			fmt.Printf("Error: received status code %d\n", response.StatusCode)
			return
		}

		err = json.NewDecoder(response.Body).Decode(&userInfos.GoogleUserInfos)
		if err != nil {
			return
		}
	}
}

func (service *googleService) FindActionByName(name string) func(channel chan string, option string, workflowId uint64, actionOption string) {
	return func(channel chan string, option string, workflowId uint64, actionOption string) {
		channel <- "google"
	}
}

func (service *googleService) FindReactionByName(name string) func(channel chan string, workflowId uint64, accessToken []schemas.ServiceToken, reactionOption string) {
	return func(channel chan string, workflowId uint64, accessToken []schemas.ServiceToken, reactionOption string) {
		channel <- "google"
	}
}
