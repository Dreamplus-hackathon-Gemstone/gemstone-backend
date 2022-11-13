package repository

import (
	"gemstone-backend/internal/gemstone/domain"
	"gemstone-backend/internal/gemstone/global"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"log"
	"time"
)

type MovieRepo struct {
	db *gorm.DB
}

func (m *MovieRepo) TableName() string {
	return "movies"
}

func (m *MovieRepo) FindByContentID(param global.FindMovieReq) (ret global.FindMovieRes) {
	var movie *domain.Movie
	tx := m.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if tx.Error != nil {
		return global.FindMovieRes{Success: false, Movie: nil, Err: tx.Error}
	}
	if err := tx.Table(m.TableName()).Preload(clause.Associations).Where(&domain.Movie{ContentID: param.CID}).Find(&movie); err.Error != nil {
		tx.Rollback()
		return global.FindMovieRes{Success: false, Movie: nil, Err: err.Error}
	}
	return global.FindMovieRes{Success: true, Movie: movie, Err: tx.Commit().Error}
}

func (m *MovieRepo) FindMany(param global.FindManyMovieReq) (ret global.FindManyMovieRes) {
	var movies []*domain.Movie
	tx := m.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if tx.Error != nil {
		return global.FindManyMovieRes{Success: false, Movies: nil, Err: tx.Error}
	}
	if err := tx.Table(m.TableName()).Preload(clause.Associations).Where(&domain.Movie{GenreID: param.GenreID}).Find(&movies).Offset(20 * param.Offset).Limit(20); err.Error != nil {
		tx.Rollback()
		return global.FindManyMovieRes{Success: false, Movies: nil, Err: err.Error}
	}
	return global.FindManyMovieRes{Success: true, Movies: movies, Err: tx.Commit().Error}
}

func (m *MovieRepo) Store(param global.StoreMovieReq) (ret global.StoreMovieRes) {
	tx := m.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if tx.Error != nil {
		return global.StoreMovieRes{Success: false, Err: tx.Error}
	}
	if err := tx.Create(&domain.Movie{
		MakerID:      param.MakerID,
		GenreID:      param.GenreID,
		ContentID:    param.ContentID,
		Description:  param.Description,
		ThumbnailURI: param.ThumbnailURI,
		MovieURI:     param.MovieURI,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}); err.Error != nil {
		log.Printf("Error loading while create movies ... err : %v\n", err)
		tx.Rollback()
		return global.StoreMovieRes{Success: false, Err: err.Error}
	}
	return global.StoreMovieRes{Success: true, Err: tx.Commit().Error}
}

func NewMovieRepo(db *gorm.DB) global.IMovieRepo {
	return &MovieRepo{db: db}
}
