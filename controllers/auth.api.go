package controllers

import (
	"authentication/requests"
	"authentication/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthAPI interface {
	CheckDTO(context *gin.Context, dto interface{})
	Login(context *gin.Context)
	SendOTP(context *gin.Context)
	Profile(context *gin.Context)
	ListUsers(c *gin.Context)
}

type authAPI struct {
	authService services.AuthService
}

func NewAuthAPI(authService services.AuthService) AuthAPI {
	return &authAPI{authService}
}

func (api authAPI) CheckDTO(context *gin.Context, dto interface{}) {
	err := context.ShouldBind(&dto)
	if err != nil {
		panic(err)
	}
}

// SendOTP godoc
// @Summary Send OTP code to phone number
// @Description Generates and sends OTP, stores it in Redis
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body requests.OTPRequest true "OTP request"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Router /api/v1/auth/send/otp [post]
func (api authAPI) SendOTP(context *gin.Context) {
	var otpRequest requests.OTPRequest
	api.CheckDTO(context, &otpRequest)

	api.authService.SendOTPCode(otpRequest, context)

	context.JSON(http.StatusOK, gin.H{
		"fa_message": "کد یکبارمصرف با موفقیت ارسال شد",
		"en_message": "OTP code sent successfully",
	})
}

// Login godoc
// @Summary Login with phone number and OTP
// @Description Verify OTP, create user if not exists, and return JWT tokens
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body requests.LoginRequest true "Login request"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Router /api/v1/auth/login [post]
func (api authAPI) Login(context *gin.Context) {
	var loginRequest requests.LoginRequest
	api.CheckDTO(context, &loginRequest)

	user := api.authService.Login(loginRequest, context)

	context.JSON(http.StatusOK, gin.H{
		"fa_message": "ورود با موفقیت انجام شد",
		"en_message": "Login successful",
		"user":       user,
	})
}

// Profile godoc
// @Summary Get user profile
// @Description Retrieve user profile by query parameters
// @Tags Auth
// @Accept json
// @Produce json
// @Param phone query string true "Phone number"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Router /api/v1/auth/profile [get]
func (api authAPI) Profile(context *gin.Context) {
	var profileRequest requests.Profile

	if err := context.ShouldBindQuery(&profileRequest); err != nil {
		panic(err)
	}

	user := api.authService.GetUserProfile(profileRequest, context)

	context.JSON(200, gin.H{"user": user})
}

// ListUsers godoc
// @Summary List users
// @Description Paginated list of users with optional phone search
// @Tags Auth
// @Accept json
// @Produce json
// @Param page query int true "Page number"
// @Param page_size query int true "Number of users per page"
// @Param phone query string false "Search by phone"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Router /api/v1/auth/users [get]
func (api authAPI) ListUsers(c *gin.Context) {
	var request requests.UsersList
	if err := c.ShouldBindQuery(&request); err != nil {
		panic(err)
	}

	users := api.authService.ListUsers(c, request)
	c.JSON(200, gin.H{
		"page":      request.Page,
		"page_size": request.PageSize,
		"search":    request.PhoneLike,
		"users":     users,
	})
}
