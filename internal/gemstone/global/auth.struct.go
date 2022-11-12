package global

// Interface for AuthHandler
type (
	VerifyNicknameReq struct {
		Nickname string `json:"nickname,omitempty"`
	}

	RegisterReq struct {
		Name     string `json:"name,omitempty"`
		Nickname string `json:"nickname,omitempty"`
		Account  string `json:"account,omitempty"`
		Email    string `json:"email,omitempty"`
		Password string `json:"password,omitempty"`
	}

	LoginReq struct {
		Email    string `json:"email,omitempty"`
		Password string `json:"password,omitempty"`
	}

	UpdateNicknameReq struct {
		ID       uint   `json:"id,omitempty"`
		Nickname string `json:"nickname,omitempty"`
	}
)

type (
	VerifyNicknameRes struct {
		Success bool  `json:"success,omitempty"`
		Err     error `json:"err,omitempty"`
	}

	RegisterRes struct {
		Success bool  `json:"success,omitempty"`
		Err     error `json:"err,omitempty"`
	}

	LoginRes struct {
		Success bool  `json:"success,omitempty"`
		Err     error `json:"err,omitempty"`
	}

	UpdateNicknameRes struct {
		Success bool  `json:"success,omitempty"`
		Err     error `json:"err,omitempty"`
	}
)
