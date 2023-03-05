package language

import "regexp"

var C_LANGUAGE_PATTERN = []LanguagePattern{
	// Primitive variable declaration.
	{
		Pattern: regexp.MustCompile(`(char|long|int|float|double)\s+\w+\s*=?`),
		Type:    TYPE_CONSTANT_TYPE,
	},
	// malloc function call
	{
		Pattern: regexp.MustCompile(`malloc\(.+\)`),
		Type:    TYPE_KEYWORD_FUNCTION,
	},
	// #include <whatever.h>
	{
		Pattern: regexp.MustCompile(`#include (<|")\w+\.h(>|")`),
		Type:    TYPE_META_IMPORT,
	},
	// pointer
	{
		Pattern: regexp.MustCompile(`(\w+)\s*\*\s*\w+`),
		Type:    TYPE_KEYWORD,
	},
	// Variable declaration and/or initialisation.
	{
		Pattern: regexp.MustCompile(`(\w+)\s+\w+(;|\s*=)`),
		Type:    TYPE_MACRO,
	},
	// Array declaration.
	{
		Pattern: regexp.MustCompile(`(\w+)\s+\w+\[.+\]`),
		Type:    TYPE_KEYWORD_OTHER,
	},
	// #define macro
	{
		Pattern: regexp.MustCompile(`#define\s+.+`),
		Type:    TYPE_MACRO,
	},
	// NULL constant
	{
		Pattern: regexp.MustCompile(`NULL`),
		Type:    TYPE_CONSTANT_NULL,
	},
	// void keyword
	{
		Pattern: regexp.MustCompile(`void`),
		Type:    TYPE_KEYWORD_OTHER,
	},
	// (else )if statement
	// { pattern: /(else )?if\s*\(.+\)\s\{/, points: 1 },
	// while loop
	// { pattern: /while\s\(.+\)\s\{/, points: 1 },
	// printf function
	{
		Pattern: regexp.MustCompile(`(printf|puts)\s*\(.+\)`),
		Type:    TYPE_KEYWORD_PRINT,
	},
	// new Keyword from C++
	{
		Pattern: regexp.MustCompile(`new \w+`),
		Type:    TYPE_NOT,
	},
	// new Keyword from Java
	{
		Pattern: regexp.MustCompile(`new [A-Z]\w*\s*\(.+\)`),
		Type:    TYPE_NOT,
	},
	// Single quote multicharacter string
	{
		Pattern: regexp.MustCompile(`'.{2,}'`),
		Type:    TYPE_NOT,
	},
	// JS variable declaration
	{
		Pattern: regexp.MustCompile(`var\s+\w+\s*=?`),
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
	// Avoiding C# confusion
	{
		Pattern: regexp.MustCompile(`Console\.(WriteLine|Write)(\s*)?\(`),
		Type:    TYPE_NOT,
	},
	{
		Pattern: regexp.MustCompile(`(using\s)?System(\..*)?(;)?`),
		Type:    TYPE_NOT,
	},
	{
		Pattern: regexp.MustCompile(`(public\s)?((partial|static|delegate)\s)?(class\s)`),
		Type:    TYPE_NOT,
	},
	{
		Pattern: regexp.MustCompile(`(public|private|protected|internal)`),
		Type:    TYPE_NOT,
	},
	{
		Pattern: regexp.MustCompile(`(new|this\s)?(List|IEnumerable)<(sbyte|byte|short|ushort|int|uint|long|ulong|float|double|decimal|bool|char|string)>`),
		Type:    TYPE_NOT,
	},
	// Avoiding Lua confusion
	{
		Pattern: regexp.MustCompile(`local\s(function|\w+)?`),
		Type:    TYPE_NOT,
	},
	// Avoiding Dart confusion
	{
		Pattern: regexp.MustCompile(`^(void\s)?main\(\)\s(async\s)?{`),
		Type:    TYPE_NOT,
	},
}
