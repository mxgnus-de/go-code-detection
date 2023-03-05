package language

import "regexp"

var JAVA_LANGUAGE_PATTERN = []LanguagePattern{
	// System.out.println() etc.
	{
		Pattern: regexp.MustCompile(`System\.(in|out)\.\w+`),
		Type:    TYPE_KEYWORD_PRINT,
	},
	// Class variable declarations
	{
		Pattern: regexp.MustCompile(`(private|protected|public)\s*\w+\s*\w+(\s*=\s*[\w])?`),
		Type:    TYPE_KEYWORD,
	},
	// Method
	{
		Pattern: regexp.MustCompile(`(private|protected|public)\s*\w+\s*[\w]+\(.+\)`),
		Type:    TYPE_KEYWORD,
	},
	// String class
	{
		Pattern: regexp.MustCompile(`(^|\s)(String)\s+[\w]+\s*=?`),
		Type:    TYPE_KEYWORD_OTHER,
	},
	// List/ArrayList
	{
		Pattern: regexp.MustCompile(`(List<\w+>|ArrayList<\w*>\s*\(.*\))(\s+[\w]+|;)`),
		Type:    TYPE_KEYWORD_VARIABLE,
	},
	// class keyword
	{
		Pattern: regexp.MustCompile(`(public\s*)?class\b.*?\{`),
		Type:    TYPE_KEYWORD,
	},
	// Array declaration.
	{
		Pattern: regexp.MustCompile(`(\w+)(\[\s*\])+\s+\w+`),
		Type:    TYPE_CONSTANT_ARRAY,
	},
	// final keyword
	{
		Pattern: regexp.MustCompile(`final\s*\w+`),
		Type:    TYPE_KEYWORD_OTHER,
	},
	// getter & setter
	{
		Pattern: regexp.MustCompile(`\w+\.(get|set)\(.+\)`),
		Type:    TYPE_KEYWORD_OTHER,
	},
	// new Keyword (Java)
	{
		Pattern: regexp.MustCompile(`new [A-Z]\w*\s*\(.+\)`),
		Type:    TYPE_KEYWORD_OTHER,
	},
	// C style variable declaration.
	{
		Pattern: regexp.MustCompile(`(^|\s)(char|long|int|float|double)\s+[\w]+\s*=?`),
		Type:    TYPE_CONSTANT_TYPE,
	},
	// extends/implements keywords
	{
		Pattern:   regexp.MustCompile(`(extends|implements)`),
		Type:      TYPE_META_MODULE,
		IsNearTop: true,
	},
	// null keyword
	{
		Pattern: regexp.MustCompile(`null`),
		Type:    TYPE_KEYWORD_OTHER,
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
	// void keyword
	{
		Pattern: regexp.MustCompile(`void`),
		Type:    TYPE_KEYWORD_OTHER,
	},
	// const
	{
		Pattern: regexp.MustCompile(`const\s*\w+`),
		Type:    TYPE_NOT,
	},
	// pointer
	{
		Pattern: regexp.MustCompile(`(\w+)\s*\*\s*\w+`),
		Type:    TYPE_NOT,
	},
	// Single quote multicharacter string
	{
		Pattern: regexp.MustCompile(`'.{2,}'`),
		Type:    TYPE_NOT,
	},
	// C style include
	{
		Pattern:   regexp.MustCompile(`#include\s*(<|")\w+(\.h)?(>|")`),
		Type:      TYPE_NOT,
		IsNearTop: true,
	},
	// Avoiding Ruby confusion
	{
		Pattern: regexp.MustCompile(`def\s+\w+\s*(\(.+\))?\s*\n`),
		Type:    TYPE_NOT,
	},
	// Avoiding C# confusion
	{
		Pattern: regexp.MustCompile(`\bnamespace\s.*(\s{)?`),
		Type:    TYPE_NOT,
	},
	{
		Pattern: regexp.MustCompile(`\[Attribute\]`),
		Type:    TYPE_NOT,
	},
	{
		Pattern: regexp.MustCompile(`Console\.(WriteLine|Write)(\s*)?\(`),
		Type:    TYPE_NOT,
	},
	{
		Pattern: regexp.MustCompile(`(#region(\s.*)?|#endregion\n)`),
		Type:    TYPE_NOT,
	},
	{
		Pattern: regexp.MustCompile(`using\sSystem(\..*)?(;)?`),
		Type:    TYPE_NOT,
	},
	// Avoiding Kotlin confusion
	{
		Pattern: regexp.MustCompile(`fun main\((.*)?\) {`),
		Type:    TYPE_NOT,
	},
	{
		Pattern: regexp.MustCompile(`(inline(\s+))?fun(\s+)([A-Za-z0-9_])(\s+)?\((.*)\)(\s+)({|=)`),
		Type:    TYPE_NOT,
	},
	{
		Pattern: regexp.MustCompile(`(const)?(\s+)?val(\s+)(.*)(:(\s)(.*)(\?)?)?(\s+)=(\s+)`),
		Type:    TYPE_NOT,
	},
	// Avoiding Dart confusion
	{
		Pattern: regexp.MustCompile(`^(void\s)?main\(\)\s{`),
		Type:    TYPE_NOT,
	},
}
