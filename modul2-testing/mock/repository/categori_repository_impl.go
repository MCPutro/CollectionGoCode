package repository

import (
	"github.com/MCPutro/CollectionGoCode/modul2-testing/mock/entity"
)

type CategoryRepositoryImpl struct {
}

func NewCategoryRepositoryImpl() *CategoryRepositoryImpl {
	return &CategoryRepositoryImpl{}
}

func (c *CategoryRepositoryImpl) FindById(id string) *entity.Category {
	return &entity.Category{
		Id:   "1",
		Name: "haha",
	}
}
