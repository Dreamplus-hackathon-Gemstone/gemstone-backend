package service

import (
	"gemstone-backend/internal/gemstone/global"
)

type MinerService[P global.Parameter, R global.ReturnType] struct {
	minerRepo global.IRepository[P, R]
}

func (m *MinerService[P, R]) Find(p P) R {
	//TODO implement me
	panic("implement me")
}

func (m *MinerService[P, R]) FindAll(p P) R {
	//TODO implement me
	panic("implement me")
}

func (m *MinerService[P, R]) Store(p P) R {
	//TODO implement me
	panic("implement me")
}

func (m *MinerService[P, R]) Update(p P) R {
	//TODO implement me
	panic("implement me")
}

func (m *MinerService[P, R]) Delete(p P) R {
	//TODO implement me
	panic("implement me")
}
