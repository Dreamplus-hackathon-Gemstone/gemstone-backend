package repository

import (
	"gemstone-backend/internal/gemstone/domain"
	"gemstone-backend/internal/gemstone/global"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"log"
	"time"
)

type ProposalRepo struct {
	db *gorm.DB
}

func (p *ProposalRepo) TableName() string {
	return "proposals"
}

func (p *ProposalRepo) FindByContentID(param global.FindProposalReq) (ret global.FindProposalRes) {
	var proposal *domain.Proposal
	tx := p.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if tx.Error != nil {
		return global.FindProposalRes{Success: false, Proposal: nil, Err: tx.Error}
	}
	if err := tx.Table(p.TableName()).Preload(clause.Associations).Where(&domain.Proposal{ContentID: param.CID}).Find(&proposal); err.Error != nil {
		return global.FindProposalRes{Success: false, Proposal: nil, Err: err.Error}
	}
	return global.FindProposalRes{Success: true, Proposal: proposal, Err: tx.Commit().Error}
}

func (p *ProposalRepo) FindMany(param global.FindManyProposalReq) (ret global.FindManyProposalRes) {
	var proposals []*domain.Proposal
	tx := p.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if tx.Error != nil {
		return global.FindManyProposalRes{Success: false, Proposals: nil}
	}
	if err := tx.Table(p.TableName()).Preload(clause.Associations).Find(&proposals).Offset(param.Offset * 20).Limit(20); err.Error != nil {
		tx.Rollback()
		return global.FindManyProposalRes{Success: false, Proposals: nil}
	}
	tx.Commit()
	return global.FindManyProposalRes{Success: true, Proposals: proposals}
}

func (p *ProposalRepo) Store(param global.StoreProposalReq) (ret global.StoreProposalRes) {
	tx := p.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if tx.Error != nil {
		return global.StoreProposalRes{Success: false, Err: tx.Error}
	}
	if err := tx.Create(&domain.Proposal{
		GenreID:                 param.GenreID,
		MakerID:                 param.MakerID,
		ContentID:               param.ContentID,
		ThumbnailURI:            param.ThumbnailURI,
		DocumentURI:             param.DocumentURI,
		Title:                   param.Title,
		Description:             param.Description,
		TargetPrice:             param.TargetPrice,
		Deadline:                param.Deadline,
		ExpectedReleaseYear:     param.ExpectedReleaseYear,
		EstimatedProductionTime: param.EstimatedProductionTime,
		CreatedAt:               time.Now().UTC(),
		UpdatedAt:               time.Now().UTC(),
	}); err.Error != nil {
		log.Printf("Error loading while create proposals ... err : %v\n", err)
		tx.Rollback()
		return global.StoreProposalRes{Success: false, Err: err.Error}
	}
	return global.StoreProposalRes{Success: true, Err: tx.Commit().Error}
}

func (p *ProposalRepo) Update(param global.UpdateProposalReq) (ret global.UpdateProposalRes) {
	//TODO implement me
	panic("implement me")
}

func (p *ProposalRepo) Delete(param global.DeleteProposalReq) (ret global.DeleteProposalRes) {
	//TODO implement me
	panic("implement me")
}

func NewProposalRepo(db *gorm.DB) global.IProposalRepo {
	return &ProposalRepo{db: db}
}
