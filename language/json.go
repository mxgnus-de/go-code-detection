package language

import "regexp"

var JSON_LANGUAGE_PATTERN = []LanguagePattern{
	// object declaration on top
	{
		Pattern:   regexp.MustCompile(`^\{$`),
		Type:      TYPE_META_MODULE,
		IsNearTop: true,
	},
	// normal data type
	{
		Pattern: regexp.MustCompile(`^\s*".+":\s*(".+"|[0-9]+|null|true|false)(,)?$`),
		Type:    TYPE_KEYWORD,
	},
	// object and array
	{
		Pattern: regexp.MustCompile(`^\s*".+":\s*(\{|\[)$`),
		Type:    TYPE_KEYWORD,
	},
	// inline key/value pair in object
	// e.g { "id": 1, "body": "some comment", "postId": 1 }
	{
		Pattern: regexp.MustCompile(`^\s*".+":\s*\{(\s*".+":\s*(".+"|[0-9]+|null|true|false)(,)?\s*){1,}\}(,)?$`),
		Type:    TYPE_KEYWORD,
	},
	// inline value in array
	// e.g "middlewares": ["./fixtures/middlewares/en", "./fixtures/middlewares/jp"]
	{
		Pattern: regexp.MustCompile(`\s*".+"\s*\[\s*((".+"|[0-9]+|null|true|false)(,)?\s*){1,}\](,)?$`),
		Type:    TYPE_KEYWORD,
	},
}
