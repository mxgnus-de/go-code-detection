package language_test

import (
	"testing"

	detection "github.com/mxgnus-de/go-code-detection"
	"github.com/mxgnus-de/go-code-detection/language"
)

func TestJSONLanguageDetection(t *testing.T) {
	t.Log("Testing JSON Language Detection")

	t.Run("1", func(t *testing.T) {
		code := `{
      "key": "value",

      "keys": "must always be enclosed in double quotes",
      "numbers": 0,
      "strings": "Hellø, wørld. All unicode is allowed, along with \\"escaping\\".",
      "has bools?": true,
      "nothingness": null,

      "big number": 1.2e+100,

      "objects": {
         "comment": "Most of your structure will come from objects.",

         "array": [0, 1, 2, 3, "Arrays can have anything in them.", 5],

         "another object": {
            "comment": "These things can be nested, very useful."
         }
      },

      "silliness": [
         {
            "sources of potassium": ["bananas"]
         },
         [
            [1, 0, 0, 0],
            [0, 1, 0, 0],
            [0, 0, 1, "neo"],
            [0, 0, 0, 1]
         ]
      ],

      "alternative style": {
         "comment": "check this out!"
      , "comma position": "doesn't matter, if it's before the next key, it's valid"
      , "another comment": "how nice"
      },



      "whitespace": "Does not matter.",



      "that was short": "And done. You now know everything JSON has to offer."
      }`
		result := detection.DetectLanguage(code, detection.NewDetectionOptions())

		if result.Language != language.LANGUAGE_JSON {
			t.Errorf("Expected language to be %s, but got %s", language.LANGUAGE_JSON, result.Language)
			t.Errorf("Result: %v", result)
		}
	})

	t.Run("2", func(t *testing.T) {
		code := `{
         "name": "John Doe",
         "age": 43,
         "address": {
            "street": "10 Downing Street",
            "city": "London"
         },
         "phones": [
            "+44 1234567",
            "+44 2345678"
         ]
      }`
		result := detection.DetectLanguage(code, detection.NewDetectionOptions())

		if result.Language != language.LANGUAGE_JSON {
			t.Errorf("Expected language to be %s, but got %s", language.LANGUAGE_JSON, result.Language)
			t.Errorf("Result: %v", result)
		}
	})

	t.Run("3", func(t *testing.T) {
		code := `{
      "keywords": [
         "JSON",
         "server",
         "fake",
         "REST",
         "API",
         "prototyping",
         "mock",
         "mocking",
         "test",
         "testing",
         "rest",
         "data",
         "dummy",
         "sandbox"
      ]
      }`
		result := detection.DetectLanguage(code, detection.NewDetectionOptions())

		if result.Language != language.LANGUAGE_JSON {
			t.Errorf("Expected language to be %s, but got %s", language.LANGUAGE_JSON, result.Language)
			t.Errorf("Result: %v", result)
		}
	})

	t.Run("4", func(t *testing.T) {
		code := `{
         "/api/": "/",
         "/blog/:resource/:id/show": "/:resource/:id",
         "/blog/:category": "/posts?category=:category"
      }`
		result := detection.DetectLanguage(code, detection.NewDetectionOptions())

		if result.Language != language.LANGUAGE_JSON {
			t.Errorf("Expected language to be %s, but got %s", language.LANGUAGE_JSON, result.Language)
			t.Errorf("Result: %v", result)
		}
	})

	t.Run("5", func(t *testing.T) {
		code := `{
         "posts": [
            { "id": 1, "title": "json-server", "author": "typicode" }
         ],
         "comments": [
            { "id": 1, "body": "some comment", "postId": 1 }
         ],
         "profile": { "name": "typicode" }
      }`
		result := detection.DetectLanguage(code, detection.NewDetectionOptions())

		if result.Language != language.LANGUAGE_JSON {
			t.Errorf("Expected language to be %s, but got %s", language.LANGUAGE_JSON, result.Language)
			t.Errorf("Result: %v", result)
		}
	})

	t.Run("6", func(t *testing.T) {
		code := `{
         "middlewares": ["./fixtures/middlewares/en", "./fixtures/middlewares/jp"]
      }`
		result := detection.DetectLanguage(code, detection.NewDetectionOptions())

		if result.Language != language.LANGUAGE_JSON {
			t.Errorf("Expected language to be %s, but got %s", language.LANGUAGE_JSON, result.Language)
			t.Errorf("Result: %v", result)
		}
	})
}
