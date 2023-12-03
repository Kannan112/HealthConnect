package interfaces

import (
	"context"

	"github.com/easy-health/pkg/domain"
	"github.com/easy-health/pkg/utils/req"
	"github.com/easy-health/pkg/utils/res"
)

type CategoryRepository interface {
	// admin side
	CreateCategory(ctx context.Context, category req.Category) error
	UpdateCategoryName(ctxc context.Context, name string, id uint) error
	UpdateCategoryDescription(ctxc context.Context, description string, id uint) error

	DeleteCategory(ctx context.Context, category_id int) error
	ListCategory(ctx context.Context) ([]domain.Categories, error)
	FeatchCategoryDetails(ctx context.Context, name string) (domain.Categories, error)

	// patienst side
	ListCategoryUser(ctx context.Context) ([]res.CategoriesUser, error)

	//doctor side
	CategoryIdCheck(ctx context.Context, categoryId uint) (bool, error)
}
