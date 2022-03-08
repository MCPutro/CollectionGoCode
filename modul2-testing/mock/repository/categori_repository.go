package repository

import (
	"github.com/MCPutro/CollectionGoCode/modul2-testing/mock/entity"
)

type CategoryRepository interface {
	FindById(id string) *entity.Category
}
