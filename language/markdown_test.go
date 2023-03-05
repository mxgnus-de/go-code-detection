package language_test

import (
	"os"
	"testing"

	detection "github.com/mxgnus-de/go-code-detection"
	"github.com/mxgnus-de/go-code-detection/language"
)

func TestMarkdownHeadingLanguageDetection(t *testing.T) {
	t.Log("Testing Markdown Headings Language Detection")

	t.Run("Heading 2", func(t *testing.T) {
		t.Log("Testing Markdown Heading Level 2")
		code := `## Heading level 2`
		result := detection.DetectLanguage(code, detection.NewDetectionOptions())

		if result.Language != language.LANGUAGE_MARKDOWN {
			t.Errorf("Expected language to be %s, but got %s", language.LANGUAGE_MARKDOWN, result.Language)
			t.Errorf("Result: %v", result)
		}
	})

	t.Run("Heading 3", func(t *testing.T) {
		t.Log("Testing Markdown Heading Level 3")
		code := `### Heading level 3`
		result := detection.DetectLanguage(code, detection.NewDetectionOptions())

		if result.Language != language.LANGUAGE_MARKDOWN {
			t.Errorf("Expected language to be %s, but got %s", language.LANGUAGE_MARKDOWN, result.Language)
			t.Errorf("Result: %v", result)
		}
	})

	t.Run("Heading 4", func(t *testing.T) {
		t.Log("Testing Markdown Heading Level 4")
		code := `#### Heading level 4`
		result := detection.DetectLanguage(code, detection.NewDetectionOptions())

		if result.Language != language.LANGUAGE_MARKDOWN {
			t.Errorf("Expected language to be %s, but got %s", language.LANGUAGE_MARKDOWN, result.Language)
			t.Errorf("Result: %v", result)
		}
	})

	t.Run("Heading 5", func(t *testing.T) {
		t.Log("Testing Markdown Heading Level 5")
		code := `##### Heading level 5`
		result := detection.DetectLanguage(code, detection.NewDetectionOptions())

		if result.Language != language.LANGUAGE_MARKDOWN {
			t.Errorf("Expected language to be %s, but got %s", language.LANGUAGE_MARKDOWN, result.Language)
			t.Errorf("Result: %v", result)
		}
	})

	t.Run("Heading 6", func(t *testing.T) {
		t.Log("Testing Markdown Heading Level 6")
		code := `###### Heading level 6`
		result := detection.DetectLanguage(code, detection.NewDetectionOptions())

		if result.Language != language.LANGUAGE_MARKDOWN {
			t.Errorf("Expected language to be %s, but got %s", language.LANGUAGE_MARKDOWN, result.Language)
			t.Errorf("Result: %v", result)
		}
	})

	t.Run("1 - alternative syntax", func(t *testing.T) {
		t.Log("Testing Markdown Heading Level 1")
		code := "Heading level 1\n============"
		result := detection.DetectLanguage(code, detection.NewDetectionOptions())

		if result.Language != language.LANGUAGE_MARKDOWN {
			t.Errorf("Expected language to be %s, but got %s", language.LANGUAGE_MARKDOWN, result.Language)
			t.Errorf("Result: %v", result)
		}
	})

	t.Run("2 - alternative syntax", func(t *testing.T) {
		t.Log("Testing Markdown Heading Level 1")
		code := "Heading level 1\n------------"
		result := detection.DetectLanguage(code, detection.NewDetectionOptions())

		if result.Language != language.LANGUAGE_MARKDOWN {
			t.Errorf("Expected language to be %s, but got %s", language.LANGUAGE_MARKDOWN, result.Language)
			t.Errorf("Result: %v", result)
		}
	})
}

func TestMarkdownImageLanguageDetection(t *testing.T) {
	t.Log("Testing Markdown Image Language Detection")

	t.Run("1", func(t *testing.T) {
		t.Log("Testing Markdown Image")
		code := `![Alt text](/path/to/img.jpg)`
		result := detection.DetectLanguage(code, detection.NewDetectionOptions())

		if result.Language != language.LANGUAGE_MARKDOWN {
			t.Errorf("Expected language to be %s, but got %s", language.LANGUAGE_MARKDOWN, result.Language)
			t.Errorf("Result: %v", result)
		}
	})

	t.Run("2", func(t *testing.T) {
		t.Log("Testing Markdown Image with title")
		code := `![Alt text](/path/to/img.jpg "Optional title")`
		result := detection.DetectLanguage(code, detection.NewDetectionOptions())

		if result.Language != language.LANGUAGE_MARKDOWN {
			t.Errorf("Expected language to be %s, but got %s", language.LANGUAGE_MARKDOWN, result.Language)
			t.Errorf("Result: %v", result)
		}
	})
}

