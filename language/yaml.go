package language

import "regexp"

var YAML_LANGUAGE_PATTERN = []LanguagePattern{
	// Regular key: value
	{
		Pattern: regexp.MustCompile(`^( )*([A-Za-z0-9_. ]+):( )?(.*)?$`),
		Type:    TYPE_KEYWORD,
	},
	// Regular array - key: value
	{
		Pattern: regexp.MustCompile(`^( )*-( )([A-Za-z0-9_. ]+):( )?(.*)?$`),
		Type:    TYPE_KEYWORD,
	},
	// Regular array - value
	{
		Pattern: regexp.MustCompile(`^( )*-( )(.*)$`),
		Type:    TYPE_KEYWORD,
	},
	// Binary tag
	{
		Pattern: regexp.MustCompile(`^( )*([A-Za-z0-9_. ]+):( )!!binary( )?(|)?$`),
		Type:    TYPE_CONSTANT_TYPE,
	},
	// Literal multiline block
	{
		Pattern: regexp.MustCompile(`^( )*([A-Za-z0-9_. ]+):( )\|$`),
		Type:    TYPE_KEYWORD,
	},
	// Folded multiline style
	{
		Pattern: regexp.MustCompile(`^( )*([A-Za-z0-9_. ]+):( )>$`),
		Type:    TYPE_KEYWORD,
	},
	// Set types
	{
		Pattern: regexp.MustCompile(`^( )*\?( )(.*)$`),
		Type:    TYPE_KEYWORD,
	},
	// Complex key / multiline key
	{
		Pattern: regexp.MustCompile(`^( )*\?( )\|$`),
		Type:    TYPE_CONSTANT_TYPE,
	},
	// Merge key
	{
		Pattern: regexp.MustCompile(`^( )*<<:( )(\*)(.*)?$`),
		Type:    TYPE_CONSTANT_TYPE,
	},
	// Avoiding confusion with CSS
	{
		Pattern: regexp.MustCompile(`^( )*([A-Za-z0-9_. ]+):(.*)?( )?{$`),
		Type:    TYPE_NOT,
	},
	{
		Pattern: regexp.MustCompile(`^( )*([A-Za-z0-9_. ]+):(.*)?( )?,$`),
		Type:    TYPE_NOT,
	},
}
