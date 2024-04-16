package service

import (
	"context"
	"github.com/Aerok925/LeoNews/internal/interfaces"
	"github.com/Aerok925/LeoNews/internal/models"
	"log"
	"math/rand"
)

var (
	NewsService *News
)

type News struct {
	db            interfaces.NewsDB
	subscribeNews map[int]chan *models.News
	// TODO: logger
}

func Init(db interfaces.NewsDB) (*News, error) {
	NewsService = &News{db: db, subscribeNews: make(map[int]chan *models.News)}
	return NewsService, nil
}

func (n *News) Create(ctx context.Context, news models.News) (*models.News, error) {
	respNews, err := n.db.CreateNews(ctx, news)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	log.Println(news)
	for _, c := range n.subscribeNews {
		c <- respNews
	}
	return respNews, nil
}

func (n *News) GetAll(ctx context.Context) ([]models.News, error) {
	newsResp, err := n.db.GetNews(ctx)
	if err != nil {
		return nil, err
	}
	return newsResp, nil
}

func (n *News) SubscribeNews(ctx context.Context) (<-chan *models.News, error) {
	id := rand.Int()
	n.subscribeNews[id] = make(chan *models.News)
	return n.subscribeNews[id], nil
}
