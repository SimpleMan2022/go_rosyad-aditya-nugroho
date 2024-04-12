package repositories

import (
	"go-wishlist-api/entities"
	"gorm.io/gorm"
)

type WishlistRepository interface {
	GetAll() ([]*entities.Wishlist, error)
	Create(wishlist *entities.Wishlist) (*entities.Wishlist, error)
}

type wishlistRepository struct {
	db *gorm.DB
}

func NewWishlistRepository(db *gorm.DB) *wishlistRepository {
	return &wishlistRepository{db}
}

func (r *wishlistRepository) GetAll() ([]*entities.Wishlist, error) {
	var wishlists []*entities.Wishlist
	if err := r.db.Find(&wishlists).Error; err != nil {
		return nil, err
	}
	return wishlists, nil
}
func (r *wishlistRepository) Create(wishlist *entities.Wishlist) (*entities.Wishlist, error) {
	if err := r.db.Create(&wishlist).Error; err != nil {
		return nil, err
	}
	return wishlist, nil
}
