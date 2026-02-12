package service

import (
	"context"
	"time"

	models "github.com/parmeet20/bloghive/internal/model"
	"github.com/parmeet20/bloghive/internal/repository"
)

type BlogService interface {
	Create(blog *models.Blog) error
	GetAll() ([]models.Blog, error)
	GetByID(id string) (*models.Blog, error)
	Update(id string, blog *models.Blog) error
}

type blogService struct {
	repo repository.BlogRepository
}

func NewBlogService(repo repository.BlogRepository) BlogService {
	return &blogService{repo}
}

func (s *blogService) Create(blog *models.Blog) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return s.repo.Create(ctx, blog)
}

func (s *blogService) GetAll() ([]models.Blog, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return s.repo.GetAll(ctx)
}

func (s *blogService) GetByID(id string) (*models.Blog, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return s.repo.GetByID(ctx, id)
}

func (s *blogService) Update(id string, blog *models.Blog) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return s.repo.Update(ctx, id, blog)
}
