package service

import (
	"github.com/yakuter/go-clean-code/pkg/model"
	"github.com/yakuter/go-clean-code/pkg/repository/post"
)

// PostService ...
type PostService struct {
	PostRepository *post.Repository
}

// NewPostService ...
func NewPostService(p *post.Repository) PostService {
	return PostService{PostRepository: p}
}

// All ...
func (p *PostService) All() ([]model.Post, error) {
	return p.PostRepository.All()
}

// FindByID ...
func (p *PostService) FindByID(id uint) (*model.Post, error) {
	return p.PostRepository.FindByID(id)
}

// Save ...
func (p *PostService) Save(post *model.Post) (*model.Post, error) {
	return p.PostRepository.Save(post)
}

// Delete ...
func (p *PostService) Delete(id uint) error {
	return p.PostRepository.Delete(id)
}

// Migrate ...
func (p *PostService) Migrate() error {
	return p.PostRepository.Migrate()
}
