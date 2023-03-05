package language

import "regexp"

var CSS_LANGUAGE_PATTERN = []LanguagePattern{
	// Properties
	{
		Pattern: regexp.MustCompile(`[a-z-]+:[^:]+;`),
		Type:    TYPE_KEYWORD,
	},
	// <style> tag from HTML
	{
		Pattern: regexp.MustCompile(`<(\/)?style>`),
		Type:    TYPE_NOT,
	},
}
