package reqvalidator

import (
	"context"
	"github.com/guregu/null"
	"testing"
)

type TestStruct struct {
	Phone    null.String `json:"phone" validate:"required"`
	Password null.String `json:"password" validate:"omitempty,min=1,max=25"`
}

func TestValidator(t *testing.T) {
	testS := TestStruct{Phone: null.StringFrom("1234567890")}

	err := validate.StructCtx(context.Background(), testS)
	if err != nil {
		t.Errorf("Error: %v", err)
	}

}
