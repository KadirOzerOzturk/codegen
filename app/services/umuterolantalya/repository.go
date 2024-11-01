package umuterolantalya

import (
	"github.com/google/uuid"
	"github.com/KadirOzerOzturk/deneme/app/entities"
	"github.com/KadirOzerOzturk/deneme/internal/paginator"
	"github.com/KadirOzerOzturk/deneme/internal/search"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Repository interface {
	Insert(payload *entities.Umuterolantalya) error
	List(page, per_page int, sort string, query string, filter []entities.Filter) (*paginator.Data, error)
	Get(id *uuid.UUID) (*entities.Umuterolantalya, error)
	Update(id *uuid.UUID, payload *entities.Umuterolantalya) error
	Delete(id *uuid.UUID) error
}

type repository struct {
	DB *gorm.DB
}

var _ Repository = &repository{}

func NewRepo(DB *gorm.DB) Repository {
	return &repository{
		DB: DB,
	}
}

// Insert repository
func (r *repository) Insert(payload *entities.Umuterolantalya) error {
	if err := r.DB.Create(payload).Error; err != nil {
		return err
	}
	return nil
}

// List repository
func (r *repository) List(page, per_page int, sort string, query string, filter []entities.Filter) (*paginator.Data, error) {
	items := []*entities.Umuterolantalya{}
	db := r.DB.Model(&entities.Umuterolantalya{}).Preload(clause.Associations)
	if query != "" {
		search.Search(query, db)
	}

	result, err := paginator.New(db, page, per_page, sort, filter).Paginate(&items)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// Get repository
func (r *repository) Get(id *uuid.UUID) (*entities.Umuterolantalya, error) {
	item := entities.Umuterolantalya{}
	if err := r.DB.Model(&item).First(&item, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return &item, nil
}

// Update repository
func (r *repository) Update(id *uuid.UUID, payload *entities.Umuterolantalya) error {
	item, err := r.Get(id)
	if err != nil {
		return err
	}

	if err := r.DB.Model(&item).Updates(&payload).Error; err != nil {
		return err
	}

	return nil
}

// Delete repository
func (r *repository) Delete(id *uuid.UUID) error {
	if err := r.DB.Where("id = ?", id).Delete(&entities.Umuterolantalya{}).Error; err != nil {
		return err
	}

	return nil
}