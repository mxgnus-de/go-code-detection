package language

import "regexp"

var PHP_LANGUAGE_PATTERN = []LanguagePattern{
	// PHP tag
	{
		Pattern: regexp.MustCompile(`<\?php`),
		Type:    TYPE_META_MODULE,
	},
	// PHP style variables.
	{
		Pattern: regexp.MustCompile(`\$\w+`),
		Type:    TYPE_KEYWORD_VARIABLE,
	},
	// use Something\Something;
	{
		Pattern:   regexp.MustCompile(`use\s+\w+(\\\w+)+\s*;`),
		Type:      TYPE_META_IMPORT,
		IsNearTop: true,
	},
	// arrow
	{
		Pattern: regexp.MustCompile(`\$\w+->\w+`),
		Type:    TYPE_KEYWORD,
	},
	// require/include
	{
		Pattern: regexp.MustCompile(`(require|include)(_once)?\s*\(?\s*('|").+\.php('|")\s*\)?\s*;`),
		Type:    TYPE_META_IMPORT,
	},
	// echo 'something';
	{
		Pattern: regexp.MustCompile(`echo\s+('|").+('|")\s*;`),
		Type:    TYPE_KEYWORD_PRINT,
	},
	// NULL constant
	{
		Pattern: regexp.MustCompile(`NULL`),
		Type:    TYPE_CONSTANT_NULL,
	},
	// new keyword
	{
		Pattern: regexp.MustCompile(`new\s+((\\\w+)+|\w+)(\(.*\))?`),
		Type:    TYPE_KEYWORD,
	},
	// Function definition
	{
		Pattern: regexp.MustCompile(`function(\s+[$\w]+\(.*\)|\s*\(.*\))`),
		Type:    TYPE_KEYWORD_CONTROL,
	},
	// (else)if statement
	{
		Pattern: regexp.MustCompile(`(else)?if\s+\(.+\)`),
		Type:    TYPE_KEYWORD_CONTROL,
	},
	// scope operator
	{
		Pattern: regexp.MustCompile(`\w+::\w+`),
		Type:    TYPE_KEYWORD,
	},
	// === operator
	{
		Pattern: regexp.MustCompile(`===`),
		Type:    TYPE_KEYWORD_OPERATOR,
	},
	// !== operator
	{
		Pattern: regexp.MustCompile(`!==`),
		Type:    TYPE_KEYWORD_OPERATOR,
	},
	// C/JS style variable declaration.
	{
		Pattern: regexp.MustCompile(`(^|\s)(var|char|long|int|float|double)\s+\w+\s*=?`),
		Type:    TYPE_NOT,
	},
	// Javascript variable declaration
	{
		Pattern: regexp.MustCompile(`(var|const|let)\s+\w+\s*=?`),
		Type:    TYPE_NOT,
	},
	// Avoiding Lua confusion
	{
		Pattern: regexp.MustCompile(`local\s(function|\w+)`),
		Type:    TYPE_NOT,
	},
}
