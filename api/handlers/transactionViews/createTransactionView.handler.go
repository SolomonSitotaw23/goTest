package transactionviews

import "test_project/config"

func CreateTransactionViews() error {
	db := config.SetupDatabase()
	defer db.Close()

	// Create the TransactionViews view using raw SQL and inner join
	if err := db.Exec("CREATE VIEW transaction_views AS SELECT t.id, t.customer_id, c.customer_name, t.item_id, i.item_name, t.qty, t.price, t.amount FROM transactions t INNER JOIN customers c ON t.customer_id = c.id INNER JOIN items i ON t.item_id = i.id").Error; err != nil {
		return err
	}
	return nil
}
