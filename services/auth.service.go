package services

import (
	"authentication/repositories"
	"authentication/requests"
	"authentication/utils"
	"context"
	"fmt"
	"time"
)

type AuthService interface {
	Login(loginRequest requests.LoginRequest, ctx context.Context) map[string]string
	SendOTPCode(otpRequest requests.OTPRequest, ctx context.Context)
	GetUserProfile(request requests.Profile, ctx context.Context) map[string]string
	ListUsers(ctx context.Context, request requests.UsersList) []map[string]string
}

type authService struct {
	authRepository repositories.AuthRepository
}

func NewAuthService(authRepository repositories.AuthRepository) AuthService {
	return &authService{authRepository: authRepository}
}

func (s *authService) SendOTPCode(otpRequest requests.OTPRequest, ctx context.Context) {
	code := utils.Generate6DigitCode()
	s.authRepository.SetOTP(ctx, otpRequest.PhoneNumber, code, 2*time.Minute)
}

func (s *authService) Login(loginRequest requests.LoginRequest, ctx context.Context) map[string]string {
	otp := s.authRepository.GetOTP(ctx, loginRequest.PhoneNumber)
	if otp != loginRequest.OTPCode {
		panic(utils.PanicMessage{MessageKey: 3})
	}

	var user map[string]string
	exists := s.authRepository.UserExists(ctx, loginRequest.PhoneNumber)
	fmt.Println(exists)
	if exists {
		user = s.authRepository.GetUser(ctx, loginRequest.PhoneNumber)
	} else {
		fmt.Println("I'm creating a user")
		user = s.authRepository.CreateUser(ctx, loginRequest.PhoneNumber)
	}

	accessToken, err := utils.GenerateAccessToken(loginRequest.PhoneNumber, time.Minute*15)
	if err != nil {
		panic(utils.PanicMessage{0, &err})
	}

	refreshToken := utils.GenerateRefreshToken()
	err = s.authRepository.SetRefreshToken(ctx, loginRequest.PhoneNumber, refreshToken, time.Hour*24*7)
	if err != nil {
		panic(utils.PanicMessage{0, &err})
	}

	user["access_token"] = accessToken
	user["refresh_token"] = refreshToken

	return user
}

func (s *authService) GetUserProfile(request requests.Profile, ctx context.Context) map[string]string {
	return s.authRepository.GetUser(ctx, request.PhoneNumber)
}

func (s *authService) ListUsers(ctx context.Context, request requests.UsersList) []map[string]string {
	users := s.authRepository.ListUsers(ctx, request)
	return users
}

//func (s *authService) SearchUsers(ctx context.Context, query string) []map[string]string {
//	users, err := s.authRepository.SearchUsers(ctx, query)
//	if err != nil {
//		panic(err)
//	}
//	return users
//}
