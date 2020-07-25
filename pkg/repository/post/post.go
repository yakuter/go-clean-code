package post

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/yakuter/go-clean-code/pkg/model"
)

// Repository ...
type Repository struct {
	db *gorm.DB
}

// NewRepository ...
func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

// All ...
func (p *Repository) Test() {
	fmt.Println("erhan")
}

// All ...
func (p *Repository) All() ([]model.Post, error) {
	posts := []model.Post{}
	err := p.db.Find(&posts).Error
	return posts, err
}

// FindByID ...
func (p *Repository) FindByID(id uint) (*model.Post, error) {
	post := new(model.Post)
	err := p.db.Where(`id = ?`, id).First(&post).Error
	return post, err
}

// Save ...
func (p *Repository) Save(post *model.Post) (*model.Post, error) {
	err := p.db.Save(&post).Error
	return post, err
}

// Delete ...
func (p *Repository) Delete(id uint) error {
	err := p.db.Delete(&model.Post{ID: id}).Error
	return err
}

// Migrate ...
func (p *Repository) Migrate() error {
	return p.db.AutoMigrate(&model.Post{}).Error
}
