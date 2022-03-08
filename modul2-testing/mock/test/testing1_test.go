package unit_test

import (
	"github.com/MCPutro/CollectionGoCode/modul2-testing/mock/entity"
	mockRepo "github.com/MCPutro/CollectionGoCode/modul2-testing/mock/mock-repo"
	"github.com/MCPutro/CollectionGoCode/modul2-testing/mock/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

var repositoryImplMock = mockRepo.NewCategoryRepositoryImplMock(mock.Mock{})
var serviceImpl = service.NewCategoryServiceImpl(repositoryImplMock)

func Test_Mock1(t *testing.T) {
	//ini bagian mock nya, saat function FindById dg arg 1 dipanggil akan mengembalikan nil
	repositoryImplMock.Mock.On("FindById", "1").Return(nil)
	category, err := serviceImpl.Get("1")

	assert.NotNil(t, err)
	assert.Nil(t, category)
}

func Test_Mock2(t *testing.T) {
	c := entity.Category{
		Id:   "2",
		Name: "2",
	}

	//ini bagian mock nya, saat function FindById dg arg 2 dipanggil akan mengembalikan variable c
	repositoryImplMock.Mock.On("FindById", "2").Return(c)
	ca, err := serviceImpl.Get("2")
	assert.Nil(t, err)
	assert.Equal(t, c.Id, ca.Id)
}
