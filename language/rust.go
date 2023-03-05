package language

import "regexp"

var RUST_LANGUAGE_PATTERN = []LanguagePattern{
	{
		Pattern: regexp.MustCompile(`fn\s+main\(\)`),
		Type:    TYPE_KEYWORD_FUNCTION,
	},
	{
		Pattern: regexp.MustCompile(`(pub\s)?fn\s[A-Za-z0-9<>,]+\(.+\)\s->\s\w+(\s\{|)`),
		Type:    TYPE_KEYWORD_VISIBILITY,
	},
	{
		Pattern: regexp.MustCompile(`let\s+mut\s\w+(\s=|)`),
		Type:    TYPE_KEYWORD_VARIABLE,
	},
	{
		Pattern: regexp.MustCompile(`(.*)!\(.+\)`),
		Type:    TYPE_MACRO,
	},
	{
		Pattern: regexp.MustCompile(`use\s+\w+::.+`),
		Type:    TYPE_META_IMPORT,
	},
	{
		Pattern: regexp.MustCompile(`\{:\?\}`),
		Type:    TYPE_KEYWORD_OTHER,
	},
	{
		Pattern: regexp.MustCompile(`loop\s+\{`),
		Type:    TYPE_KEYWORD_CONTROL,
	},
	// Rust keywords
	{
		Pattern: regexp.MustCompile(`(impl|crate|extern|macro|box)`),
		Type:    TYPE_KEYWORD_OTHER,
	},
	{
		Pattern: regexp.MustCompile(`match\s+\w+\s+\{`),
		Type:    TYPE_KEYWORD_CONTROL,
	},
	{
		Pattern: regexp.MustCompile(`\w+\.len\(\)`),
		Type:    TYPE_KEYWORD_OTHER,
	},
	// Data types
	{
		Pattern: regexp.MustCompile(`(&str|(i|u)(8|16|32|64|128|size))`),
		Type:    TYPE_CONSTANT_TYPE,
	},
	// Vector
	{
		Pattern: regexp.MustCompile(`(Vec|Vec::new)|vec!`),
		Type:    TYPE_CONSTANT_TYPE,
	},
	// Traits
	{
		Pattern: regexp.MustCompile(`(Ok|Err|Box|ToOwned|Clone)`),
		Type:    TYPE_KEYWORD_OTHER,
	},
	// Panic!!
	{
		Pattern: regexp.MustCompile(`panic!\(.+\)`),
		Type:    TYPE_KEYWORD_FUNCTION,
	},
	// Avoiding clash with C#
	{
		Pattern: regexp.MustCompile(`using\s+System`),
		Type:    TYPE_NOT,
	},
	{
		Pattern: regexp.MustCompile(`Console\.WriteLine\s*\(`),
		Type:    TYPE_NOT,
	},
	{
		Pattern: regexp.MustCompile(`(public\s)?((partial|static)\s)?class\s+`),
		Type:    TYPE_NOT,
	},
	{
		Pattern: regexp.MustCompile(`(function|func)\s+`),
		Type:    TYPE_NOT,
	},
}
