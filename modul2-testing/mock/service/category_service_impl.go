package service

import (
	"errors"
	"github.com/MCPutro/CollectionGoCode/modul2-testing/mock/entity"
	"github.com/MCPutro/CollectionGoCode/modul2-testing/mock/repository"
)

type CategoryServiceImpl struct {
	CategoryRepository repository.CategoryRepository
}

func NewCategoryServiceImpl(categoryRepository repository.CategoryRepository) *CategoryServiceImpl {
	return &CategoryServiceImpl{CategoryRepository: categoryRepository}
}

func (c *CategoryServiceImpl) Get(id string) (*entity.Category, error) {
	category := c.CategoryRepository.FindById(id)

	if category == nil {
		return nil, errors.New("data not found")
	} else {
		return category, nil
	}
}
