package language

import "regexp"

var PYTHON_LANGUAGE_PATTERN = []LanguagePattern{
	// Function definition
	{
		Pattern: regexp.MustCompile(`def\s+\w+\(.*\)\s*:`),
		Type:    TYPE_KEYWORD_FUNCTION,
	},
	// while loop
	{
		Pattern: regexp.MustCompile(`while (.+):`),
		Type:    TYPE_KEYWORD_CONTROL,
	},
	// from library import something
	{
		Pattern: regexp.MustCompile(`from [\w.]+ import (\w+|\*)`),
		Type:    TYPE_META_IMPORT,
	},
	// class keyword
	{
		Pattern: regexp.MustCompile(`class\s*\w+(\(\s*\w+\s*\))?\s*:`),
		Type:    TYPE_KEYWORD,
	},
	// if keyword
	{
		Pattern: regexp.MustCompile(`if\s+(.+)\s*:`),
		Type:    TYPE_KEYWORD_CONTROL,
	},
	// elif keyword
	{
		Pattern: regexp.MustCompile(`elif\s+(.+)\s*:`),
		Type:    TYPE_KEYWORD_CONTROL,
	},
	// else keyword
	{
		Pattern: regexp.MustCompile(`else:`),
		Type:    TYPE_KEYWORD_CONTROL,
	},
	// for loop
	{
		Pattern: regexp.MustCompile(`for (\w+|\(?\w+,\s*\w+\)?) in (.+):`),
		Type:    TYPE_KEYWORD_CONTROL,
	},
	// Python variable declaration.
	{
		Pattern: regexp.MustCompile(`\w+\s*=\s*\w+[^;]*(\n|$)`),
		Type:    TYPE_KEYWORD,
	},
	// import something
	{
		Pattern:   regexp.MustCompile(`import ([^\.]\w)+`),
		Type:      TYPE_META_IMPORT,
		IsNearTop: true,
	},
	// print statement/function
	{
		Pattern: regexp.MustCompile(`print((\s*\(.+\))|\s+.+)`),
		Type:    TYPE_KEYWORD_PRINT,
	},
	// &&/|| operators
	{
		Pattern: regexp.MustCompile(`(&{2}|\|{2})`),
		Type:    TYPE_NOT,
	},
	// avoiding lua
	{
		Pattern: regexp.MustCompile(`elseif`),
		Type:    TYPE_NOT,
	},
	{
		Pattern: regexp.MustCompile(`local\s(function|\w+)?\s=\s`),
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
	// Avoiding Dart confusion with the Future<any> keyword
	{
		Pattern: regexp.MustCompile(`Future<(.*)>`),
		Type:    TYPE_NOT,
	},
}
