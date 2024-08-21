package product

import (
	"database/sql"

	"github.com/mukul-pixel/ims-bbs/cmd/types"

)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

// add product
func (s *Store) CreateProduct(product types.Product) error {
	_, err := s.db.Exec("INSERT INTO products(name,upc,category,image,quantity,location)VALUES(?,?,?,?,?,?)", product.Name, product.Upc, product.Category, product.Image, product.Quantity, product.Location)
	if err != nil {
		return err
	}
	return nil
}

//get product by upc
func(s *Store) GetProductByUpcBool(Upc string)bool{
	row := s.db.QueryRow("SELECT COUNT(*) FROM products WHERE upc = ?", Upc)
    var count int
    err := row.Scan(&count)
    if err != nil {
        return false
    }
    return count > 0
}

func (s *Store) GetProducts() ([]types.Product, error) {
	rows, err := s.db.Query("SELECT * FROM products")
	if err != nil {
		return nil, err
	}

	products := make([]types.Product, 0)
	for rows.Next() {
		p, err := scanRowsIntoProducts(rows)
		if err != nil {
			return nil, err
		}

		products = append(products, *p)
	}

	return products, nil
}

func scanRowsIntoProducts(rows *sql.Rows) (*types.Product, error) {
	products := new(types.Product)

	err := rows.Scan(
		&products.ID,
		&products.Name,
		&products.Upc,
		&products.Category,
		&products.Image,
		&products.Quantity,
		&products.Location,
		&products.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return products, nil
}
