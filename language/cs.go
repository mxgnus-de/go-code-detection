package language

import "regexp"

var CS_LANGUAGE_PATTERN = []LanguagePattern{
	{
		Pattern: regexp.MustCompile(`using\sSystem(\..*)?(;)?`),
		Type:    TYPE_META_IMPORT,
	},
	{
		Pattern: regexp.MustCompile(`Console\.(WriteLine|Write)(\s*)?\(`),
		Type:    TYPE_KEYWORD_PRINT,
	},
	{
		Pattern: regexp.MustCompile(`Console\.ReadLine\(\)`),
		Type:    TYPE_KEYWORD_OTHER,
	},
	{
		Pattern: regexp.MustCompile(`(public\s)?((partial|static|delegate)\s)?class\s`),
		Type:    TYPE_KEYWORD,
	},
	// Modifiers
	{
		Pattern: regexp.MustCompile(`(extern|override|sealed|readonly|virtual|volatile)`),
		Type:    TYPE_KEYWORD_OTHER,
	},
	{
		Pattern: regexp.MustCompile(`namespace\s(.*)(\.(.*))?(\s{)?`),
		Type:    TYPE_KEYWORD,
	},
	// Regions
	{
		Pattern: regexp.MustCompile(`(#region(\s.*)?|#endregion\n)`),
		Type:    "section.scope",
	},
	// Functions
	{
		Pattern: regexp.MustCompile(`(public|private|protected|internal)\s`),
		Type:    TYPE_KEYWORD_VISIBILITY,
	},
	// class keyword
	{
		Pattern: regexp.MustCompile(`\bclass\s+\w+`),
		Type:    TYPE_KEYWORD,
	},
	// (else )if statement
	{
		Pattern: regexp.MustCompile(`(else )?if\s*\(.+\)`),
		Type:    TYPE_KEYWORD_CONTROL,
	},
	// while loop
	{
		Pattern: regexp.MustCompile(`\bwhile\s+\(.+\)`),
		Type:    TYPE_KEYWORD_CONTROL,
	},
	// Variable declaration
	{
		Pattern: regexp.MustCompile(`(const\s)?(sbyte|byte|short|ushort|int|uint|long|ulong|float|double|decimal|bool|char|string)(\[\])?\s(.*)\s=\s(.*);`),
		Type:    TYPE_CONSTANT_TYPE,
	},
	// Lists
	{
		Pattern: regexp.MustCompile(`(new|this\s)?(List|IEnumerable)<(sbyte|byte|short|ushort|int|uint|long|ulong|float|double|decimal|bool|char|string)>`),
		Type:    TYPE_CONSTANT_DICTIONARY,
	},
	// Macro
	{
		Pattern: regexp.MustCompile(`#define\s(.*)`),
		Type:    TYPE_MACRO,
	},
	// Plus point if you're doing PascalCase
	{
		Pattern: regexp.MustCompile(`\s([A-Z]([A-Z0-9]*[a-z][a-z0-9]*[A-Z]|[a-z0-9]*[A-Z][A-Z0-9]*[a-z])[A-Za-z0-9]*)\s=`),
		Type:    TYPE_MACRO,
	},
	// Avoiding Java confusion
	{
		Pattern: regexp.MustCompile(`(extends|throws|@Attribute)`),
		Type:    TYPE_NOT,
	},
	{
		Pattern: regexp.MustCompile(`System\.(in|out)\.\w+`),
		Type:    TYPE_NOT,
	},
	// Avoiding Ruby confusion
	{
		Pattern: regexp.MustCompile(`\bmodule\s\S`),
		Type:    TYPE_NOT,
	},
	// Avoiding Dart confusion
	{
		Pattern: regexp.MustCompile(`^\s*import\s("|')dart:\w+("|')`),
		Type:    TYPE_NOT,
	},
}
