package language

import "regexp"

var HTML_LANGUAGE_PATTERN = []LanguagePattern{
	{
		Pattern:   regexp.MustCompile(`<!DOCTYPE (html|HTML PUBLIC .+)>`),
		Type:      TYPE_META_MODULE,
		IsNearTop: true,
	},
	{
		Pattern: regexp.MustCompile(`<[a-z0-9]+(\s*[\w]+=('|").+('|")\s*)?>.*<\/[a-z0-9]+>`),
		Type:    TYPE_KEYWORD,
	},
	{
		Pattern: regexp.MustCompile(`<!--(.*)(-->)?`),
		Type:    TYPE_COMMENT_BLOCK,
	},
	{
		Pattern: regexp.MustCompile(`[a-z-]+=("|').+("|')`),
		Type:    TYPE_KEYWORD_OTHER,
	},
	{
		Pattern: regexp.MustCompile(`<\?php`),
		Type:    TYPE_NOT,
	},
}
