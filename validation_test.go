package inivadil_test

import (
	"bytes"
	"errors"
	"testing"

	"github.com/ItsMalma/inivadil"
)

func TestValidateFunction(t *testing.T) {
	t.Run("not error", func(t *testing.T) {
		if err := inivadil.Validate(
			inivadil.JSONValidation,
			inivadil.Rules{},
			bytes.NewReader([]byte("{}")),
		); err != nil {
			t.Fatalf("expected return nil, but got %v instead", err)
		}
	})

	t.Run("error invalid type", func(t *testing.T) {
		err := inivadil.Validate(
			69,
			inivadil.Rules{},
			bytes.NewReader([]byte("{}")),
		)
		if !errors.Is(err, inivadil.ErrInvalidValidationType) {
			t.Fatalf("expected return err (ErrValidationInvalidType), but got %v instead", err)
		}
	})
}

func TestValidateFunctionJSON(t *testing.T) {
	t.Run("not error", func(t *testing.T) {
		if err := inivadil.Validate(
			inivadil.JSONValidation,
			inivadil.Rules{},
			bytes.NewReader([]byte("{}")),
		); err != nil {
			t.Fatalf("expected return nil, but got %v instead", err)
		}
	})

	t.Run("error", func(t *testing.T) {
		if err := inivadil.Validate(
			inivadil.JSONValidation,
			inivadil.Rules{},
			bytes.NewReader([]byte("not a json")),
		); err == nil {
			t.Fatalf("expected return err, but got nil instead")
		}
	})
}
