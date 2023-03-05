package language

import "regexp"

var CPP_LANGUAGE_PATTERN = []LanguagePattern{
	// Primitive variable declaration.
	{
		Pattern: regexp.MustCompile(`(char|long|int|float|double)\s+\w+\s*=?`),
		Type:    TYPE_CONSTANT_TYPE,
	},
	// #include <whatever.h>
	{
		Pattern: regexp.MustCompile(`#include\s*(<|")\w+(\.h)?(>|")`),
		Type:    TYPE_META_IMPORT,
	},
	// using namespace something
	{
		Pattern: regexp.MustCompile(`using\s+namespace\s+.+\s*;`),
		Type:    TYPE_KEYWORD,
	},
	// template declaration
	{
		Pattern: regexp.MustCompile(`template\s*<.*>`),
		Type:    TYPE_KEYWORD,
	},
	// std
	{
		Pattern: regexp.MustCompile(`std::\w+`),
		Type:    TYPE_KEYWORD_OTHER,
	},
	// cout/cin/endl
	{
		Pattern: regexp.MustCompile(`(cout|cin|endl)`),
		Type:    TYPE_KEYWORD_PRINT,
	},
	// Visibility specifiers
	{
		Pattern: regexp.MustCompile(`(public|protected|private):`),
		Type:    TYPE_KEYWORD_VISIBILITY,
	},
	// nullptr
	{
		Pattern: regexp.MustCompile(`nullptr`),
		Type:    TYPE_KEYWORD,
	},
	// new Keyword
	{
		Pattern: regexp.MustCompile(`new \w+(\(.*\))?`),
		Type:    TYPE_KEYWORD,
	},
	// #define macro
	{
		Pattern: regexp.MustCompile(`#define\s+.+`),
		Type:    TYPE_MACRO,
	},
	// template usage
	{
		Pattern: regexp.MustCompile(`\w+<\w+>`),
		Type:    TYPE_KEYWORD_OTHER,
	},
	// class keyword
	{
		Pattern: regexp.MustCompile(`class\s+\w+`),
		Type:    TYPE_KEYWORD,
	},
	// void keyword
	{
		Pattern: regexp.MustCompile(`void`),
		Type:    TYPE_KEYWORD,
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
	// Scope operator
	{
		Pattern: regexp.MustCompile(`\w*::\w+`),
		Type:    TYPE_MACRO,
	},
	// Single quote multicharacter string
	{
		Pattern: regexp.MustCompile(`'.{2,}'`),
		Type:    TYPE_NOT,
	},
	// Java List/ArrayList
	{
		Pattern: regexp.MustCompile(`(List<\w+>|ArrayList<\w*>\s*\(.*\))(\s+[\w]+|;)`),
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
		Pattern: regexp.MustCompile(`\bmodule\s\S`),
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
		Pattern: regexp.MustCompile(`static\s+\S+\s+Main\(.*\)`),
		Type:    TYPE_NOT,
	},
	{
		Pattern: regexp.MustCompile(`(public|private|protected|internal)\s`),
		Type:    TYPE_NOT,
	},
	// Avoiding Kotlin confusion
	{
		Pattern: regexp.MustCompile(`fun main\((.*)?\) {`),
		Type:    TYPE_NOT,
	},
	{
		Pattern: regexp.MustCompile(`(inline|private|public|protected|override|operator(\s+))?fun(\s+)([A-Za-z0-9_])(\s+)?\((.*)\)(\s+)({|=)`),
		Type:    TYPE_NOT,
	},
	{
		Pattern: regexp.MustCompile(`(const)?(\s+)?val(\s+)(.*)(:(\s)(.*)(\?)?)?(\s+)=(\s+)`),
		Type:    TYPE_NOT,
	},
	// Avoiding Dart confusion
	{
		Pattern: regexp.MustCompile(`^(void\s)?main\(\)\s(async\s)?{`),
		Type:    TYPE_NOT,
	},
	{
		Pattern: regexp.MustCompile(`typedef\s+\w+\s+\w+`),
		Type:    TYPE_NOT,
	},
}
