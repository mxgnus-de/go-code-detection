package language

import "regexp"

var RUBY_LANGUAGE_PATTERN = []LanguagePattern{
	// require/include
	{
		Pattern:   regexp.MustCompile(`(require|include)\s+'\w+(\.rb)?'`),
		Type:      TYPE_META_IMPORT,
		IsNearTop: true,
	},
	// Function definition
	{
		Pattern: regexp.MustCompile(`def\s+\w+\s*(\(.+\))?\s*\n`),
		Type:    TYPE_KEYWORD_FUNCTION,
	},
	// Instance variables
	{
		Pattern: regexp.MustCompile(`@\w+`),
		Type:    TYPE_KEYWORD_OTHER,
	},
	// Boolean property
	{
		Pattern: regexp.MustCompile(`\.\w+\?`),
		Type:    TYPE_CONSTANT_BOOLEAN,
	},
	// puts (Ruby print)
	{
		Pattern: regexp.MustCompile(`puts\s+("|').+("|')`),
		Type:    TYPE_KEYWORD_PRINT,
	},
	// Inheriting class
	{
		Pattern: regexp.MustCompile(`class [A-Z]\w*\s*<\s*([A-Z]\w*(::)?)+`),
		Type:    TYPE_KEYWORD,
	},
	// attr_accessor
	{
		Pattern: regexp.MustCompile(`attr_accessor\s+(:\w+(,\s*)?)+`),
		Type:    TYPE_KEYWORD_FUNCTION,
	},
	// new
	{
		Pattern: regexp.MustCompile(`\w+\.new\s+`),
		Type:    TYPE_KEYWORD,
	},
	// elsif keyword
	{
		Pattern: regexp.MustCompile(`elsif`),
		Type:    TYPE_KEYWORD_CONTROL,
	},
	// module
	{
		Pattern: regexp.MustCompile(`\bmodule\s\S`),
		Type:    TYPE_KEYWORD_OTHER,
	},
	// BEGIN and END
	{
		Pattern: regexp.MustCompile(`\bBEGIN\s\{.*\}`),
		Type:    TYPE_KEYWORD_OTHER,
	},
	{
		Pattern: regexp.MustCompile(`\bEND\s\{.*\}`),
		Type:    TYPE_KEYWORD_OTHER,
	},
	// do
	{
		Pattern: regexp.MustCompile(`do\s*[|]\w+(,\s*\w+)*[|]`),
		Type:    TYPE_KEYWORD_CONTROL,
	},
	// for loop
	{
		Pattern: regexp.MustCompile(`for (\w+|\(?\w+,\s*\w+\)?) in (.+)`),
		Type:    TYPE_KEYWORD_CONTROL,
	},
	// nil keyword
	{
		Pattern: regexp.MustCompile(`nil`),
		Type:    TYPE_CONSTANT_NULL,
	},
	// Scope operator
	{
		Pattern: regexp.MustCompile(`[A-Z]\w*::[A-Z]\w*`),
		Type:    TYPE_MACRO,
	},
}
