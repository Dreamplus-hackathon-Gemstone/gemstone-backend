package handler

import (
	"gemstone-backend/internal/gemstone/global"
	"github.com/gofiber/fiber/v2"
	"log"
	_ "time"
)

type ProposalHandler struct {
	proposalService global.IProposalService
}

func NewProposalHandler(proposalService global.IProposalService) global.IProposalHandler {
	return &ProposalHandler{proposalService: proposalService}
}

// Find godoc
// @Title       Find Single Proposal
// @Summary     Find Single Proposal
// @Description Proposal 세부 정보를 조회합니다.
// @Tags        proposal
// @Accept      json
// @Param       cid path string true "ContentID"
// @Produce     json
// @Success     200 {object} global.FindProposalRes
// @Failure     400 {object} global.FindProposalRes
// @Router      /api/v1/proposal/{cid} [get]
func (p *ProposalHandler) Find(c *fiber.Ctx) error {
	var param = global.FindProposalReq{}
	if err := c.ParamsParser(&param); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(global.FindProposalRes{Success: false, Proposal: nil, Err: err})
	}
	res := p.proposalService.Find(param)
	if res.Success == false {
		return c.Status(fiber.StatusInternalServerError).JSON(res)
	}
	return c.Status(fiber.StatusOK).JSON(res)
}

// FindMany godoc
// @Title       Find Many Proposal
// @Summary     Find Many Proposal
// @Description Proposal 조회합니다. Offset 받을 수 있고, Limit 20으로 고정입니다.
// @Tags        proposal
// @Accept      json
// @Param       offset query string true "Offset"
// @Produce     json
// @Success     200 {object} global.FindManyProposalRes
// @Failure     400 {object} global.FindManyProposalRes
// @Router      /api/v1/proposal [get]
func (p *ProposalHandler) FindMany(c *fiber.Ctx) error {
	var param = global.FindManyProposalReq{}
	if err := c.QueryParser(&param); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(global.FindManyProposalRes{Success: false, Proposals: nil})
	}
	res := p.proposalService.FindMany(param)
	if res.Success == false {
		return c.Status(fiber.StatusInternalServerError).JSON(res)
	}
	return c.Status(fiber.StatusOK).JSON(res)
}

// Store godoc
// @Title       Store Proposal
// @Summary     Store Proposal
// @Description 테스트용 API 입니다. Polling 혹은 Listen 시스템 제작 시 삭제 예정입니다.
// @Tags        proposal
// @Accept      json
// @Param       StoreProposalReq body global.StoreProposalReq true "Store Req"
// @Produce     json
// @Success     200 {object} global.StoreProposalRes
// @Failure     400 {object} global.StoreProposalRes
// @Router      /api/v1/proposal/store [post]
func (p *ProposalHandler) Store(c *fiber.Ctx) error {
	var param = global.StoreProposalReq{}
	if err := c.BodyParser(&param); err != nil {
		log.Println("err :", err)
		log.Println("Body :", string(c.Body()))
		return c.Status(fiber.StatusBadRequest).JSON(global.StoreProposalRes{Success: false, Err: err})
	}
	res := p.proposalService.Store(param)
	if res.Success == false {
		return c.Status(fiber.StatusInternalServerError).JSON(res)
	}
	return c.Status(fiber.StatusOK).JSON(res)
}

func (p *ProposalHandler) Update(c *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func (p *ProposalHandler) Delete(c *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}
