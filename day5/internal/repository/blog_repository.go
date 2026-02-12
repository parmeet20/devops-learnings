package repository

import (
	"context"
	"time"

	models "github.com/parmeet20/bloghive/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type BlogRepository interface {
	Create(ctx context.Context, blog *models.Blog) error
	GetAll(ctx context.Context) ([]models.Blog, error)
	GetByID(ctx context.Context, id string) (*models.Blog, error)
	Update(ctx context.Context, id string, blog *models.Blog) error
}

type blogRepository struct {
	collection *mongo.Collection
}

func NewBlogRepository(db *mongo.Database) BlogRepository {
	return &blogRepository{
		collection: db.Collection("blogs"),
	}
}

func (r *blogRepository) Create(ctx context.Context, blog *models.Blog) error {
	blog.CreatedAt = time.Now()
	blog.UpdatedAt = time.Now()
	_, err := r.collection.InsertOne(ctx, blog)
	return err
}

func (r *blogRepository) GetAll(ctx context.Context) ([]models.Blog, error) {
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var blogs []models.Blog
	if err := cursor.All(ctx, &blogs); err != nil {
		return nil, err
	}
	return blogs, nil
}

func (r *blogRepository) GetByID(ctx context.Context, id string) (*models.Blog, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var blog models.Blog
	err = r.collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&blog)
	if err != nil {
		return nil, err
	}
	return &blog, nil
}

func (r *blogRepository) Update(ctx context.Context, id string, blog *models.Blog) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	update := bson.M{
		"$set": bson.M{
			"title":      blog.Title,
			"content":    blog.Content,
			"author":     blog.Author,
			"updated_at": time.Now(),
		},
	}

	_, err = r.collection.UpdateByID(ctx, objID, update)
	return err
}
