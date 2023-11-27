package repository

import (
	"context"

	"github.com/easy-health/pkg/domain"
	interfaces "github.com/easy-health/pkg/repository/interface"
	"github.com/easy-health/pkg/utils/req"
	"github.com/easy-health/pkg/utils/res"
	"gorm.io/gorm"
)

type CategoriesDatabase struct {
	DB *gorm.DB
}

func NewCategoriesRepository(DB *gorm.DB) interfaces.CategoryRepository {
	return &CategoriesDatabase{DB: DB}
}

// admin side
func (c *CategoriesDatabase) CreateCategory(ctx context.Context, category req.Category) error {
	query := `INSERT INTO categories (name, created_at, description) VALUES ($1,NOW(), $2)`
	err := c.DB.Exec(query, category.Name, category.Description).Error
	if err != nil {
		return err
	}
	return nil
}
func (c *CategoriesDatabase) DeleteCategory(ctx context.Context, category_id int) error {
	query := `DELETE FROM categories WHERE id = $1`
	if err := c.DB.Exec(query, category_id).Error; err != nil {
		return err
	}
	return nil
}
func (c *CategoriesDatabase) ListCategory(ctx context.Context) ([]domain.Categories, error) {
	var data []domain.Categories
	query := `select * from categories`
	err := c.DB.Raw(query).Scan(&data).Error
	if err != nil {
		return data, err
	}
	return data, nil
}

// patienst side
func (c *CategoriesDatabase) ListCategoryUser(ctx context.Context) ([]res.CategoriesUser, error) {
	var Categories []res.CategoriesUser
	query := `select * from categories`
	if err := c.DB.Raw(query).Scan(&Categories).Error; err != nil {
		return nil, err
	}
	return Categories, nil
}

// doctor side
func (c *CategoriesDatabase) CategoryIdCheck(ctx context.Context, categoryId uint) (bool, error) {
	var check bool
	query := `select Exists(select * from categories where id=$1)`
	if err := c.DB.Raw(query, categoryId).Scan(&check).Error; err != nil {
		return false, err
	}
	return check, nil
}

func (c *CategoriesDatabase) FeatchCategoryDetails(ctx context.Context, name string) (domain.Categories, error) {
	var cat domain.Categories
	query := "select * from categories where name=$1"
	if err := c.DB.Raw(query, name).Scan(&cat).Error; err != nil {
		return cat, err
	}
	return cat, nil
}

func (c *CategoriesDatabase) UpdateCategoryName(ctxc context.Context, name string, id uint) error {
	query := `update categories SET name=$1 Where id=$2`
	if err := c.DB.Exec(query, name).Error; err != nil {
		return nil
	}
	return nil
}

func (c *CategoriesDatabase) UpdateCategoryDescription(ctxc context.Context, description string, id uint) error {
	query := `update categories SET Description=$1 Where id=$2`
	if err := c.DB.Exec(query, description).Error; err != nil {
		return err
	}
	return nil
}
