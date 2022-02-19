package phone_call_test

import (
	"testing"

	"github.com/pedrokunz/go_backend/entity/tax"
	repositoryTaxMock "github.com/pedrokunz/go_backend/infra/repository/mock/tax"
	"github.com/pedrokunz/go_backend/usecase/phone_call"
)

func TestSuccess(t *testing.T) {
	calculatorRepositoryMock := repositoryTaxMock.New()
	calculator, err := tax.NewTaxCalculator(calculatorRepositoryMock)
	if err != nil {
		t.Fatalf("Error creating tax calculator: %s", err)
	}

	usecase := phone_call.NewCalculateTax(calculator)

	t.Run("plan FaleMais30 during 20 minutes", func(t *testing.T) {
		input := &phone_call.CalculateTaxInput{
			Plan:        "FaleMais30",
			Duration:    20,
			Origin:      "011",
			Destination: "016",
		}

		withPlanExpected := 0.0
		withoutPlanExpected := 38.0

		output, err := usecase.Perform(input)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if output.WithPlan != withPlanExpected {
			t.Fatalf("expected %f, got %f", withPlanExpected, output.WithPlan)
		}

		if output.WithoutPlan != withoutPlanExpected {
			t.Fatalf("expected %f, got %f", withoutPlanExpected, output.WithoutPlan)
		}
	})
}
