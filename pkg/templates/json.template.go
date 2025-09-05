package MessageTemplate

import (
	"github.com/gin-gonic/gin"
)

var MessageTemplates = map[int]struct {
	Status  int
	Message gin.H
}{
	0: {400, gin.H{"en_message": "An error occurred", "fa_message": "خطایی پیش آمد"}},
	1: {401, gin.H{"en_message": "User not authenticated", "fa_message": "کاربر احراز هویت نشد"}},
	2: {400, gin.H{"en_message": "OTP code has expired", "fa_message": "کد تایید منقضی شده است"}},
	3: {400, gin.H{"en_message": "OTP code is wrong", "fa_message": "کد تایید نادرست است"}},
	4: {400, gin.H{"en_message": "No user found with this phone number", "fa_message": "کاربری با این شماره تماس پیدا نشد"}},
	5: {400, gin.H{"en_message": "The OTP code has sent before", "fa_message": "کد تایید از قبل ارسال شده است"}},
	6: {400, gin.H{"en_message": "Too many requests. Please try again later.", "fa_message": "درخواست بیش از حد لطفا چند لحظه بعد دوباره تلاش کنید"}},
}
