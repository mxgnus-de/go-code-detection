package detection_test

import (
	"testing"

	"github.com/mxgnus-de/go-code-detection/detection"
	"github.com/mxgnus-de/go-code-detection/language"
)

func TestShebangLanguageDetection(t *testing.T) {
	t.Log("Testing Shebang Language Detection")

	t.Run("Empty", func(t *testing.T) {
		t.Log("Testing Empty Shebang Language Detection")
		code := ``
		result := detection.DetectLanguage(code, detection.NewDetectionOptions())

		if result.Language != language.LANGUAGE_UNKNOWN {
			t.Errorf("Expected language to be %s, but got %s", language.LANGUAGE_UNKNOWN, result.Language)
			t.Errorf("Result: %v", result)
		}
	})

	t.Run("JavaScript", func(t *testing.T) {
		t.Log("Testing JavaScript Shebang Language Detection")
		code := `#!/usr/bin/env node`
		result := detection.DetectLanguage(code, detection.NewDetectionOptions())

		if result.Language != language.LANGUAGE_JAVASCRIPT {
			t.Errorf("Expected language to be %s, but got %s", language.LANGUAGE_JAVASCRIPT, result.Language)
			t.Errorf("Result: %v", result)
		}
	})

	t.Run("Python", func(t *testing.T) {
		t.Log("Testing Python Shebang Language Detection")
		code := `#!/usr/bin/env python`
		result := detection.DetectLanguage(code, detection.NewDetectionOptions())

		if result.Language != language.LANGUAGE_PYTHON {
			t.Errorf("Expected language to be %s, but got %s", language.LANGUAGE_PYTHON, result.Language)
			t.Errorf("Result: %v", result)
		}
	})

	t.Run("Ruby", func(t *testing.T) {
		t.Log("Testing Ruby Shebang Language Detection")
		code := `#!/usr/bin/env ruby`
		result := detection.DetectLanguage(code, detection.NewDetectionOptions())

		if result.Language != language.LANGUAGE_RUBY {
			t.Errorf("Expected language to be %s, but got %s", language.LANGUAGE_RUBY, result.Language)
			t.Errorf("Result: %v", result)
		}
	})

	t.Run("Bash", func(t *testing.T) {
		t.Log("Testing Bash Shebang Language Detection")
		code := `#!/usr/bin/env bash`
		result := detection.DetectLanguage(code, detection.NewDetectionOptions())

		if result.Language != language.LANGUAGE_BASH {
			t.Errorf("Expected language to be %s, but got %s", language.LANGUAGE_BASH, result.Language)
			t.Errorf("Result: %v", result)
		}
	})

	t.Run("PHP", func(t *testing.T) {
		t.Log("Testing PHP Shebang Language Detection")
		code := `#!/usr/bin/env php`
		result := detection.DetectLanguage(code, detection.NewDetectionOptions())

		if result.Language != language.LANGUAGE_PHP {
			t.Errorf("Expected language to be %s, but got %s", language.LANGUAGE_PHP, result.Language)
			t.Errorf("Result: %v", result)
		}
	})

	t.Run("Perl", func(t *testing.T) {
		t.Log("Testing Perl Shebang Language Detection")
		code := `#!/usr/bin/env perl`
		result := detection.DetectLanguage(code, detection.NewDetectionOptions())

		if result.Language != language.LANGUAGE_PERL {
			t.Errorf("Expected language to be %s, but got %s", language.LANGUAGE_PERL, result.Language)
			t.Errorf("Result: %v", result)
		}
	})
}
