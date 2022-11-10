package repository

import (
	"gemstone-backend/internal/gemstone/transactor"
	"github.com/gofiber/fiber/v2"
)

type MakerRepository struct {
	transactor transactor.ITransactor
}

func NewMakerRepository(ITransactor transactor.ITransactor) *MakerRepository {
	return &MakerRepository{transactor: ITransactor}
}

func (m *MakerRepository) Find(c *fiber.Ctx) error {
	m.transactor.Execute(c, func(c *fiber.Ctx) error {
		// ...
	})
}

func (m *MakerRepository) Store() error {
	//TODO implement me
	panic("implement me")
}

func (m *MakerRepository) Update() error {
	//TODO implement me
	panic("implement me")
}

func (m *MakerRepository) FindAll() error {
	//TODO implement me
	panic("implement me")
}

func (m *MakerRepository) Delete() error {
	//TODO implement me
	panic("implement me")
}
