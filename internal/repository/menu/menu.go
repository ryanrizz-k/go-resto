package menu

import (
	"context"

	"github.com/rynr00/go-resto/internal/model"
	"github.com/rynr00/go-resto/internal/tracing"
	"gorm.io/gorm"
)

type menuRepo struct {
	db *gorm.DB
}

func GetRepository(db *gorm.DB) Repository {
	return &menuRepo{
		db: db,
	}
}

func (m *menuRepo) GetMenuList(ctx context.Context, menuType string) ([]model.MenuItem, error) {
	ctx, span := tracing.CreateSpan(ctx, "GetMenuList")
	defer span.End()

	var menuData []model.MenuItem

	if err := m.db.WithContext(ctx).Where(model.MenuItem{Type: model.MenuType(menuType)}).Find(&menuData).Error; err != nil {
		return nil, err
	}

	return menuData, nil
}

func (m *menuRepo) GetMenu(ctx context.Context, orderCode string) (model.MenuItem, error) {
	ctx, span := tracing.CreateSpan(ctx, "GetMenu")
	defer span.End()

	var menuData model.MenuItem

	if err := m.db.WithContext(ctx).Where((model.MenuItem{OrderCode: orderCode})).First(&menuData).Error; err != nil {
		return menuData, err
	}

	return menuData, nil
}
