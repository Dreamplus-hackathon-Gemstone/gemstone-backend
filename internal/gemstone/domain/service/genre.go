package service

import (
	"gemstone-backend/internal/gemstone/global"
)

type GenreService[P global.Parameter, R global.ReturnType] struct {
	genreRepo global.IRepository[P, R]
}

func (g *GenreService[P, R]) Find(p P) R {
	//TODO implement me
	panic("implement me")
}

func (g *GenreService[P, R]) FindAll(p P) R {
	//TODO implement me
	panic("implement me")
}

func (g *GenreService[P, R]) Store(p P) R {
	//TODO implement me
	panic("implement me")
}

func (g *GenreService[P, R]) Update(p P) R {
	//TODO implement me
	panic("implement me")
}

func (g *GenreService[P, R]) Delete(p P) R {
	//TODO implement me
	panic("implement me")
}
