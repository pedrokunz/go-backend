package create_booking

import (
	"context"
	"log"
	"time"

	"github.com/pedrokunz/go_backend/entity/restaurant"
	"github.com/pedrokunz/go_backend/infra/repository/mongo/internal"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type db struct {
	client     *mongo.Client
	database   string
	collection string
}

func New() (*db, error) {
	client, err := internal.Connect()
	if err != nil {
		return nil, err
	}

	return &db{
		client:     client,
		database:   "restaurant",
		collection: "bookings",
	}, nil
}

type booking struct {
	ID           string    `bson:"_id"`
	Username     string    `bson:"username"`
	CustomerName string    `bson:"customer_name"`
	Date         time.Time `bson:"date"`
	TableID      int       `bson:"table_id"`
}

func toBookingMongo(b *restaurant.Booking) (*primitive.D, error) {
	return internal.ToDoc(booking{
		Username:     b.Username,
		CustomerName: b.CustomerName,
		Date:         b.Date,
		TableID:      b.TableID,
	})
}

func (d *db) Create(ctx context.Context, booking *restaurant.Booking) error {
	doc, err := toBookingMongo(booking)
	if err != nil {
		return err
	}

	result, err := d.client.
		Database(d.database).
		Collection(d.collection).
		InsertOne(ctx, doc)
	if err != nil {
		return err
	}

	log.Printf("Inserted a single document: %v", result.InsertedID)

	return nil
}
