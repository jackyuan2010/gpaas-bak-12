package dto

// User register structure
type RegisterDto struct {
	Mobile     string `json:"mobile"`
	Password     string `json:"password"`
	UserName     string `json:"userName"`
	HeaderImg    string `json:"headerImg"`
}

// User login structure
type LoginDto struct {
	Mobile  string `json:"mobile"`  // 用户名
	Password  string `json:"password"`  // 密码
	Captcha   string `json:"captcha"`   // 验证码
	CaptchaId string `json:"captchaId"` // 验证码ID
}


// Modify password structure
type ChangePasswordDto struct {
	Id          uint   `json:"-"`           // 从 JWT 中提取 user id，避免越权
	Password    string `json:"password"`    // 密码
	NewPassword string `json:"newPassword"` // 新密码
}