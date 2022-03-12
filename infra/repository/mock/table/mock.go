package table

import "github.com/pedrokunz/go_backend/entity/restaurant"

type Mock struct {
	tables []restaurant.Table
}

func New() *Mock {
	tables := make([]restaurant.Table, 0)
	tables = append(tables, restaurant.Table{ID: 1})
	tables = append(tables, restaurant.Table{ID: 2})
	tables = append(tables, restaurant.Table{ID: 3})
	tables = append(tables, restaurant.Table{ID: 4})
	tables = append(tables, restaurant.Table{ID: 5})
	tables = append(tables, restaurant.Table{ID: 6})
	tables = append(tables, restaurant.Table{ID: 7})
	tables = append(tables, restaurant.Table{ID: 8})
	tables = append(tables, restaurant.Table{ID: 9})
	tables = append(tables, restaurant.Table{ID: 10})

	return &Mock{
		tables: tables,
	}
}

func (m *Mock) GetAvailableTables() ([]restaurant.Table, error) {
	return m.tables, nil
}
