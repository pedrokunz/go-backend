package booking

import (
	"context"
	"log"
	"time"

	"github.com/pedrokunz/go_backend/entity/restaurant"
	"github.com/pedrokunz/go_backend/helper"
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
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	Username     string             `bson:"username"`
	CustomerName string             `bson:"customer_name"`
	Date         time.Time          `bson:"date"`
	TableID      int                `bson:"table_id"`
	Status       string             `bson:"status"`
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

func (d *db) GetBookingsForDay(ctx context.Context, bookingDate time.Time) ([]*restaurant.Booking, error) {
	firstMomentOfDay := helper.RemoveTimeFromDate(bookingDate)
	lastMomentOfDay := helper.SetDateToLastMomentOfTheDay(bookingDate)

	result, err := d.client.
		Database(d.database).
		Collection(d.collection).
		Find(ctx, bson.M{
			"date": bson.M{
				"$gte": primitive.NewDateTimeFromTime(firstMomentOfDay),
				"$lte": primitive.NewDateTimeFromTime(lastMomentOfDay),
			},
			"status": bson.M{
				"$ne": "deleted",
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
			ID:           b.ID.Hex(),
			Username:     b.Username,
			CustomerName: b.CustomerName,
			Date:         b.Date,
			TableID:      b.TableID,
			Status:       b.Status,
		})
	}

	return bookings, nil
}

func (d *db) GetBookingsFromDay(ctx context.Context, bookingDate time.Time) ([]*restaurant.Booking, error) {
	firstMomentOfDay := helper.RemoveTimeFromDate(bookingDate)

	result, err := d.client.
		Database(d.database).
		Collection(d.collection).
		Find(ctx, bson.M{
			"date": bson.M{
				"$gte": primitive.NewDateTimeFromTime(firstMomentOfDay),
			},
			"status": bson.M{
				"$ne": "deleted",
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
			ID:           b.ID.Hex(),
			Username:     b.Username,
			CustomerName: b.CustomerName,
			Date:         b.Date,
			TableID:      b.TableID,
			Status:       b.Status,
		})
	}

	return bookings, nil
}

func (d *db) Delete(ctx context.Context, id string) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = d.client.
		Database(d.database).
		Collection(d.collection).
		UpdateByID(ctx, oid, bson.M{
			"$set": bson.M{
				"status": "deleted",
			},
		})

	if err != nil {
		return err
	}

	log.Printf("Deleted a single document: %s", oid)

	return nil
}
