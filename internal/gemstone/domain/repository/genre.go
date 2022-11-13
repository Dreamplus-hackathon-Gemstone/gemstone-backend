package repository

import (
	"gemstone-backend/internal/gemstone/domain"
	"gemstone-backend/internal/gemstone/global"
	"gorm.io/gorm"
	"log"
)

type GenreRepo struct {
	db *gorm.DB
}

func (g *GenreRepo) FindMovieByGenre(param global.FindMovieByGenreReq) (ret global.FindMovieByGenreRes) {
	var movies []*domain.Movie
	tx := g.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	log.Println("GenreID = ", param.GenreID)
	log.Println("Offset = ", param.Offset)
	if tx.Error != nil {
		return global.FindMovieByGenreRes{Success: false, Err: tx.Error, Movies: nil}
	}
	if err := tx.Table("movies").Where(&domain.Movie{GenreID: param.GenreID}).Find(&movies).Offset(20 * param.Offset).Limit(20); err.Error != nil {
		return global.FindMovieByGenreRes{Success: false, Err: tx.Error, Movies: nil}
	}
	return global.FindMovieByGenreRes{Success: true, Err: tx.Commit().Error, Movies: movies}
}

func (g *GenreRepo) FindProposalByGenre(param global.FindProposalByGenreReq) (ret global.FindProposalByGenreRes) {
	tx := g.db.Begin()
	var proposals []*domain.Proposal
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if tx.Error != nil {
		return global.FindProposalByGenreRes{Success: false, Err: tx.Error, Proposals: nil}
	}
	if err := tx.Table("proposals").Where(&domain.Proposal{GenreID: param.GenreID}).Find(&proposals).Offset(20 * param.Offset).Limit(20); err.Error != nil {
		return global.FindProposalByGenreRes{Success: false, Err: tx.Error, Proposals: nil}
	}
	return global.FindProposalByGenreRes{Success: true, Err: tx.Commit().Error, Proposals: proposals}
}

func (g *GenreRepo) Store(param global.StoreGenreReq) (ret global.StoreGenreRes) {
	tx := g.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if tx.Error != nil {
		return global.StoreGenreRes{Success: false, Err: tx.Error}
	}
	if err := tx.Create(&domain.Genres{GenreType: param.GenreType}); err.Error != nil {
		return global.StoreGenreRes{Success: false, Err: err.Error}
	}
	return global.StoreGenreRes{Success: true, Err: tx.Commit().Error}
}

func NewGenreRepo(db *gorm.DB) global.IGenreRepo {
	return &GenreRepo{db: db}
}
