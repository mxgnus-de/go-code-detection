package language

import "regexp"

var PASCAL_LANGUAGE_PATTERN = []LanguagePattern{
	{
		Pattern:   regexp.MustCompile(`^program (.*);$`),
		Type:      TYPE_META_MODULE,
		IsNearTop: true,
	},
	{
		Pattern:   regexp.MustCompile(`(?i)var$`),
		Type:      TYPE_CONSTANT_TYPE,
		IsNearTop: true,
	},
	{
		Pattern:   regexp.MustCompile(`(?i)const$`),
		Type:      TYPE_CONSTANT_TYPE,
		IsNearTop: true,
	},
	{
		Pattern:   regexp.MustCompile(`(?i)type$`),
		Type:      TYPE_CONSTANT_TYPE,
		IsNearTop: true,
	},
	{
		Pattern: regexp.MustCompile(`(?i)(write|writeln)(\s+)?(\((.*)\))?;`),
		Type:    TYPE_KEYWORD_PRINT,
	},
	{
		Pattern: regexp.MustCompile(`(?i)^(\s*)?(function|procedure)(\s*)(.*)\((.*)\)(\s)?:(\s)?(.*);$`),
		Type:    TYPE_KEYWORD_FUNCTION,
	},
	{
		Pattern: regexp.MustCompile(`(?i)end(\.|;)`),
		Type:    TYPE_KEYWORD_CONTROL,
	},
	{
		Pattern: regexp.MustCompile(`(?i):(\s*)?(cardinal|shortint|smallint|word|extended|comp)(\s*);$`),
		Type:    TYPE_CONSTANT_TYPE,
	},
	{
		Pattern: regexp.MustCompile(`(?i)if(\s+)(.*)(\s+)then`),
		Type:    TYPE_KEYWORD_CONTROL,
	},
	{
		Pattern: regexp.MustCompile(`(?i)for(\s+)(.*):=(.*)(\s+)(downto|to)(\s+)(.*)(\s+)do`),
		Type:    TYPE_KEYWORD_CONTROL,
	},
	{
		Pattern: regexp.MustCompile(`(?i)with(\s+)(.*)(\s+)do`),
		Type:    TYPE_KEYWORD_CONTROL,
	},
	{
		Pattern: regexp.MustCompile(`repeat$`),
		Type:    TYPE_KEYWORD,
	},
	{
		Pattern: regexp.MustCompile(`begin$`),
		Type:    TYPE_KEYWORD,
	},
	{
		Pattern: regexp.MustCompile(`(?i)until(\s+)(.*);`),
		Type:    TYPE_KEYWORD_CONTROL,
	},
	{
		Pattern: regexp.MustCompile(`(?i)\w+(\s*)?:=(\s*)?.+;$`),
		Type:    TYPE_KEYWORD_VARIABLE,
	},
}
