package detection_test

import (
	"testing"

	"github.com/mxgnus-de/go-code-detection/detection"
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

func TestLowerCaseOptionLanguageDetection(t *testing.T) {
	t.Log("Testing Lower Case Option Language Detection")

	t.Run("Test C++", func(t *testing.T) {
		code := `cout << \"Hello world\" << endl;`
		result := detection.DetectLanguage(code, detection.NewDetectionOptions().WithLowerCaseOutput().WithNoUnknown())

		if language.ConvertLowerCaseLangIntoLang(result.Language) != language.LANGUAGE_CPP {
			t.Errorf("Expected language to be %s, but got %s", language.LANGUAGE_CPP, result.Language)
			t.Errorf("Result: %v", result)
		}
	})

	t.Run("Test C#", func(t *testing.T) {
		code := `Console.WriteLine(\"Hello world\");`
		result := detection.DetectLanguage(code, detection.NewDetectionOptions().WithLowerCaseOutput().WithNoUnknown())

		if language.ConvertLowerCaseLangIntoLang(result.Language) != language.LANGUAGE_CS {
			t.Errorf("Expected language to be %s, but got %s", language.LANGUAGE_CS, result.Language)
			t.Errorf("Result: %v", result)
		}
	})

	t.Run("Test Go", func(t *testing.T) {
		code := `package main

      import "fmt"

      func main() {
         fmt.Println("Hello World")
      }`
		result := detection.DetectLanguage(code, detection.NewDetectionOptions().WithLowerCaseOutput().WithNoUnknown())

		if result.Language != "go" {
			t.Errorf("Expected language to be %s, but got %s", "go", result.Language)
			t.Errorf("Result: %v", result)
		} else if language.ConvertLowerCaseLangIntoLang(result.Language) != language.LANGUAGE_GO {
			t.Errorf("Expected converted language to be %s, but got %s", language.LANGUAGE_GO, language.ConvertLowerCaseLangIntoLang(result.Language))
			t.Errorf("Result: %v", result)
		}
	})
}
