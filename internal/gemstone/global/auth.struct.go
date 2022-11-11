package global

import "gemstone-backend/internal/gemstone/domain"

// Interface for AuthHandler
type (
	VerifyNicknameReq struct {
		Nickname string `json:"nickname,omitempty"`
	}

	FindMinerReq struct {
		Account string `json:"account,omitempty"`
	}

	RegisterMakerReq struct {
		Name        string `json:"name,omitempty"`
		Nickname    string `json:"nickname,omitempty"`
		PhoneNumber string `json:"phone_number,omitempty"`
		HomeAddress string `json:"home_address,omitempty"`
		Email       string `json:"email,omitempty"`
		Password    string `json:"password,omitempty"`
	}

	RegisterMinerReq struct {
		Account  string `json:"account,omitempty"`
		Nickname string `json:"nickname,omitempty"`
	}

	MakerLoginReq struct {
		Email    string `json:"email,omitempty"`
		Password string `json:"password,omitempty"`
	}

	MinerLoginReq struct {
		Account string `json:"account,omitempty"`
		AuthID  string `json:"auth_id,omitempty"`
	}

	UpdateMakerNicknameReq struct {
		ID       uint   `json:"id,omitempty"`
		Nickname string `json:"nickname,omitempty"`
	}
)

type (
	FindMinerRes struct {
		Success bool           `json:"success,omitempty"`
		Miner   *domain.Miners `json:"miner,omitempty"`
	}

	VerifyNicknameRes struct {
		Success bool  `json:"success,omitempty"`
		Err     error `json:"err,omitempty"`
	}

	RegisterMinerRes struct {
		Success bool  `json:"success,omitempty"`
		Err     error `json:"err,omitempty"`
	}

	RegisterMakerRes struct {
		Success bool  `json:"success,omitempty"`
		Err     error `json:"err,omitempty"`
	}

	LoginRes struct {
		Success bool  `json:"success,omitempty"`
		Err     error `json:"err,omitempty"`
	}

	UpdateMakerNicknameRes struct {
		Success bool  `json:"success,omitempty"`
		Err     error `json:"err,omitempty"`
	}
)
