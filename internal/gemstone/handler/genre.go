package handler

import (
	"gemstone-backend/internal/gemstone/global"
	"github.com/gofiber/fiber/v2"
	"log"
)

type GenreHandler struct {
	genreService global.IGenreService
}

func NewGenreHandler(genreService global.IGenreService) global.IGenreHandler {
	return &GenreHandler{genreService: genreService}
}

// FindMovieByGenre godoc
// @Title       Find Movie By Genre ID
// @Summary     Find Movie By Genre ID
// @Description Genre 별 Movie 조회합니다. Offset 받을 수 있고, Limit 20으로 고정입니다.
// @Tags        genre
// @Accept      json
// @Param       offset query string true "Offset"
// @Param       genreId query string true "GenreID"
// @Produce     json
// @Success     200 {object} global.FindMovieByGenreRes
// @Failure     400 {object} global.FindMovieByGenreRes
// @Router      /api/v1/genre/movie [get]
func (g *GenreHandler) FindMovieByGenre(c *fiber.Ctx) error {
	var param global.FindMovieByGenreReq
	if err := c.QueryParser(&param); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(global.FindMovieByGenreRes{Success: false, Movies: nil, Err: err})
	}
	log.Println("Param : ", param)
	res := g.genreService.FindMovie(param)
	if res.Success == false {
		return c.Status(fiber.StatusInternalServerError).JSON(res)
	}
	return c.Status(fiber.StatusOK).JSON(res)
}

// FindProposalByGenre godoc
// @Title       Find Proposal By Genre ID
// @Summary     Find Proposal By Genre ID
// @Description Genre 별 Proposal 조회합니다. Offset 받을 수 있고, Limit 20으로 고정입니다.
// @Tags        genre
// @Accept      json
// @Param       offset query string true "Offset"
// @Param       genreId query string true "GenreID"
// @Produce     json
// @Success     200 {object} global.FindProposalByGenreRes
// @Failure     400 {object} global.FindProposalByGenreRes
// @Router      /api/v1/genre/proposal [get]
func (g *GenreHandler) FindProposalByGenre(c *fiber.Ctx) error {
	var param global.FindProposalByGenreReq
	if err := c.QueryParser(&param); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(global.FindProposalByGenreRes{Success: false, Proposals: nil, Err: err})
	}
	res := g.genreService.FindProposal(param)
	if res.Success == false {
		return c.Status(fiber.StatusInternalServerError).JSON(res)
	}
	return c.Status(fiber.StatusOK).JSON(res)
}

// Store godoc
// @Title       Store Genre
// @Summary     Store Genre
// @Description 테스트용 API 입니다.
// @Tags        genre
// @Accept      json
// @Param       StoreGenreReq body global.StoreGenreReq true "Store Req"
// @Produce     json
// @Success     200 {object} global.StoreGenreRes
// @Failure     400 {object} global.StoreGenreRes
// @Router      /api/v1/genre/create [post]
func (g *GenreHandler) Store(c *fiber.Ctx) error {
	var param global.StoreGenreReq
	if err := c.BodyParser(&param); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(global.StoreGenreRes{Success: false, Err: err})
	}
	res := g.genreService.Store(param)
	if res.Success == false {
		return c.Status(fiber.StatusInternalServerError).JSON(res)
	}
	return c.Status(fiber.StatusOK).JSON(res)
}
