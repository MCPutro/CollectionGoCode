package test

import (
	"context"
	"fmt"
	db_driver "github.com/MCPutro/CollectionGoCode/modul8-database/db-driver"
	"github.com/MCPutro/CollectionGoCode/modul8-database/entity"
	"github.com/MCPutro/CollectionGoCode/modul8-database/repository"
	"testing"
)

func TestInsert(t *testing.T) {
	commentRepository := repository.GetCommentRepository(db_driver.GetConnection())

	ctx := context.Background()

	commentBaru := entity.Comment{
		Email:   "coba@gmail.com",
		Comment: "123456789",
	}

	insert, err := commentRepository.Insert(ctx, commentBaru)
	if err != nil {
		panic(err)
	}

	fmt.Println(insert)
}
