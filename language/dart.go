package language

import "regexp"

var DART_LANGUAGE_PATTERN = []LanguagePattern{
	// Variable declaration
	{
		Pattern: regexp.MustCompile(`^\s*(const|final|var|dynamic|late)?\s*(int|double|String|bool|List<[A-Za-z [\](),]+>|HashMap<[A-Za-z [\](),]+>|Iterator<[A-Za-z [\](),]+>|Set<[A-Za-z [\](),]+>)?(\?)?\s\w+(\s=\s.+)?(;|,)$`),
		Type:    TYPE_KEYWORD_VARIABLE,
	},
	{
		Pattern: regexp.MustCompile(`\bstdout.write\(.+\);`),
		Type:    TYPE_KEYWORD_PRINT,
	},
	{
		Pattern: regexp.MustCompile(`\bprint\(.+\);`),
		Type:    TYPE_KEYWORD_PRINT,
	},
	{
		Pattern:   regexp.MustCompile(`^\s*import\s("|')dart:\w+("|')`),
		Type:      TYPE_META_IMPORT,
		IsNearTop: true,
	},
	{
		Pattern:   regexp.MustCompile(`^\s*import\s("|')package:\w+("|')`),
		Type:      TYPE_META_IMPORT,
		IsNearTop: true,
	},
	{
		Pattern:   regexp.MustCompile(`^\s*library\s\w+;`),
		Type:      TYPE_META_MODULE,
		IsNearTop: true,
	},
	{
		Pattern: regexp.MustCompile(`^\s*(void\s)?main\(\)\s(async\s)?{`),
		Type:    TYPE_KEYWORD_FUNCTION,
	},
	// function with type definition
	{
		Pattern: regexp.MustCompile(`^\s*(List<[A-Za-z [\](),]+>|HashMap<[A-Za-z [\](),]+>|int|double|String|bool|void|Iterator<[A-Za-z [\](),]+>|Set<[A-Za-z [\](),]+>)\s\w+\(.+\)\s*\{$`),
		Type:    TYPE_KEYWORD_FUNCTION,
	},
	// arrow function
	{
		Pattern: regexp.MustCompile(`^\s*(int|double|String|bool|List<[A-Za-z [\](),]+>|HashMap<[A-Za-z [\](),]+>|Iterator<[A-Za-z [\](),]+>|Set<[A-Za-z [\](),]+>)\s\w+\(.+\)\s=>`),
		Type:    TYPE_KEYWORD_FUNCTION,
	},
	{
		Pattern: regexp.MustCompile(`\bnew\s(List|Map|Iterator|HashMap|Set)<\w+>\(\);$`),
		Type:    TYPE_KEYWORD_VARIABLE,
	},
	{
		Pattern: regexp.MustCompile(`^(abstract\s)?class\s\w+\s(extends\s\w+\s)?(with\s\w+\s)?(implements\s\w+\s)?{(})?$`),
		Type:    TYPE_KEYWORD_CONTROL,
	},
	{
		Pattern: regexp.MustCompile(`\bget\s\w+=>\w+`),
		Type:    TYPE_KEYWORD_CONTROL,
	},
	{
		Pattern: regexp.MustCompile(`^\s*@override$`),
		Type:    TYPE_KEYWORD_CONTROL,
	},
	{
		Pattern: regexp.MustCompile(`\bset\s\w+\(.+\)`),
		Type:    TYPE_KEYWORD_CONTROL,
	},
	{
		Pattern: regexp.MustCompile(`^\s*Future<w+>\s\w+\(.+\)\sasync`),
		Type:    TYPE_KEYWORD_CONTROL,
	},
	{
		Pattern: regexp.MustCompile(`^\s*await\sfor`),
		Type:    TYPE_KEYWORD_CONTROL,
	},
	{
		Pattern: regexp.MustCompile(`^\s*typedef\s.+\s=`),
		Type:    TYPE_KEYWORD_CONTROL,
	},
	// Avoiding confusion with C
	{
		Pattern: regexp.MustCompile(`\blong\s`),
		Type:    TYPE_NOT,
	},
	{
		Pattern: regexp.MustCompile(`\s*function\b`),
		Type:    TYPE_NOT,
	},
	// Avoiding confusion with Java
	{
		Pattern: regexp.MustCompile(`\bArrayList`),
		Type:    TYPE_NOT,
	},
}
