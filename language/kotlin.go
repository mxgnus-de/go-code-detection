package language

import "regexp"

var KOTLIN_LANGUAGE_PATTERN = []LanguagePattern{
	// main function
	{
		Pattern: regexp.MustCompile(`fun main\((.*)?\) {`),
		Type:    TYPE_KEYWORD_FUNCTION,
	},
	// function
	{
		Pattern: regexp.MustCompile(`(inline|private|public|protected|override|operator(\s+))?fun(\s+)([A-Za-z0-9_])(\s+)?\((.*)\)(\s+)({|=)`),
		Type:    TYPE_KEYWORD_FUNCTION,
	},
	// println
	{
		Pattern: regexp.MustCompile(`println\((.*)\)(\n|;)`),
		Type:    TYPE_KEYWORD_PRINT,
	},
	// (else )if statement
	{
		Pattern: regexp.MustCompile(`(else )?if\s*\(.+\)`),
		Type:    TYPE_KEYWORD_CONTROL,
	},
	// while loop
	{
		Pattern: regexp.MustCompile(`while\s+\(.+\)`),
		Type:    TYPE_KEYWORD_CONTROL,
	},
	// Variables
	{
		Pattern: regexp.MustCompile(`(const)?(\s+)?val(\s+)(.*)(:(\s)(.*)(\?)?)?(\s+)=(\s+)`),
		Type:    TYPE_KEYWORD_VARIABLE,
	},
	{
		Pattern: regexp.MustCompile(`^(\s+)?(inner|open|data)(\s+)class`),
		Type:    TYPE_KEYWORD,
	},
	// import
	{
		Pattern:   regexp.MustCompile(`^import(\s+)(.*)$`),
		Type:      TYPE_META_IMPORT,
		IsNearTop: true,
	},
	// package
	{
		Pattern: regexp.MustCompile(`typealias(\s+)(.*)(\s+)=`),
		Type:    TYPE_KEYWORD_CONTROL,
	},
	// compare
	{
		Pattern: regexp.MustCompile(`companion(\s+)object`),
		Type:    TYPE_KEYWORD,
	},
	// when
	{
		Pattern: regexp.MustCompile(`when(\s+)(\((.*)\)\s+)?{$`),
		Type:    TYPE_KEYWORD_CONTROL,
	},
}
