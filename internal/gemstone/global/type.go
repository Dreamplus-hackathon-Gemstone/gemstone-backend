package global

type IService[P Parameter, R ReturnType] interface {
	Find(param P) (ret R)
	FindAll(param P) (ret R)
	Store(param P) (ret R)
	Update(param P) (ret R)
	Delete(param P) (ret R)
}

type IRepository[P Parameter, R ReturnType] interface {
	Find(param P) (ret R)
	FindAll(param P) (ret R)
	Store(param P) (ret R)
	Update(param P) (ret R)
	Delete(param P) (ret R)
}

// 외부에서 사용되는 인터페이스
type IExample interface {
	Find(param Parameter) (ret ReturnType)
	FindAll(param Parameter) (ret ReturnType)
	Store(param Parameter) (ret ReturnType)
	Update(param Parameter) (ret ReturnType)
	Delete(param Parameter) (ret ReturnType)
}

// 세부 구현 인터페이스
type IGenreRepo interface {
	Find(param string) (ret ReturnType)
	FindAll(param Parameter) (ret ReturnType)
	Store(param Parameter) (ret ReturnType)
	Update(param Parameter) (ret ReturnType)
	Delete(param Parameter) (ret ReturnType)
}

type (
	Parameter interface {
		string | VerifyNicknameReq | GetProposal
	}

	ReturnType struct {
		success bool
		err     error
		data    interface{}
	}

	VerifyNicknameReq interface {
		string
	}

	GetProposal interface {
		uint
	}
)
