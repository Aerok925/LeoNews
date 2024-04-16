package postgres

import (
	"context"
	"github.com/Aerok925/LeoNews/internal/models"
)

func (r *repos) GetNews(ctx context.Context) ([]models.News, error) {
	sqlQuery := `select 
    news.id, 
    news.title,
    news.description,
    news.img,
    news.create_at
	from news.news`
	var newsDst = make([]news, 0)
	err := r.db.SelectContext(ctx, &newsDst, sqlQuery)
	if err != nil {
		return nil, err
	}
	var newsResp []models.News

	for _, n := range newsDst {
		newsResp = append(newsResp, models.News{
			ID:          n.ID,
			Title:       n.Title,
			Description: n.Description,
			Img:         n.Img,
			CreatedAt:   n.CreatedAt,
		})
	}
	return newsResp, nil
}

func (r *repos) CreateNews(ctx context.Context, news models.News) (*models.News, error) {
	sqlQuery := `insert into news.news (title, description, img) VALUES (:title, :description, :img) returning id, title, description, img, create_at`

	tx := r.db.MustBegin()

	rows, err := tx.NamedQuery(sqlQuery, news)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	respModel := &models.News{}
	for rows.Next() {
		rows.Scan(
			&respModel.ID,
			&respModel.Title,
			&respModel.Description,
			&respModel.Img,
			&respModel.CreatedAt,
		)
	}
	defer rows.Close()

	return respModel, tx.Commit()
}
