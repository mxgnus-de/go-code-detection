package detection

import (
	"strings"

	"github.com/mxgnus-de/go-code-detection/language"
)

var SHEBANG_MAP = map[string]language.Language{
	"bash":   language.LANGUAGE_BASH,
	"sh":     language.LANGUAGE_BASH,
	"python": language.LANGUAGE_PYTHON,
	"ruby":   language.LANGUAGE_RUBY,
	"perl":   language.LANGUAGE_PERL,
	"node":   language.LANGUAGE_JAVASCRIPT,
	"php":    language.LANGUAGE_PHP,
	"lua":    language.LANGUAGE_LUA,
	"g++":    language.LANGUAGE_CPP,
	"gcc":    language.LANGUAGE_C,
	"java":   language.LANGUAGE_JAVA,
}

func isShebangPresent(lines []string) bool {
	return strings.HasPrefix(lines[0], "#!")
}

func hasLanguageSpecificShebang(lines []string) language.Language {
	if !isShebangPresent(lines) {
		return ""
	}

	shebangLanguage := ""

	if strings.HasPrefix(lines[0], "#!/usr/bin/env") {
		shebangLanguage = strings.TrimPrefix(lines[0], "#!/usr/bin/env ")
	} else if strings.HasPrefix(lines[0], "#!/usr/bin/") {
		shebangLanguage = strings.TrimPrefix(lines[0], "#!/usr/bin/")
	} else if strings.HasPrefix(lines[0], "#!/bin/") {
		shebangLanguage = strings.TrimPrefix(lines[0], "#!/bin/")
	}

	if shebangLanguage == "" {
		return ""
	}

	if SHEBANG_MAP[shebangLanguage] != "" {
		return SHEBANG_MAP[shebangLanguage]
	}

	shebangLanguage = strings.ToUpper(shebangLanguage[0:1]) + strings.ToLower(shebangLanguage[1:])
	return language.Language(shebangLanguage)
}
