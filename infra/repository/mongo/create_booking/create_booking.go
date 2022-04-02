package create_booking

import (
	"context"
	"log"
	"time"

	"github.com/pedrokunz/go_backend/entity/restaurant"
	"github.com/pedrokunz/go_backend/infra/repository/mongo/internal"
	"go.mongodb.org/mongo-driver/bson"
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

func (d *db) GetBookingsByDay(ctx context.Context, bookingDate time.Time) ([]*restaurant.Booking, error) {
	today, err := time.Parse("2006-01-02", bookingDate.Format("2006-01-02"))
	if err != nil {
		return nil, err
	}
	
	result, err := d.client.
		Database(d.database).
		Collection(d.collection).
		Find(ctx, bson.M{
			"date": bson.M{
				"$gte": primitive.NewDateTimeFromTime(today),
			},
		})
	if err != nil {
		return nil, err
	}

	var bookings []*restaurant.Booking
	for result.Next(ctx) {
		var b booking
		err := result.Decode(&b)
		if err != nil {
			return nil, err
		}

		bookings = append(bookings, &restaurant.Booking{
			Username:     b.Username,
			CustomerName: b.CustomerName,
			Date:         b.Date,
			TableID:      b.TableID,
		})
	}

	return bookings, nil
}
