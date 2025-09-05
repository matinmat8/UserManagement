package requests

type LoginRequest struct {
	PhoneNumber string `json:"phoneNumber" binding:"required"`
	OTPCode     string `json:"OTPCode" binding:"required,min=6"`
}

type OTPRequest struct {
	PhoneNumber string `json:"phoneNumber" binding:"required"`
}

type Profile struct {
	PhoneNumber string `json:"phone" form:"phone"`
}

type UsersList struct {
	Page      int64  `form:"page" binding:"required,min=1"`
	PageSize  int64  `form:"page_size" binding:"required,min=1,max=100"`
	PhoneLike string `form:"phone"`
}
