package language

import "regexp"

var MARKDOWN_LANGUAGE_PATTERN = []LanguagePattern{
	// headings
	{
		Pattern: regexp.MustCompile(`^(#){2,6}\s.+`),
		Type:    TYPE_KEYWORD,
	},
	// headings alternate syntax
	{
		Pattern: regexp.MustCompile(`^(=|-)(=|-)+([^>=]|$)$`),
		Type:    TYPE_META_MODULE,
	},
	// images
	{
		Pattern: regexp.MustCompile(`(!)?\[.+\]\(.+\)`),
		Type:    TYPE_KEYWORD,
	},
	// links 2
	{
		Pattern: regexp.MustCompile(`\[.+\]\[.+\]`),
		Type:    TYPE_KEYWORD,
	},
	// links 3
	{
		Pattern: regexp.MustCompile(`^\[.+\]:\s?(<)?(http)?`),
		Type:    TYPE_KEYWORD,
	},
	// blockquotes
	{
		Pattern: regexp.MustCompile(`^(> .*)+`),
		Type:    TYPE_MACRO,
	},
	// code block
	{
		Pattern: regexp.MustCompile("^```([A-Za-z0-9#_]+)?$"),
		Type:    TYPE_KEYWORD,
	},
	// frontmatter
	{
		Pattern:   regexp.MustCompile(`^---$`),
		Type:      TYPE_META_MODULE,
		IsNearTop: true,
	},
}
