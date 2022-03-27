package create_booking

import (
	"context"
	"testing"
	"time"

	"github.com/pedrokunz/go_backend/entity/restaurant"
	"go.mongodb.org/mongo-driver/bson"
)

func Test_db_Create(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode.")
	}

	type args struct {
		booking *restaurant.Booking
	}

	t.Setenv("MONGO_URI", "mongodb://mongo:27017")
	mongo, err := New()
	if err != nil {
		t.FailNow()
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				booking: &restaurant.Booking{
					Username:     "user test",
					CustomerName: "customer test",
					Date:         time.Now(),
					TableID:      1,
				},
			},
		},
	}

	ctx := context.Background()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := mongo.Create(ctx, tt.args.booking); (err != nil) != tt.wantErr {
				t.Errorf("db.Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}

	// delete all bookings
	mongo.client.Database(mongo.database).Collection(mongo.collection).DeleteMany(ctx, bson.D{})
}
