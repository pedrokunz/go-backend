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

	saturdayAfternoon := time.Date(2032, time.March, 21, 15, 0, 0, 0, time.UTC)

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
					BookingDate:  saturdayAfternoon.Format(time.RFC3339),
					CustomerName: "customer_test",
					TableID:      1,
				},
			},
		},
		{
			name: "SUCCESS - Case booking is 3 hours after existing booking",
			args: args{
				input: restaurant.CreateBookingInput{
					Username:     "user_test",
					BookingDate:  saturdayAfternoon.Add(3 * time.Hour).Format(time.RFC3339),
					CustomerName: "customer_test",
					TableID:      1,
				},
			},
		},
		{
			name: "SUCCESS - Case booking is 2 hours before existing booking",
			args: args{
				input: restaurant.CreateBookingInput{
					Username:     "user_test",
					BookingDate:  saturdayAfternoon.Add(-2 * time.Hour).Format(time.RFC3339),
					CustomerName: "customer_test",
					TableID:      1,
				},
			},
		},
		{
			name: "SUCCESS - Case booking is for other table",
			args: args{
				input: restaurant.CreateBookingInput{
					Username:     "user_test",
					BookingDate:  saturdayAfternoon.Format(time.RFC3339),
					CustomerName: "customer_test",
					TableID:      2,
				},
			},
		},
		{
			name: "ERROR - Case booking is out of working hours",
			args: args{
				input: restaurant.CreateBookingInput{
					Username:     "user_test",
					BookingDate:  saturdayAfternoon.Add(8 * time.Hour).Format(time.RFC3339),
					CustomerName: "customer_test",
					TableID:      1,
				},
			},
			expectedErr: "not working datetime",
		},
		{
			name: "ERROR - Case booking is out of booking hours",
			args: args{
				input: restaurant.CreateBookingInput{
					Username:     "user_test",
					BookingDate:  saturdayAfternoon.Add(7 * time.Hour).Format(time.RFC3339),
					CustomerName: "customer_test",
					TableID:      1,
				},
			},
			expectedErr: "not booking datetime",
		},
		{
			name: "ERROR - Case booking is less or equal than 2 hours from existing booking",
			args: args{
				input: restaurant.CreateBookingInput{
					Username:     "user_test",
					BookingDate:  saturdayAfternoon.Add(2 * time.Hour).Format(time.RFC3339),
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
