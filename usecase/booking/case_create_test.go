package booking_test

import (
	"github.com/pedrokunz/go_backend/usecase/booking"
	"testing"
	"time"
)

func TestCreateBooking(t *testing.T) {
	type args struct {
		input booking.CreateBookingInput
	}
	tests := []struct {
		name string
		args args
		//want    *booking.CreateBookingOutput
		wantErr bool
	}{
		{
			name: "SUCCESS",
			args: args{
				input: booking.CreateBookingInput{
					Username:     "user_test",
					BookingDate:  time.Now().String(),
					CustomerName: "customer_test",
					TableID:      1,
				},
			},
			wantErr: false,
		},
	}

	//todo mocks

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//todo mocks
			u := booking.NewCreate(nil, nil, nil)
			err := u.Perform(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Perform() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
