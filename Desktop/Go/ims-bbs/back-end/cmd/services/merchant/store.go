package merchant

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

func (s *Store) CreateMerchant(merchant types.Merchant) error {
	_, err := s.db.Exec("INSERT INTO merchant(merchant_name,merchant_address,in_contact,contact_info,category) VALUES (?,?,?,?,?)", merchant.Merchant_Name, merchant.Merchant_Address, merchant.In_Contact, merchant.Contact_Info, merchant.Category)
	if err != nil {
		return err
	}

	return nil
}

func (s *Store) GetMerchants() ([]types.Merchant, error) {
	rows, err := s.db.Query("SELECT * FROM merchant")
	if err != nil {
		return nil, err
	}

	merchants := make([]types.Merchant, 0)
	for rows.Next() {
		m, err := scanRowsIntoMerchants(rows)
		if err != nil {
			return nil, err
		}

		merchants = append(merchants, *m)
	}
	return merchants,nil
}

func scanRowsIntoMerchants(rows *sql.Rows) (*types.Merchant, error) {
	merchants := new(types.Merchant)

	err:=rows.Scan(
		&merchants.ID,
		&merchants.Merchant_Name,
		&merchants.Merchant_Address,
		&merchants.In_Contact,
		&merchants.Contact_Info,
		&merchants.Category,
		&merchants.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return merchants, nil
}
