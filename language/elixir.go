package language

import "regexp"

var ELIXIR_LANGUAGE_PATTERN = []LanguagePattern{
	{
		Pattern: regexp.MustCompile(`^\s*defmodule\s+.+\s+do$`),
		Type:    TYPE_META_MODULE,
	},
	// Alias
	{
		Pattern: regexp.MustCompile(`\s*alias\s+.+as:.+`),
		Type:    TYPE_KEYWORD_OTHER,
	},
	// IO.puts()
	{
		Pattern: regexp.MustCompile(`IO\.puts.+`),
		Type:    TYPE_KEYWORD_PRINT,
	},
	// Anonymous func
	{
		Pattern: regexp.MustCompile(`fn\s+[A-Za-z0-9_:<>()]+\s+->\s+.+(end)?$`),
		Type:    TYPE_KEYWORD_FUNCTION,
	},
	{
		Pattern: regexp.MustCompile(`^\s*(def|defp)\s+.+\s+do$`),
		Type:    TYPE_KEYWORD_FUNCTION,
	},
	{
		Pattern: regexp.MustCompile(`^\s*(if|unless|cond|case|try|defimpl|defprotocol)\s+.+\s+do$`),
		Type:    TYPE_KEYWORD_CONTROL,
	},
	{
		Pattern: regexp.MustCompile(`^\s*defstruct\s+`),
		Type:    TYPE_KEYWORD,
	},
	// Spec
	{
		Pattern: regexp.MustCompile(`^\s*@spec\s+.+::.+`),
		Type:    TYPE_MACRO,
	},
	// Lists
	{
		Pattern: regexp.MustCompile(`\{:.+,.+\}`),
		Type:    TYPE_CONSTANT_ARRAY,
	},
	// Maps
	{
		Pattern: regexp.MustCompile(`%\{(.+(=>|:).+(,)?){1,}\}`),
		Type:    TYPE_CONSTANT_DICTIONARY,
	},
}
