package mock_repo

import (
	"github.com/MCPutro/CollectionGoCode/modul2-testing/mock/entity"
	"github.com/stretchr/testify/mock"
)

type CategoryRepositoryImplMock struct {
	Mock mock.Mock
}

func NewCategoryRepositoryImplMock(mock mock.Mock) *CategoryRepositoryImplMock {
	return &CategoryRepositoryImplMock{Mock: mock}
}

func (c *CategoryRepositoryImplMock) FindById(id string) *entity.Category {
	arguments := c.Mock.Called(id)

	if arguments.Get(0) == nil {
		return nil
	}

	category := arguments.Get(0).(entity.Category)
	return &category
}
