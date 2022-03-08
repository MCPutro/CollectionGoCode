package service

import (
	"github.com/MCPutro/CollectionGoCode/modul2-testing/mock/entity"
)

type CategoryService interface {
	Get(id string) (*entity.Category, error)
}
