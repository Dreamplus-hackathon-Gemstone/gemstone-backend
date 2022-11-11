package service

import "gemstone-backend/internal/gemstone/global"

type TokenService[P global.Parameter, R global.ReturnType] struct {
	tokenRepo global.IRepository[P, R]
}

func (t *TokenService[P, R]) Find(p P) R {
	//TODO implement me
	panic("implement me")
}

func (t *TokenService[P, R]) FindAll(p P) R {
	//TODO implement me
	panic("implement me")
}

func (t *TokenService[P, R]) Store(p P) R {
	//TODO implement me
	panic("implement me")
}

func (t *TokenService[P, R]) Update(p P) R {
	//TODO implement me
	panic("implement me")
}

func (t *TokenService[P, R]) Delete(p P) R {
	//TODO implement me
	panic("implement me")
}