func TestMarkdownLinkLanguageDetection(t *testing.T) {
	t.Log("Testing Markdown Link Language Detection")

	t.Run("1", func(t *testing.T) {
		t.Log("Testing Markdown Link")
		code := `[I'm an inline-style link](https://www.google.com)`
		result := detection.DetectLanguage(code, detection.NewDetectionOptions())

		if result.Language != language.LANGUAGE_MARKDOWN {
			t.Errorf("Expected language to be %s, but got %s", language.LANGUAGE_MARKDOWN, result.Language)
			t.Errorf("Result: %v", result)
		}
	})

	t.Run("2", func(t *testing.T) {
		t.Log("Testing Markdown Link with title")
		code := `[I'm an inline-style link](https://www.google.com "Google's Homepage")`
		result := detection.DetectLanguage(code, detection.NewDetectionOptions())

		if result.Language != language.LANGUAGE_MARKDOWN {
			t.Errorf("Expected language to be %s, but got %s", language.LANGUAGE_MARKDOWN, result.Language)
			t.Errorf("Result: %v", result)
		}
	})

	t.Run("3", func(t *testing.T) {
		t.Log("Testing Markdown Reference Link")
		code := `[I'm a reference-style link][Arbitrary case-insensitive reference text]`
		result := detection.DetectLanguage(code, detection.NewDetectionOptions())

		if result.Language != language.LANGUAGE_MARKDOWN {
			t.Errorf("Expected language to be %s, but got %s", language.LANGUAGE_MARKDOWN, result.Language)
			t.Errorf("Result: %v", result)
		}
	})

	t.Run("4", func(t *testing.T) {
		t.Log("Testing Markdown Reference Link with hash")
		code := `[1]: https://en.wikipedia.org/wiki/Hobbit#Lifestyle`
		result := detection.DetectLanguage(code, detection.NewDetectionOptions())

		if result.Language != language.LANGUAGE_MARKDOWN {
			t.Errorf("Expected language to be %s, but got %s", language.LANGUAGE_MARKDOWN, result.Language)
			t.Errorf("Result: %v", result)
		}
	})

	t.Run("5", func(t *testing.T) {
		t.Log("Testing Markdown Reference Link with Hobbit lifestyles")
		code := `[1]: <https://en.wikipedia.org/wiki/Hobbit#Lifestyle> "Hobbit lifestyles"`
		result := detection.DetectLanguage(code, detection.NewDetectionOptions())

		if result.Language != language.LANGUAGE_MARKDOWN {
			t.Errorf("Expected language to be %s, but got %s", language.LANGUAGE_MARKDOWN, result.Language)
			t.Errorf("Result: %v", result)
		}
	})
}

func TestMarkdownBlockQuotesLanguageDetection(t *testing.T) {
	t.Log("Testing Markdown Block Quotes Language Detection")
	code := `> We're living the future so
   > the present is our past.`
	result := detection.DetectLanguage(code, detection.NewDetectionOptions())

	if result.Language != language.LANGUAGE_MARKDOWN {
		t.Errorf("Expected language to be %s, but got %s", language.LANGUAGE_MARKDOWN, result.Language)
		t.Errorf("Result: %v", result)
	}
}

func TestMarkdownExampleLanguageDetection(t *testing.T) {
	t.Log("Testing Markdown Example Language Detection")

	t.Run("1", func(t *testing.T) {
		t.Log("Testing Markdown Example 1")
		fileBytes, err := os.ReadFile("../testdata/markdown/example1.md")

		if err != nil {
			t.Fatal(err)
		}

		code := string(fileBytes)
		result := detection.DetectLanguage(code, detection.NewDetectionOptions())

		if result.Language != language.LANGUAGE_MARKDOWN {
			t.Errorf("Expected language to be %s, but got %s", language.LANGUAGE_MARKDOWN, result.Language)
			t.Errorf("Result: %v", result)
		}
	})

	t.Run("2", func(t *testing.T) {
		t.Log("Testing Markdown Example 2")
		fileBytes, err := os.ReadFile("../testdata/markdown/example2.md")

		if err != nil {
			t.Fatal(err)
		}

		code := string(fileBytes)
		result := detection.DetectLanguage(code, detection.NewDetectionOptions())

		if result.Language != language.LANGUAGE_MARKDOWN {
			t.Errorf("Expected language to be %s, got %s", language.LANGUAGE_MARKDOWN, result.Language)
			t.Errorf("Result: %v", result)
		}
	})

	t.Run("3", func(t *testing.T) {
		t.Log("Testing Markdown Example 3")
		fileBytes, err := os.ReadFile("../testdata/markdown/example3.md")

		if err != nil {
			t.Fatal(err)
		}

		code := string(fileBytes)
		result := detection.DetectLanguage(code, detection.NewDetectionOptions())

		if result.Language != language.LANGUAGE_MARKDOWN {
			t.Errorf("Expected language to be %s, got %s", language.LANGUAGE_MARKDOWN, result.Language)
			t.Errorf("Result: %v", result)
		}
	})

	t.Run("4", func(t *testing.T) {
		t.Log("Testing Markdown Example 4")
		fileBytes, err := os.ReadFile("../testdata/markdown/example4.md")

		if err != nil {
			t.Fatal(err)
		}

		code := string(fileBytes)
		result := detection.DetectLanguage(code, detection.NewDetectionOptions())

		if result.Language != language.LANGUAGE_MARKDOWN {
			t.Errorf("Expected language to be %s, got %s", language.LANGUAGE_MARKDOWN, result.Language)
			t.Errorf("Result: %v", result)
		}
	})
}
