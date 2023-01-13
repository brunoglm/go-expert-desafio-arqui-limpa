package database

import (
	"database/sql"
	"testing"

	"project/internal/entity"

	"github.com/stretchr/testify/suite"

	// sqlite3
	_ "github.com/mattn/go-sqlite3"
)

type OrderRepositoryTestSuite struct {
	suite.Suite
	Db *sql.DB
}

// func (suite *OrderRepositoryTestSuite) SetupSuite() {
// 	db, err := sql.Open("sqlite3", ":memory:")
// 	suite.NoError(err)
// 	db.Exec("CREATE TABLE orders (id varchar(255) NOT NULL, price float NOT NULL, tax float NOT NULL, final_price float NOT NULL, PRIMARY KEY (id))")
// 	suite.Db = db
// }

// func (suite *OrderRepositoryTestSuite) TearDownAllSuite() {
// 	suite.Db.Close()
// }

func (suite *OrderRepositoryTestSuite) SetupTest() {
	db, err := sql.Open("sqlite3", ":memory:")
	suite.NoError(err)
	db.Exec("CREATE TABLE orders (id varchar(255) NOT NULL, price float NOT NULL, tax float NOT NULL, final_price float NOT NULL, PRIMARY KEY (id))")
	suite.Db = db
}

func (suite *OrderRepositoryTestSuite) TearDownTest() {
	suite.Db.Close()
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(OrderRepositoryTestSuite))
}

func (suite *OrderRepositoryTestSuite) TestGivenAnOrder_WhenSave_ThenShouldSaveOrder() {
	order, err := entity.NewOrder("123", 10.0, 2.0)
	suite.NoError(err)
	suite.NoError(order.CalculateFinalPrice())
	repo := NewOrderRepository(suite.Db)
	err = repo.Save(order)
	suite.NoError(err)

	var orderResult entity.Order
	err = suite.Db.QueryRow("Select id, price, tax, final_price from orders where id = ?", order.ID).
		Scan(&orderResult.ID, &orderResult.Price, &orderResult.Tax, &orderResult.FinalPrice)

	suite.NoError(err)
	suite.Equal(order.ID, orderResult.ID)
	suite.Equal(order.Price, orderResult.Price)
	suite.Equal(order.Tax, orderResult.Tax)
	suite.Equal(order.FinalPrice, orderResult.FinalPrice)
}

func (suite *OrderRepositoryTestSuite) TestgivenThatTheListMethodWasCalled_ThenShouldListAllOrders() {
	order1, _ := entity.NewOrder("1", 10.0, 2.0)
	order2, _ := entity.NewOrder("2", 20.0, 3.0)
	suite.NoError(order1.CalculateFinalPrice())
	suite.NoError(order2.CalculateFinalPrice())

	repo := NewOrderRepository(suite.Db)
	repo.Save(order1)
	repo.Save(order2)

	orders, err := repo.List()
	suite.NoError(err)

	suite.Equal(len(orders), 2)
	suite.Equal(orders[0], order1)
	suite.Equal(orders[1], order2)
}
