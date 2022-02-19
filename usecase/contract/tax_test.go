package contract_test

import (
	"fmt"
	"github.com/pedrokunz/go_backend/entity/tax"
	repositoryTaxMock "github.com/pedrokunz/go_backend/infra/repository/mock/tax"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTaxCalculator(t *testing.T) {
	type args struct {
		origin      string
		destination string
		plan        string
		duration    float64
	}
	tests := []struct {
		name            string
		args            args
		wantWithPlan    float64
		wantWithoutPlan float64
		wantErr         bool
		taxError        tax.Error
	}{
		{
			args:            args{origin: "011", destination: "016", plan: "FaleMais30", duration: 20},
			wantWithPlan:    0.0,
			wantWithoutPlan: 38.00,
			wantErr:         false,
		},
		{
			args:            args{origin: "011", destination: "017", plan: "FaleMais60", duration: 80},
			wantWithPlan:    37.40,
			wantWithoutPlan: 136.00,
			wantErr:         false,
		},
		{
			args:            args{origin: "018", destination: "011", plan: "FaleMais120", duration: 200},
			wantWithPlan:    167.20,
			wantWithoutPlan: 380.00,
			wantErr:         false,
		},
		{
			args:            args{origin: "018", destination: "017", plan: "FaleMais30", duration: 100},
			wantWithPlan:    0.0,
			wantWithoutPlan: 0.0,
			wantErr:         true,
			taxError:        tax.ErrInvalidOriginOrDestination,
		},
		{
			args:            args{plan: "invalid plan"},
			wantWithPlan:    0.0,
			wantWithoutPlan: 0.0,
			wantErr:         true,
			taxError:        tax.ErrPlanNotFound,
		},
	}

	for _, tt := range tests {
		tt.name = fmt.Sprintf("Test origin %s destination %s plan %s duration %f", tt.args.origin, tt.args.destination, tt.args.plan, tt.args.duration)
		t.Run(tt.name, func(t *testing.T) {
			calculatorRepositoryMock := repositoryTaxMock.New()
			c, err := tax.NewTaxCalculator(calculatorRepositoryMock)
			if err != nil {
				t.Fatalf("Error creating tax calculator: %s", err.Error())
			}

			gotWithPlan, gotWithoutPlan, err := c.Calculate(tt.args.origin, tt.args.destination, tt.args.plan, tt.args.duration)
			if (err != nil) != tt.wantErr {
				t.Errorf("Calculate() error = %v, wantErr %v", err, tt.wantErr)
				return
			} else if tt.wantErr && err != tt.taxError {
				t.Errorf("Calculate() error = %v, err %v", err, tt.taxError)
				return
			}

			if gotWithPlan != tt.wantWithPlan {
				t.Errorf("Calculate() gotWithPlan = %v, want %v", gotWithPlan, tt.wantWithPlan)
			}

			if gotWithoutPlan != tt.wantWithoutPlan {
				t.Errorf("Calculate() gotWithoutPlan = %v, want %v", gotWithoutPlan, tt.wantWithoutPlan)
			}
		})
	}

	t.Run("Test invalid repository", func(t *testing.T) {
		taxCalculator, err := tax.NewTaxCalculator(nil)
		assert.Nil(t, taxCalculator, "Tax calculator should be nil")
		assert.Error(t, err, "Error should be returned")
		assert.EqualError(t, err, tax.ErrInvalidRepository.Error(), "expected error %s, got %s", tax.ErrInvalidRepository.Error(), err.Error())
	})
}
