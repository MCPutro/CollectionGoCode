package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/MCPutro/CollectionGoCode/modul8-database/entity"
	"strconv"
)

type commentRepositoryImpl struct {
	DB *sql.DB
}

func GetCommentRepository(db *sql.DB) CommentRepositoryInterface {
	return &commentRepositoryImpl{DB: db}
}

func (repository *commentRepositoryImpl) Insert(ctx context.Context, comment entity.Comment) (entity.Comment, error) {
	defer func(DB *sql.DB) {
		err := DB.Close()
		if err != nil {
		}
		fmt.Println("close koneksi")
	}(repository.DB)

	script := "insert into comments ( email, comment) values ( ?, ?);"

	result, err := repository.DB.ExecContext(ctx, script, comment.Email, comment.Comment)

	if err != nil {
		fmt.Println("error1 :", err)
	}

	insertId, err2 := result.LastInsertId()
	if err2 != nil {
		fmt.Println("error2 :", err2)
	}
	comment.Id = int32(insertId)

	return comment, nil
}

func (repository *commentRepositoryImpl) FindById(ctx context.Context, id int32) (entity.Comment, error) {
	script := "SELECT id, email, comment FROM comments WHERE id = ? LIMIT 1"
	rows, err := repository.DB.QueryContext(ctx, script, id)
	comment := entity.Comment{}
	if err != nil {
		return comment, err
	}
	defer rows.Close()
	if rows.Next() {
		// ada
		rows.Scan(&comment.Id, &comment.Email, &comment.Comment)
		return comment, nil
	} else {
		// tidak ada
		return comment, errors.New("Id " + strconv.Itoa(int(id)) + " Not Found")
	}
}

func (repository *commentRepositoryImpl) FindAll(ctx context.Context) ([]entity.Comment, error) {
	script := "SELECT id, email, comment FROM comments"

	rows, err := repository.DB.QueryContext(ctx, script)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var comments []entity.Comment
	for rows.Next() {
		comment := entity.Comment{}
		err := rows.Scan(&comment.Id, &comment.Email, &comment.Comment)
		if err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}
	return comments, nil
}
