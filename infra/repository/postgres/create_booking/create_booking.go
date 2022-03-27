package create_booking

import (
	"context"
	"time"

	"github.com/pedrokunz/go_backend/entity/restaurant"
	"github.com/pedrokunz/go_backend/infra/repository/postgres/internal"
	"gorm.io/gorm"
)

type db struct {
	client   *gorm.DB
	database string
}

func New() (*db, error) {
	client, err := internal.Connect()
	if err != nil {
		return nil, err
	}

	client.AutoMigrate(&booking{})

	return &db{
		client: client,
	}, nil
}

type booking struct {
	gorm.Model
	Username     string    `sql:"username"`
	CustomerName string    `sql:"customer_name"`
	Date         time.Time `sql:"date"`
	TableID      int       `sql:"table_id"`
}

func toBookingPostgres(b *restaurant.Booking) *booking {
	return &booking{
		Username:     b.Username,
		CustomerName: b.CustomerName,
		Date:         b.Date,
		TableID:      b.TableID,
	}
}

func (d *db) Create(ctx context.Context, booking *restaurant.Booking) error {
	result := d.client.Create(toBookingPostgres(booking))
	if result.Error != nil {
		return result.Error
	}

	return nil
}
