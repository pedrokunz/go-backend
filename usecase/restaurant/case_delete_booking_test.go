package restaurant_test

import (
	"context"
	"testing"
	"time"

	restaurantMockRepository "github.com/pedrokunz/go_backend/infra/repository/mock/booking"
	"github.com/pedrokunz/go_backend/usecase/restaurant"
)

func TestDeleteBooking(t *testing.T) {
	bookingMock := restaurantMockRepository.New()
	ctx := context.Background()
	now := time.Now()
	createValidBookings(t, ctx, bookingMock, now)

	type args struct {
		id string
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "SUCCESS",
			args: args{
				id: bookingMock.Bookings[0].ID,
			},
			wantErr: false,
		},
		{
			name: "FAILURE",
			args: args{
				id: "invalid_id",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := restaurant.NewDeleteBooking(bookingMock)
			if err := u.Perform(ctx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("usecaseDeleteBooking.Perform() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
