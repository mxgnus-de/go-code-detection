package language

import (
	"regexp"
	"strings"
)

type Type string
type Language string

const (
	TYPE_COMMENT_LINE          Type = "comment.line"
	TYPE_COMMENT_DOCUMENTATION Type = "comment.documentation"
	TYPE_COMMENT_BLOCK         Type = "comment.block"
	TYPE_META_IMPORT           Type = "meta.import"
	TYPE_META_MODULE           Type = "meta.module"
	TYPE_SECTION_SCOPE         Type = "section.scope"
	TYPE_CONSTANT_TYPE         Type = "constant.type"
	TYPE_CONSTANT_STRING       Type = "constant.string"
	TYPE_CONSTANT_NUMERIC      Type = "constant.numeric"
	TYPE_CONSTANT_BOOLEAN      Type = "constant.boolean"
	TYPE_CONSTANT_DICTIONARY   Type = "constant.dictionary"
	TYPE_CONSTANT_ARRAY        Type = "constant.array"
	TYPE_CONSTANT_NULL         Type = "constant.null"
	TYPE_KEYWORD               Type = "keyword"
	TYPE_KEYWORD_PRINT         Type = "keyword.print"
	TYPE_KEYWORD_VARIABLE      Type = "keyword.variable"
	TYPE_KEYWORD_CONTROL       Type = "keyword.control"
	TYPE_KEYWORD_VISIBILITY    Type = "keyword.visibility"
	TYPE_KEYWORD_OTHER         Type = "keyword.other"
	TYPE_KEYWORD_OPERATOR      Type = "keyword.operator"
	TYPE_KEYWORD_FUNCTION      Type = "keyword.function"
	TYPE_MACRO                 Type = "macro"
	TYPE_NOT                   Type = "not"
)

const (
	LANGUAGE_UNKNOWN    Language = "Unknown"
	LANGUAGE_C          Language = "C"
	LANGUAGE_CLOJURE    Language = "Clojure"
	LANGUAGE_CPP        Language = "C++"
	LANGUAGE_CS         Language = "C#"
	LANGUAGE_CSS        Language = "CSS"
	LANGUAGE_DART       Language = "Dart"
	LANGUAGE_DOCKERFILE Language = "Dockerfile"
	LANGUAGE_ELIXIR     Language = "Elixir"
	LANGUAGE_GO         Language = "Go"
	LANGUAGE_HTML       Language = "HTML"
	LANGUAGE_JAVA       Language = "Java"
	LANGUAGE_JAVASCRIPT Language = "JavaScript"
	LANGUAGE_JSON       Language = "JSON"
	LANGUAGE_JULIA      Language = "Julia"
	LANGUAGE_KOTLIN     Language = "Kotlin"
	LANGUAGE_LUA        Language = "Lua"
	LANGUAGE_MARKDOWN   Language = "Markdown"
	LANGUAGE_PASCAL     Language = "Pascal"
	LANGUAGE_PHP        Language = "PHP"
	LANGUAGE_PYTHON     Language = "Python"
	LANGUAGE_R          Language = "R"
	LANGUAGE_RUBY       Language = "Ruby"
	LANGUAGE_RUST       Language = "Rust"
	LANGUAGE_SQL        Language = "SQL"
	LANGUAGE_YAML       Language = "YAML"
	LANGUAGE_PERL       Language = "Perl"
	LANGUAGE_BASH       Language = "Bash"
)

type LanguagePattern struct {
	Pattern   *regexp.Regexp
	Type      Type
	IsNearTop bool
}

var LANGUAGES = map[Language][]LanguagePattern{
	LANGUAGE_C:          C_LANGUAGE_PATTERN,
	LANGUAGE_CLOJURE:    CLOJURE_LANGUAGE_PATTERN,
	LANGUAGE_CPP:        CPP_LANGUAGE_PATTERN,
	LANGUAGE_CS:         CS_LANGUAGE_PATTERN,
	LANGUAGE_CSS:        CSS_LANGUAGE_PATTERN,
	LANGUAGE_DART:       DART_LANGUAGE_PATTERN,
	LANGUAGE_DOCKERFILE: DOCKERFILE_LANGUAGE_PATTERN,
	LANGUAGE_ELIXIR:     ELIXIR_LANGUAGE_PATTERN,
	LANGUAGE_GO:         GO_LANGUAGE_PATTERN,
	LANGUAGE_HTML:       HTML_LANGUAGE_PATTERN,
	LANGUAGE_JAVA:       JAVA_LANGUAGE_PATTERN,
	LANGUAGE_JAVASCRIPT: JAVASCRIPT_LANGUAGE_PATTERN,
	LANGUAGE_JSON:       JSON_LANGUAGE_PATTERN,
	LANGUAGE_JULIA:      JULIA_LANGUAGE_PATTERN,
	LANGUAGE_KOTLIN:     KOTLIN_LANGUAGE_PATTERN,
	LANGUAGE_LUA:        LUA_LANGUAGE_PATTERN,
	LANGUAGE_MARKDOWN:   MARKDOWN_LANGUAGE_PATTERN,
	LANGUAGE_PASCAL:     PASCAL_LANGUAGE_PATTERN,
	LANGUAGE_PHP:        PHP_LANGUAGE_PATTERN,
	LANGUAGE_PYTHON:     PYTHON_LANGUAGE_PATTERN,
	LANGUAGE_RUBY:       RUBY_LANGUAGE_PATTERN,
	LANGUAGE_RUST:       RUST_LANGUAGE_PATTERN,
	LANGUAGE_SQL:        SQL_LANGUAGE_PATTERN,
	LANGUAGE_YAML:       YAML_LANGUAGE_PATTERN,
}

var SPECIAL_LOWERCASE_LANGUAGES = map[Language]Language{
	LANGUAGE_CPP: "cpp",
	LANGUAGE_CS:  "csharp",
}

func ConvertLowerCaseLangIntoLang(lang Language) Language {
	var found Language

	for k, v := range SPECIAL_LOWERCASE_LANGUAGES {
		if v == lang {
			found = k
			break
		}
	}

	if found != "" {
		return found
	}

	return Language(strings.ToUpper(string(lang[:1])) + string(lang[1:]))
}
