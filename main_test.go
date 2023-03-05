package detection_test

import (
	"testing"

	detection "github.com/mxgnus-de/go-code-detection"
	"github.com/mxgnus-de/go-code-detection/language"
)

func TestUnknwonLanguageDetection(t *testing.T) {
	t.Log("Testing Unknown Language Detection")

	t.Run("Plain Text", func(t *testing.T) {
		t.Log("Testing Plain Text Language Detection")
		code := `Plain text`
		result := detection.DetectLanguage(code, detection.NewDetectionOptions())

		if result.Language != language.LANGUAGE_UNKNOWN {
			t.Errorf("Expected language to be %s, but got %s", language.LANGUAGE_UNKNOWN, result.Language)
			t.Errorf("Result: %v", result)
		}
	})

	t.Run("Hello World", func(t *testing.T) {
		t.Log("Testing Hello World Language Detection")
		code := `Hello World`
		result := detection.DetectLanguage(code, detection.NewDetectionOptions())

		if result.Language != language.LANGUAGE_UNKNOWN {
			t.Errorf("Expected language to be %s, but got %s", language.LANGUAGE_UNKNOWN, result.Language)
			t.Errorf("Result: %v", result)
		}
	})

	t.Run("Should not return unknown", func(t *testing.T) {
		t.Log("Testing Hello World Language Detection")
		code := `Random text`
		result := detection.DetectLanguage(code, detection.NewDetectionOptions().WithNoUnknown())

		if result.Language == language.LANGUAGE_UNKNOWN {
			t.Errorf("Expected language to be not %s, but got %s", language.LANGUAGE_UNKNOWN, result.Language)
			t.Errorf("Result: %v", result)
		}
	})
}
