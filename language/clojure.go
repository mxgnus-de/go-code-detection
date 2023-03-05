package language

import "regexp"

var CLOJURE_LANGUAGE_PATTERN = []LanguagePattern{
	{
		Pattern: regexp.MustCompile(`^(\s+)?\(ns(\s+)(.*)(\))?`),
		Type:    TYPE_META_MODULE,
	},
	{
		Pattern: regexp.MustCompile(`^(\s+)?\(print(ln)?(\s+)(.*)(\))$`),
		Type:    TYPE_KEYWORD_PRINT,
	},
	{
		Pattern: regexp.MustCompile(`^(\s+)?\((de)?fn(-)?(\s+)(.*)(\))?$`),
		Type:    TYPE_KEYWORD_FUNCTION,
	},
	{
		Pattern: regexp.MustCompile(`^(\s+)?\((let|def)(\s+)(.*)(\))?$`),
		Type:    TYPE_KEYWORD_VARIABLE,
	},
	// Collection & sequences
	{
		Pattern: regexp.MustCompile(`^(\s+)?\((class|coll\?|seq\?|range|cons|conj|concat|map|filter|reduce)(\s+)(.*)(\))?`),
		Type:    TYPE_KEYWORD,
	},
	// Threading macro
	{
		Pattern: regexp.MustCompile(`^(\s+)?\((as)?->(>)?`),
		Type:    TYPE_MACRO,
	},
	// Modules
	{
		Pattern: regexp.MustCompile(`^(\s+)?\((use|require|import|:import)(\s+)(.*)(\))?`),
		Type:    TYPE_META_MODULE,
	},
	// Control keywords
	{
		Pattern: regexp.MustCompile(`^(\s+)?\((do|if|loop|cond|when|or|and|condp|case)`),
		Type:    TYPE_KEYWORD_CONTROL,
	},
}
