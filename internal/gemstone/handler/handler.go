package handler

import (
	"gemstone-backend/internal/gemstone/domain/service"
	"gemstone-backend/internal/gemstone/global"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	UserHandler     global.IUserHandler
	ProposalHandler global.IProposalHandler
	MovieHandler    global.IMovieHandler
	GenreHandler    global.IGenreHandler
}

func NewHandler(svc *service.Service) *Handler {
	user := NewUserHandler(svc.UserService)
	proposal := NewProposalHandler(svc.ProposalService)
	movie := NewMovieHandler(svc.MovieService)
	genre := NewGenreHandler(svc.GenreService)
	return &Handler{
		UserHandler:     user,
		ProposalHandler: proposal,
		MovieHandler:    movie,
		GenreHandler:    genre,
	}
}

func (h *Handler) Ping(c *fiber.Ctx) error {
	return c.SendString("pong")
}
