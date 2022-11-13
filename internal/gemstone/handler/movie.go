package handler

import (
	"gemstone-backend/internal/gemstone/global"
	"github.com/gofiber/fiber/v2"
)

type MovieHandler struct {
	movieService global.IMovieService
}

func NewMovieHandler(movieService global.IMovieService) *MovieHandler {
	return &MovieHandler{movieService: movieService}
}

// Find godoc
// @Title       Find Single Movie
// @Summary     Find Single Movie
// @Description Movie 세부 정보를 조회합니다.
// @Tags        movie
// @Accept      json
// @Param       cid path string true "ContentID"
// @Produce     json
// @Success     200 {object} global.FindMovieRes
// @Failure     400 {object} global.FindMovieRes
// @Router      /api/v1/movie/{cid} [get]
func (m *MovieHandler) Find(c *fiber.Ctx) error {
	var param global.FindMovieReq
	if err := c.ParamsParser(&param); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(global.FindMovieRes{Success: false, Movie: nil})
	}
	res := m.movieService.FindByContentID(param)
	if res.Success == false {
		return c.Status(fiber.StatusInternalServerError).JSON(res)
	}
	return c.Status(fiber.StatusOK).JSON(res)
}

// FindMany godoc
// @Title       Find Many Movie
// @Summary     Find Many Movie
// @Description Movie 조회합니다. Offset 받을 수 있고, Limit 20으로 고정입니다.
// @Tags        movie
// @Accept      json
// @Param       offset query string true "Offset"
// @Produce     json
// @Success     200 {object} global.FindManyMovieRes
// @Failure     400 {object} global.FindManyMovieRes
// @Router      /api/v1/movie [get]
func (m *MovieHandler) FindMany(c *fiber.Ctx) error {
	var param global.FindManyMovieReq
	if err := c.ParamsParser(&param); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(global.FindManyMovieRes{Success: false, Movies: nil})
	}
	res := m.movieService.FindMany(param)
	if res.Success == false {
		return c.Status(fiber.StatusInternalServerError).JSON(res)
	}
	return c.Status(fiber.StatusOK).JSON(res)
}

// Store godoc
// @Title       Store Movie
// @Summary     Store Movie
// @Description 테스트용 API 입니다. Polling 혹은 Listen 시스템 제작 시 삭제 예정입니다.
// @Tags        movie
// @Accept      json
// @Param       StoreMovieReq body global.StoreMovieReq true "Store Req"
// @Produce     json
// @Success     200 {object} global.StoreMovieRes
// @Failure     400 {object} global.StoreMovieRes
// @Router      /api/v1/movie/create [post]
func (m *MovieHandler) Store(c *fiber.Ctx) error {
	var param global.StoreMovieReq
	if err := c.BodyParser(&param); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(global.StoreMovieRes{Success: false, Err: err})
	}
	res := m.movieService.Store(param)
	if res.Success == false {
		return c.Status(fiber.StatusInternalServerError).JSON(res)
	}
	return c.Status(fiber.StatusOK).JSON(res)
}
