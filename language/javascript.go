package language

import "regexp"

var JAVASCRIPT_LANGUAGE_PATTERN = []LanguagePattern{
	// undefined keyword
	{
		Pattern: regexp.MustCompile(`undefined`),
		Type:    TYPE_KEYWORD,
	},
	// window keyword
	{
		Pattern: regexp.MustCompile(`window\.`),
		Type:    TYPE_KEYWORD,
	},
	// console.log('ayy lmao')
	{
		Pattern: regexp.MustCompile(`console\.log\s*\(`),
		Type:    TYPE_KEYWORD_PRINT,
	},
	// Variable declaration
	{
		Pattern: regexp.MustCompile(`(var|const|let)\s+\w+\s*=?`),
		Type:    TYPE_KEYWORD_VARIABLE,
	},
	// Array/Object declaration
	{
		Pattern: regexp.MustCompile(`(('|\").+('|\")\s*|\w+):\s*[\[{]`),
		Type:    TYPE_CONSTANT_ARRAY,
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
	// Function definition
	{
		Pattern: regexp.MustCompile(`function\*?\s*([A-Za-z$_][\w$]*)?\s*[(][^:;()]*[)]\s*{`),
		Type:    TYPE_KEYWORD_FUNCTION,
	},
	// arrow function
	{
		Pattern: regexp.MustCompile(`\(* => {`),
		Type:    TYPE_KEYWORD_FUNCTION,
	},
	// null keyword
	{
		Pattern: regexp.MustCompile(`null`),
		Type:    TYPE_CONSTANT_NULL,
	},
	// lambda expression
	{
		Pattern: regexp.MustCompile(`\(.*\)\s*=>\s*.+`),
		Type:    TYPE_KEYWORD_CONTROL,
	},
	// (else )if statement
	{
		Pattern: regexp.MustCompile(`(else )?if\s+\(.+\)`),
		Type:    TYPE_KEYWORD_CONTROL,
	},
	// while loop
	{
		Pattern: regexp.MustCompile(`while\s+\(.+\)`),
		Type:    TYPE_KEYWORD_CONTROL,
	},
	// C style variable declaration.
	{
		Pattern: regexp.MustCompile(`(^|\s)(char|long|int|float|double)\s+\w+\s*=?`),
		Type:    TYPE_NOT,
	},
	// pointer
	{
		Pattern: regexp.MustCompile(`\*\w+`),
		Type:    TYPE_NOT,
	},
	// HTML <script> tag
	{
		Pattern: regexp.MustCompile(`<(\/)?script( type=('|\")text\/javascript('|\"))?\s*>`),
		Type:    TYPE_NOT,
	},
	{
		Pattern: regexp.MustCompile(`fn\s[A-Za-z0-9<>,]+\(.*\)\s->\s\w+(\s\{|)`),
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
		Pattern: regexp.MustCompile(`(func|fn)\s+`),
		Type:    TYPE_NOT,
	},
	{
		Pattern: regexp.MustCompile(`(begin|end)\n`),
		Type:    TYPE_NOT,
	},
	// Avoiding Lua confusion
	{
		Pattern: regexp.MustCompile(`local\s(function|(\w+)\s=)`),
		Type:    TYPE_NOT,
	},
	// Avoiding Kotlin confusion
	{
		Pattern: regexp.MustCompile(`fun main\((.*)?\)\s{`),
		Type:    TYPE_NOT,
	},
	{
		Pattern: regexp.MustCompile(`(inline(\s+))?fun(\s+)([A-Za-z0-9_])(\s+)?\((.*)\)(\s+)(\{|=)`),
		Type:    TYPE_NOT,
	},
	{
		Pattern: regexp.MustCompile(`(const)?(\s+)?val(\s+)(.*)(:(\s)(.*)(\?)?)?(\s+)=(\s+)`),
		Type:    TYPE_NOT,
	},
	// Avoiding Dart confusion
	{
		Pattern: regexp.MustCompile(`^(void\s)?main()\s{`),
		Type:    TYPE_NOT,
	},
}
