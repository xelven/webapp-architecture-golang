package product_test

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"

	"webapp-core/core/domain/product"
	mockDB "webapp-core/mock/db"
	testUtil "webapp-core/util/test"
)

func TestRepository_GetList(test *testing.T) {
	test.Parallel()

	db, mock, err := mockDB.NewMockDB()
	testUtil.NoError(test, err)

	query := product.NewQuerier(db)

	mockRows := sqlmock.NewRows([]string{"id", "name", "creator"}).
		AddRow(uuid.New(), "Product1", "Creator1").
		AddRow(uuid.New(), "Product2", "Creator2")

	mock.ExpectQuery("^SELECT (.+) FROM \"products\"").WillReturnRows(mockRows)

	products, err := query.GetList()
	testUtil.NoError(test, err)
	testUtil.Equal(test, len(products), 2)
}
