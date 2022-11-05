package services

import (
	"auth-api/auth"
	"auth-api/db"
	"auth-api/models/domains"
	"auth-api/models/dtos"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	GOOGLE_AUTH_PROVIDER     = "google"
	CREDENTIAL_AUTH_PROVIDER = "credential"
)

func CredentialSignup(data *dtos.CredentialSignupDto) (*dtos.RegistrationDto, *dtos.ErrorDto) {
	user := &domains.User{
		Name:            data.Name,
		Phone:           data.Phone,
		Email:           data.Email,
		AuthProvider:    CREDENTIAL_AUTH_PROVIDER,
		IsEmailVerified: false,
	}
	user.SetPassword(data.Password)

	if err := db.GetInstance().Create(user).Error; err != nil {
		return nil, &dtos.ErrorDto{Message: err.Error()}
	}

	// dispatching verification email to that user
	go func() {
		emailToken, _ := auth.GetEmailVerifierToken(data.Email)
		DispatchEmail(&EmailMessage{Subject: "Email Verification", Receiver: data.Email, Data: map[string]string{
			"Name": data.Name,
			"URL":  fmt.Sprintf("http://localhost:3000/auth/email-verify?token=%s", emailToken),
		}})
	}()

	token, err := auth.GenerateToken(user)
	if err != nil {
		return nil, &dtos.ErrorDto{Message: err.Error()}
	}

	userDto := &dtos.UserDto{
		ID:              user.ID,
		Name:            user.Name,
		Email:           user.Email,
		Phone:           user.Phone,
		IsEmailVerified: user.IsEmailVerified,
	}

	response := &dtos.RegistrationDto{}
	response.UserDto = userDto
	response.AccessToken = token

	return response, nil
}

func OAuthSignup(data *dtos.OAuthDto) (*dtos.RegistrationDto, *dtos.ErrorDto) {
	if data.Provider != GOOGLE_AUTH_PROVIDER {
		return nil, &dtos.ErrorDto{Message: "system is not support this auth provider"}
	}
	googleProfile, err := getGoogleProfile(data.AccessToken)
	if err != nil {
		return nil, &dtos.ErrorDto{Message: err.Error()}
	}

	if !(googleProfile["email_verified"].(bool)) {
		// this parameter is required to verify specific orgranization/domain based auth
		return nil, &dtos.ErrorDto{Message: "email is not verified for signup"}
	}

	user := &domains.User{
		Name:            googleProfile["name"].(string),
		Phone:           "",
		Email:           googleProfile["email"].(string),
		AuthProvider:    GOOGLE_AUTH_PROVIDER,
		IsEmailVerified: true,
	}

	if err := db.GetInstance().Create(user).Error; err != nil {
		return nil, &dtos.ErrorDto{Message: err.Error()}
	}
	token, err := auth.GenerateToken(user)
	if err != nil {
		return nil, &dtos.ErrorDto{Message: err.Error()}
	}

	userDto := &dtos.UserDto{
		ID:              user.ID,
		Name:            user.Name,
		Email:           user.Email,
		Phone:           user.Phone,
		IsEmailVerified: user.IsEmailVerified,
	}

	response := &dtos.RegistrationDto{}
	response.UserDto = userDto
	response.AccessToken = token

	return response, nil
}

func UserCredentialLogin(data *dtos.CredentialSigninDto) (*dtos.AccessDto, *dtos.ErrorDto) {
	user := &domains.User{}
	if err := db.GetInstance().Table(domains.UserTableName).
		Where("email=? AND auth_provider=?", data.Email, CREDENTIAL_AUTH_PROVIDER).
		First(user).Error; err != nil {
		return nil, &dtos.ErrorDto{Message: err.Error()}
	}

	if !user.CheckIfPasswordIsCorrect(data.Password) {
		return nil, &dtos.ErrorDto{Message: "incorrect password"}
	}
	token, err := auth.GenerateToken(user)
	if err != nil {
		return nil, &dtos.ErrorDto{Message: err.Error()}
	}
	return &dtos.AccessDto{Bearer: token}, nil

}

func UserOAuthLogin(data *dtos.OAuthDto) (*dtos.AccessDto, *dtos.ErrorDto) {
	if data.Provider != GOOGLE_AUTH_PROVIDER {
		return nil, &dtos.ErrorDto{Message: "invalid auth provider"}
	}

	googleProfile, err := getGoogleProfile(data.AccessToken)
	if err != nil {
		return nil, &dtos.ErrorDto{Message: err.Error()}
	}

	user := &domains.User{}
	if err := db.GetInstance().Table(domains.UserTableName).
		Where("email=? AND auth_provider=?", googleProfile["email"], GOOGLE_AUTH_PROVIDER).
		First(user).Error; err != nil {
		return nil, &dtos.ErrorDto{Message: err.Error()}
	}

	token, err := auth.GenerateToken(user)
	if err != nil {
		return nil, &dtos.ErrorDto{Message: err.Error()}
	}
	return &dtos.AccessDto{Bearer: token}, nil
}

func getGoogleProfile(token string) (map[string]interface{}, error) {
	url := fmt.Sprintf("https://www.googleapis.com/oauth2/v3/userinfo?access_token=%s", token)
	httpReq, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	httpClient := &http.Client{}
	httpRes, err := httpClient.Do(httpReq)

	if err != nil {
		return nil, err
	}

	data := make(map[string]interface{})

	err = json.NewDecoder(httpRes.Body).Decode(&data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
