package global

// Interface for AuthHandler
type (
	VerifyNicknameReq struct {
		Nickname string `json:"nickname"`
	}

	RegisterReq struct {
		Name     string `json:"name"`
		Nickname string `json:"nickname"`
		Account  string `json:"account"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	LoginReq struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	UpdateNicknameReq struct {
		ID       uint   `json:"id"`
		Nickname string `json:"nickname"`
	}
)

type (
	VerifyNicknameRes struct {
		Success bool  `json:"success"`
		Err     error `json:"err"`
	}

	RegisterRes struct {
		Success bool  `json:"success"`
		Err     error `json:"err"`
	}

	LoginRes struct {
		Success bool  `json:"success"`
		Err     error `json:"err"`
	}

	UpdateNicknameRes struct {
		Success bool  `json:"success"`
		Err     error `json:"err"`
	}
)
