package postgres

import (
	"context"
	"github.com/Aerok925/LeoNews/internal/models"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestGetNews(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err.Error())
	}
	defer mockDB.Close()

	type mockBehavior func()
	timeStart := time.Now()
	db := sqlx.NewDb(mockDB, "sqlmock")
	r := repos{db}
	testTable := []struct {
		name         string
		wantErr      bool
		mockBehavior mockBehavior
		want         []models.News
	}{
		{
			name: "OK",
			mockBehavior: func() {
				rows := sqlmock.NewRows([]string{"id", "title", "description", "img", "create_at"}).AddRow(1, "test", "test", nil, timeStart)
				mock.ExpectQuery("select news.id,     news.title,    news.description,    news.img,    news.create_at from news.news").WillReturnRows(rows)
			},
			want: []models.News{
				{
					ID:          1,
					Title:       "test",
					Description: "test",
					Img:         nil,
					CreatedAt:   timeStart,
				},
			},
			wantErr: false,
		},
		{
			name:    "NO DATA",
			wantErr: false,
			mockBehavior: func() {
				rows := sqlmock.NewRows([]string{"id", "title", "description", "img", "create_at"})
				mock.ExpectQuery("select news.id,     news.title,    news.description,    news.img,    news.create_at from news.news").WillReturnRows(rows)
			},
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mockBehavior()

			got, err := r.GetNews(context.Background())
			if !testCase.wantErr {
				assert.NoError(t, err)
				assert.Equal(t, got, testCase.want)
			} else {
				assert.Error(t, err)
			}

		})
	}

}
