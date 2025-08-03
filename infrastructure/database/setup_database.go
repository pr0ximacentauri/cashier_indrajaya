package database

import(
	"database/sql"
	_ "modernc.org/sqlite"
)

func SetupDatabase() (*sql.DB, error) {
	db, err := sql.Open("sqlite", "./indrajaya.db")
	if err != nil {
		return nil, err
	}

	createTables := `
	CREATE TABLE IF NOT EXISTS products (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		price REAL NOT NULL,
		stock INTEGER NOT NULL,
		category TEXT NOT NULL
	);

	CREATE TABLE IF NOT EXISTS sales (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		product_id INTEGER NOT NULL,
		quantity INTEGER NOT NULL,
		price REAL NOT NULL,
		total REAL NOT NULL,
		sale_date DATETIME NOT NULL,
		FOREIGN KEY (product_id) REFERENCES products (id)
	);
	`

	if _, err := db.Exec(createTables); err != nil {
		return nil, err
	}

	return db, nil
}