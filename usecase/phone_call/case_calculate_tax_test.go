package phone_call_test

import (
	"testing"

	repositoryTaxMock "github.com/pedrokunz/go_backend/infra/repository/mock/tax"
	"github.com/pedrokunz/go_backend/usecase/phone_call"
)

func TestSuccess(t *testing.T) {
	repositoryMock := repositoryTaxMock.New()
	usecase, err := phone_call.NewCalculateTax(repositoryMock)
	if err != nil {
		t.Fatalf("Error creating usecase: %s", err)
	}

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
