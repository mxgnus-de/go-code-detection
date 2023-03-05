package language_test

import (
	"testing"

	"github.com/mxgnus-de/go-code-detection/detection"
	"github.com/mxgnus-de/go-code-detection/language"
)

func TestCSSHelloWorldLanguageDetection(t *testing.T) {
	t.Log("Testing CSS Hello World Language Detection")
	code := `.hello-world {\n\tfont-size: 100px;\n}`
	result := detection.DetectLanguage(code, detection.NewDetectionOptions())

	if result.Language != language.LANGUAGE_CSS {
		t.Errorf("Expected language to be %s, but got %s", language.LANGUAGE_CSS, result.Language)
		t.Errorf("Result: %v", result)
	}
}

func TestCSSLongLanguageDetection(t *testing.T) {
	t.Log("Testing CSS Long Language Detection")
	code := `/**
   * Improve readability when focused and also mouse hovered in all browsers.
   */
   
   a:active,
   a:hover {
      outline: 0;
   }
   
   /**
   * Address styling not present in IE 8/9/10/11, Safari, and Chrome.
   */
   
   abbr[title] {
      border-bottom: 1px dotted;
   }`
	result := detection.DetectLanguage(code, detection.NewDetectionOptions())

	if result.Language != language.LANGUAGE_CSS {
		t.Errorf("Expected language to be %s, but got %s", language.LANGUAGE_CSS, result.Language)
		t.Errorf("Result: %v", result)
	}
}
