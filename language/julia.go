package language

import "regexp"

var JULIA_LANGUAGE_PATTERN = []LanguagePattern{
	// Module import
	{
		Pattern: regexp.MustCompile(`(using)\s\w+`),
		Type:    TYPE_META_IMPORT,
	},
	{
		Pattern: regexp.MustCompile(`(bare\s)?module`),
		Type:    TYPE_META_MODULE,
	},
	// Avoiding Python's import
	{
		Pattern: regexp.MustCompile(`from\s.+import\s.+`),
		Type:    TYPE_NOT,
	},
	// Stdout / print line
	{
		Pattern: regexp.MustCompile(`println\(.*\)`),
		Type:    TYPE_KEYWORD_PRINT,
	},
	{
		Pattern: regexp.MustCompile(`(.*)!\(.*\)`),
		Type:    TYPE_MACRO,
	},
	// for x in / for x =
	{
		Pattern: regexp.MustCompile(`for\s(\w+)\s(in|=)\s`),
		Type:    TYPE_KEYWORD_CONTROL,
	},
	// It's not Julia if the function ends with {
	{
		Pattern: regexp.MustCompile(`function\s\w+\(.*\)\s\{`),
		Type:    TYPE_NOT,
	},
	// It's not Julia either if the while loop has a brackets
	{
		Pattern: regexp.MustCompile(`while\s+\(.+\)\n`),
		Type:    TYPE_NOT,
	},
	// The end keyword
	{
		Pattern: regexp.MustCompile(`end\n?`),
		Type:    TYPE_KEYWORD,
	},
	// Struct with <: annotation
	{
		Pattern: regexp.MustCompile(`struct\s(.*)\s<:\s`),
		Type:    TYPE_KEYWORD_OTHER,
	},
	// Data types
	{
		Pattern: regexp.MustCompile(`(::)?(Int|Uint)(8|16|32|64|128)`),
		Type:    TYPE_KEYWORD_VARIABLE,
	},
	{
		Pattern: regexp.MustCompile(`[0-9]+im`),
		Type:    TYPE_KEYWORD,
	},
	// Avoiding Rust confusion
	{
		Pattern: regexp.MustCompile(`\{:\?\}`),
		Type:    TYPE_NOT,
	},
	{
		Pattern: regexp.MustCompile(`fn\smain\(\)`),
		Type:    TYPE_NOT,
	},
	// Avoiding Ruby confusion
	{
		Pattern: regexp.MustCompile(`def\s+\w+\s*(\(.+\))?\s*\n`),
		Type:    TYPE_NOT,
	},
	{
		Pattern: regexp.MustCompile(`puts\s+("|').+("|')`),
		Type:    TYPE_NOT,
	},
	{
		Pattern: regexp.MustCompile(`class\s`),
		Type:    TYPE_NOT,
	},
	// Avoiding Lua confusion
	{
		Pattern: regexp.MustCompile(`local\s(function|\w+)`),
		Type:    TYPE_NOT,
	},
	{
		Pattern: regexp.MustCompile(`\bmodule\(.*\)`),
		Type:    TYPE_NOT,
	},
	// Avoiding Kotlin confusion
	{
		Pattern: regexp.MustCompile(`fun main\((.*)?\) {`),
		Type:    TYPE_NOT,
	},
	{
		Pattern: regexp.MustCompile(`fun(\s+)([A-Za-z0-9_])(\s+)?\((.*)\)(\s+){`),
		Type:    TYPE_NOT,
	},
}
