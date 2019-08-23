package repository

import (
	"context"
	"github.com/YuukiRen/GO_Api/models"
)

// Article repo is for...
type ArticleRepo interface{
	Fetch(ctx context.Context, num int64) ([]*models.Article,error)
	GetByID(ctx context.Context, id int64) (*models.Article,error)
	Create(ctx context.Context, p *models.Article) (int64,error)
	Update(ctx context.Context, p *models.Article) (*models.Article,error)
	Delete(ctx context.Context, id int64) (bool,error)
}