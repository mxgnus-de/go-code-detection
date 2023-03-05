package language

import (
	"regexp"
)

var GO_LANGUAGE_PATTERN = []LanguagePattern{
	// package [name]
	{
		Pattern:   regexp.MustCompile(`package\s+[a-z]+\n`),
		Type:      TYPE_META_IMPORT,
		IsNearTop: true,
	},
	// import
	{
		Pattern:   regexp.MustCompile(`(import\s*\(\s*\n)|(import\s+"[a-z0-9/.]+")`),
		Type:      TYPE_META_IMPORT,
		IsNearTop: true,
	},
	// error check
	{
		Pattern: regexp.MustCompile(`if\s+err\s*!=\s*nil\s*\{\s*\n`),
		Type:    TYPE_KEYWORD_FUNCTION,
	},
	// Go print
	{
		Pattern: regexp.MustCompile(`fmt\.Println\(`),
		Type:    TYPE_KEYWORD_PRINT,
	},
	// function
	{
		Pattern: regexp.MustCompile(`func(\s+\w+\s*)?\(.*\).*{`),
		Type:    TYPE_KEYWORD_FUNCTION,
	},
	// variable initialization
	{
		Pattern: regexp.MustCompile(`\w+\s*:=\s*.+[^;\n]`),
		Type:    TYPE_KEYWORD_VARIABLE,
	},
	// variable declaration (var/const)
	{
		Pattern: regexp.MustCompile(`(var|const)\s+\w+\s+[\w*]+(\n|\s*=|$)`),
		Type:    TYPE_KEYWORD_VARIABLE,
	},
	// if/else if
	{
		Pattern: regexp.MustCompile(`(}\s*else\s*)?if[^()]+{`),
		Type:    TYPE_KEYWORD_CONTROL,
	},
	// public access on package
	{
		Pattern: regexp.MustCompile(`[a-z]+\.[A-Z]\w*`),
		Type:    TYPE_MACRO,
	},
	// nil keyword
	{
		Pattern: regexp.MustCompile(`nil`),
		Type:    TYPE_CONSTANT_NULL,
	},
	// single quote multicharacter string
	{
		Pattern: regexp.MustCompile(`'.{2,}'`),
		Type:    TYPE_NOT,
	},
	// avoiding c# confusion
	{
		Pattern: regexp.MustCompile(`Console\.(WriteLine|Write)(\s*)?\(`),
		Type:    TYPE_NOT,
	},
	{
		Pattern: regexp.MustCompile(`using\sSystem(\..*)?(;)?`),
		Type:    TYPE_NOT,
	},
	{
		Pattern: regexp.MustCompile(`(public|private|protected|internal)\s`),
		Type:    TYPE_NOT,
	},
}
