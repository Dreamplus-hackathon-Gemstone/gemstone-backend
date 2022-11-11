package service

import (
	"gemstone-backend/internal/gemstone/global"
)

type MakerService[P global.Parameter, R global.ReturnType] struct {
	makerRepo global.IRepository[P, R]
}

func (m *MakerService[P, R]) Find(p P) R {
	//TODO implement me
	panic("implement me")
}

func (m *MakerService[P, R]) FindAll(p P) R {
	//TODO implement me
	panic("implement me")
}

func (m *MakerService[P, R]) Store(p P) R {
	//TODO implement me
	panic("implement me")
}

func (m *MakerService[P, R]) Update(p P) R {
	//TODO implement me
	panic("implement me")
}

func (m *MakerService[P, R]) Delete(p P) R {
	//TODO implement me
	panic("implement me")
}
