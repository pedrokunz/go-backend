package contract_test

import (
	"testing"

	"github.com/pedrokunz/go_backend/entity"
)

func TestTaxCalculator(t *testing.T) {
	t.Run("calculate tax with FaleMais30 with 20 minutes duration", func(t *testing.T) {
		taxCalculator := entity.NewTaxCalculator()

		withPlanExpected := 0.0
		withoutPlanExpected := 38.0

		withPlan, withoutPlan, err := taxCalculator.Calculate("011", "016", "FaleMais30", 20)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if withPlan != withPlanExpected {
			t.Fatalf("expected %f, got %f", withPlanExpected, withPlan)
		}

		if withoutPlan != withoutPlanExpected {
			t.Fatalf("expected %f, got %f", withoutPlanExpected, withoutPlan)
		}
	})
}
