package interfaces

import (
	"context"
	"github.com/Aerok925/LeoNews/internal/models"
)

type NewsDB interface {
	GetNews(ctx context.Context) ([]models.News, error)
	CreateNews(ctx context.Context, news models.News) (*models.News, error)
}
