package inivadil_test

import (
	"testing"

	"github.com/ItsMalma/inivadil"
)

func TestValidationError(t *testing.T) {
	t.Run("with NewValidationError function", func(t *testing.T) {
		t.Run("empty field and message", func(t *testing.T) {
			err := inivadil.NewValidationError("", "")
			if err.Field != "" {
				t.Fatalf("expected err.Field = \"\", but got \"%s\" instead", err.Field)
			}
			if err.Message != "" {
				t.Fatalf("expected err.Message = \"\", but got \"%s\" instead", err.Message)
			}
			if err.Error() != "" {
				t.Fatalf("expected err.Error() = \"\", but got \"%s\" instead", err.Error())
			}
			if !err.Nil() {
				t.Fatal("expedted err.Nil() = true, but got false instead")
			}
		})

		t.Run("with field only", func(t *testing.T) {
			err := inivadil.NewValidationError("x", "")
			if err.Field != "x" {
				t.Fatalf("expected err.Field = \"x\", but got \"%s\" instead", err.Field)
			}
			if err.Message != "" {
				t.Fatalf("expected err.Message = \"\", but got \"%s\" instead", err.Message)
			}
			if err.Error() != "x: " {
				t.Fatalf("expected err.Error() = \"x: \", but got \"%s\" instead", err.Error())
			}
			if err.Nil() {
				t.Fatal("expedted err.Nil() = false, but got true instead")
			}
		})

		t.Run("with message only", func(t *testing.T) {
			err := inivadil.NewValidationError("", "its x")
			if err.Field != "" {
				t.Fatalf("expected err.Field = \"\", but got \"%s\" instead", err.Field)
			}
			if err.Message != "its x" {
				t.Fatalf("expected err.Message = \"its x\", but got \"%s\" instead", err.Message)
			}
			if err.Error() != ": its x" {
				t.Fatalf("expected err.Error() = \": its x\", but got \"%s\" instead", err.Error())
			}
			if err.Nil() {
				t.Fatal("expedted err.Nil() = false, but got true instead")
			}
		})

		t.Run("with both field and message", func(t *testing.T) {
			err := inivadil.NewValidationError("x", "its x")
			if err.Field != "x" {
				t.Fatalf("expected err.Field = \"x\", but got \"%s\" instead", err.Field)
			}
			if err.Message != "its x" {
				t.Fatalf("expected err.Message = \"its x\", but got \"%s\" instead", err.Message)
			}
			if err.Error() != "x: its x" {
				t.Fatalf("expected err.Error() = \"x: its x\", but got \"%s\" instead", err.Error())
			}
			if err.Nil() {
				t.Fatal("expedted err.Nil() = false, but got true instead")
			}
		})
	})

	t.Run("with NilValidationError function", func(t *testing.T) {
		err := inivadil.NilValidationError()
		if err.Field != "" {
			t.Fatalf("expected err.Field = \"\", but got \"%s\" instead", err.Field)
		}
		if err.Message != "" {
			t.Fatalf("expected err.Message = \"\", but got \"%s\" instead", err.Message)
		}
		if err.Error() != "" {
			t.Fatalf("expected err.Error() = \"\", but got \"%s\" instead", err.Error())
		}
		if !err.Nil() {
			t.Fatal("expedted err.Nil() = true, but got false instead")
		}
	})
}

func TestValidationErrors(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		validationErrors := inivadil.ValidationErrors{}
		if str := validationErrors.Error(); str != "" {
			t.Fatalf("expected \"\", but got \"%s\" instead", str)
		}
	})

	t.Run("1 error", func(t *testing.T) {
		t.Run("with NewValidationError function", func(t *testing.T) {
			validationErrors := inivadil.ValidationErrors{}
			validationErrors = append(validationErrors, inivadil.NewValidationError("x", "this is x"))
			if str := validationErrors.Error(); str != "x: this is x" {
				t.Fatalf("expected \"x: this is x\", but got \"%s\" instead", str)
			}
		})

		t.Run("with NilValidationError function", func(t *testing.T) {
			validationErrors := inivadil.ValidationErrors{}
			validationErrors = append(validationErrors, inivadil.NilValidationError())
			if str := validationErrors.Error(); str != "" {
				t.Fatalf("expected \"\", but got \"%s\" instead", str)
			}
		})
	})

	t.Run("2 error", func(t *testing.T) {
		t.Run("with NewValidationError function", func(t *testing.T) {
			validationErrors := inivadil.ValidationErrors{}
			validationErrors = append(validationErrors, inivadil.NewValidationError("x", "this is x"))
			validationErrors = append(validationErrors, inivadil.NewValidationError("y", "this is y"))
			if str := validationErrors.Error(); str != "x: this is x\ny: this is y" {
				t.Fatalf("expected \"x: this is x\", but got \"%s\" instead", str)
			}
		})

		t.Run("with NilValidationError function", func(t *testing.T) {
			validationErrors := inivadil.ValidationErrors{}
			validationErrors = append(validationErrors, inivadil.NilValidationError())
			validationErrors = append(validationErrors, inivadil.NilValidationError())
			if str := validationErrors.Error(); str != "" {
				t.Fatalf("expected \"\", but got \"%s\" instead", str)
			}
		})

		t.Run("with NewValidationError and NilValidationError function", func(t *testing.T) {
			t.Run("NewValidationError first", func(t *testing.T) {
				validationErrors := inivadil.ValidationErrors{}
				validationErrors = append(validationErrors, inivadil.NewValidationError("x", "this is x"))
				validationErrors = append(validationErrors, inivadil.NilValidationError())
				if str := validationErrors.Error(); str != "x: this is x" {
					t.Fatalf("expected \"\", but got \"%s\" instead", str)
				}
			})

			t.Run("NilValidationError first", func(t *testing.T) {
				validationErrors := inivadil.ValidationErrors{}
				validationErrors = append(validationErrors, inivadil.NilValidationError())
				validationErrors = append(validationErrors, inivadil.NewValidationError("y", "this is y"))
				if str := validationErrors.Error(); str != "y: this is y" {
					t.Fatalf("expected \"\", but got \"%s\" instead", str)
				}
			})
		})
	})
}
