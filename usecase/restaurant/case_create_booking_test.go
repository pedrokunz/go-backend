package restaurant_test

import (
	"context"
	"testing"
	"time"

	bookingMockRepository "github.com/pedrokunz/go_backend/infra/repository/mock/booking"
	"github.com/pedrokunz/go_backend/usecase/restaurant"
	"github.com/stretchr/testify/assert"
)

func TestCreateBooking(t *testing.T) {
	type args struct {
		input restaurant.CreateBookingInput
	}

	saturdayNoon := time.Date(2032, time.March, 21, 12, 0, 0, 0, time.UTC)

	tests := []struct {
		name        string
		args        args
		expectedErr string
	}{
		{
			name: "SUCCESS",
			args: args{
				input: restaurant.CreateBookingInput{
					Username:     "user_test",
					BookingDate:  saturdayNoon.Format(time.RFC3339),
					CustomerName: "customer_test",
					TableID:      1,
				},
			},
		},
		{
			name: "SUCCESS - Case booking is more than 2 hours from existing booking",
			args: args{
				input: restaurant.CreateBookingInput{
					Username:     "user_test",
					BookingDate:  saturdayNoon.Add(3 * time.Hour).Format(time.RFC3339),
					CustomerName: "customer_test",
					TableID:      1,
				},
			},
		},
		{
			name: "ERROR - Case booking is less or equal than 2 hours from existing booking",
			args: args{
				input: restaurant.CreateBookingInput{
					Username:     "user_test",
					BookingDate:  saturdayNoon.Add(2 * time.Hour).Format(time.RFC3339),
					CustomerName: "customer_test",
					TableID:      1,
				},
			},
			expectedErr: "table not available for booking",
		},
	}

	bookingMock := bookingMockRepository.New()
	ctx := context.Background()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := restaurant.NewCreateBooking(bookingMock)
			err := u.Perform(ctx, tt.args.input)
			if err != nil {
				assert.EqualError(t, err, tt.expectedErr, "expected: %s, got: %s", tt.expectedErr, err.Error())
			}
		})
	}
}
