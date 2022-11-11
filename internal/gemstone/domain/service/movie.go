package service

import (
	"gemstone-backend/internal/gemstone/global"
)

type MovieService[P global.Parameter, R global.ReturnType] struct {
	movieRepo global.IRepository[P, R]
}

func (m *MovieService[P, R]) Find(p P) R {
	//TODO implement me
	panic("implement me")
}

func (m *MovieService[P, R]) FindAll(p P) R {
	//TODO implement me
	panic("implement me")
}

func (m *MovieService[P, R]) Store(p P) R {
	//TODO implement me
	panic("implement me")
}

func (m *MovieService[P, R]) Update(p P) R {
	//TODO implement me
	panic("implement me")
}

func (m *MovieService[P, R]) Delete(p P) R {
	//TODO implement me
	panic("implement me")
}
