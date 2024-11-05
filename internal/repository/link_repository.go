package repository

import (
	"github.com/elonghi/encurtago/internal/domain"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type linkRepository struct {
	db    *gorm.DB
	redis *redis.Client
}

type LinkRepository interface {
	FindAll() ([]domain.Link, error)
	FindById(id int) (domain.Link, error)
	FindByShortURL(shortURL string) (domain.Link, error)
	FindByUrl(url string) (domain.Link, error)
	Create(link domain.Link) (domain.Link, error)
	SetCache(key string, value string) error
	GetCache(key string) (string, error)
	Update(link domain.Link) (domain.Link, error)
}

func NewLinkRepository(db *gorm.DB, redis *redis.Client) LinkRepository {
	return &linkRepository{db: db, redis: redis}
}

func (r *linkRepository) Update(link domain.Link) (domain.Link, error) {
	result := r.db.Save(&link)
	return link, result.Error
}

func (r *linkRepository) Create(link domain.Link) (domain.Link, error) {
	result := r.db.Create(&link)
	return link, result.Error
}

func (r *linkRepository) FindAll() ([]domain.Link, error) {
	var links []domain.Link
	result := r.db.Find(&links)
	return links, result.Error
}

func (r *linkRepository) FindById(id int) (domain.Link, error) {
	var link domain.Link
	result := r.db.First(&link, id)
	return link, result.Error
}

func (r *linkRepository) FindByShortURL(shortURL string) (domain.Link, error) {
	var link domain.Link
	result := r.db.Where("short_url = ?", shortURL).First(&link)
	return link, result.Error
}

func (r *linkRepository) FindByUrl(url string) (domain.Link, error) {
	var link domain.Link
	result := r.db.Where("url = ?", url).First(&link)
	return link, result.Error
}

func (r *linkRepository) SetCache(key string, value string) error {
	return r.redis.Set(r.redis.Context(), key, value, 0).Err()
}

func (r *linkRepository) GetCache(key string) (string, error) {
	return r.redis.Get(r.redis.Context(), key).Result()
}
