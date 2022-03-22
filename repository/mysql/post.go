package mysql

import (
	"encoding/json"
	"github.com/VerTox/rentateam_test_blog/domain/model"
	"github.com/jinzhu/gorm"
	"time"
)

type Post struct {
	Id        int        `gorm:"primary_key"`
	Title     string     `gorm:"column:title"`
	Text      string     `gorm:"column:text;type:text"`
	Tags      string     `gorm:"column:tags"`
	CreatedAt *time.Time `gorm:"column:created_at"`
}

func (Post) TableName() string {
	return "posts"
}

func (p *Post) Model() (*model.Post, error) {
	var tags []string

	err := json.Unmarshal([]byte(p.Tags), &tags)

	if err != nil {
		return nil, err
	}

	return &model.Post{
		Id:        p.Id,
		Title:     p.Title,
		Text:      p.Text,
		Tags:      tags,
		CreatedAt: p.CreatedAt,
	}, nil
}

type PostCast model.Post

func (pc PostCast) Repository() (*Post, error) {
	tags, err := json.Marshal(pc.Tags)

	if err != nil {
		return nil, err
	}

	return &Post{
		Id:        pc.Id,
		Title:     pc.Title,
		Text:      pc.Text,
		Tags:      string(tags),
		CreatedAt: pc.CreatedAt,
	}, nil
}

type PostRepository struct {
	Connection *gorm.DB
}

func NewPostRepository(c *gorm.DB) *PostRepository {
	return &PostRepository{
		Connection: c,
	}
}

func (r *PostRepository) GetList() ([]*model.Post, error) {
	var ps []*Post

	err := r.Connection.Find(&ps).Error

	if err != nil {
		return nil, err
	}

	pm := make([]*model.Post, len(ps))

	for i, p := range ps {
		pm[i], err = p.Model()

		if err != nil {
			return nil, err
		}
	}

	return pm, err
}

func (r *PostRepository) Store(p *model.Post) (*model.Post, error) {
	pc, err := PostCast(*p).Repository()

	if err != nil {
		return nil, err
	}

	now := time.Now()

	pc.CreatedAt = &now

	err = r.Connection.Create(pc).Error

	if err != nil {
		return nil, err
	}

	pm, err := pc.Model()

	if err != nil {
		return nil, err
	}

	return pm, nil
}
