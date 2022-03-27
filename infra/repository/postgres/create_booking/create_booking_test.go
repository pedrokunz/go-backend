package create_booking

import (
	"context"
	"testing"
	"time"

	"github.com/pedrokunz/go_backend/entity/restaurant"
)

func Test_db_Create(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode.")
	}

	type args struct {
		booking *restaurant.Booking
	}

	t.Setenv("POSTGRES_HOST", "postgres")
	t.Setenv("POSTGRES_USER", "root")
	t.Setenv("POSTGRES_PASSWORD", "root")
	t.Setenv("POSTGRES_SSLMODE", "disable")
	db, err := New()
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
			if err := db.Create(ctx, tt.args.booking); (err != nil) != tt.wantErr {
				t.Errorf("db.Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}

	db.client.Exec("DELETE FROM bookings")
}
