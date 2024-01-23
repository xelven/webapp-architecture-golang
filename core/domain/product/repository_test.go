package product_test

import (
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"

	"webapp-core/core/domain/product"
	mockDB "webapp-core/mock/db"
	testUtil "webapp-core/util/test"
)

func TestRepository_Create(test *testing.T) {
	test.Parallel()

	db, mock, err := mockDB.NewMockDB()
	testUtil.NoError(test, err)

	repo := product.NewRepository(db)

	id := uuid.New()
	mock.ExpectBegin()
	mock.ExpectExec("^INSERT INTO \"products\" ").
		WithArgs(id, "test name", "111", mockDB.AnyTime{}, "", "", mockDB.AnyTime{}, mockDB.AnyTime{}, nil).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	newProduct := &product.Product{ID: id, Name: "test name", Creator: "111", PublishedDate: time.Now()}
	_, err = repo.Create(newProduct)
	testUtil.NoError(test, err)
}
