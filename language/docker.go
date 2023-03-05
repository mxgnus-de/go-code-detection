package language

import (
	"regexp"
	"strings"
)

var DOCKERFILE_KEYWORDS = []string{
	"ADD",
	"ARG",
	"AS",
	"CMD",
	"COPY",
	"CROSS_BUILD",
	"ENTRYPOINT",
	"ENV",
	"EXPOSE",
	"FROM",
	"HEALTHCHECK",
	"LABEL",
	"MAINTAINER",
	"ONBUILD",
	"RUN",
	"SHELL",
	"STOPSIGNAL",
	"USER",
	"VOLUME",
	"WORKDIR",
}

var DOCKERFILE_LANGUAGE_PATTERN = []LanguagePattern{
	// Keywords
	{
		Pattern: regexp.MustCompile(`(?i)(` + strings.Join(DOCKERFILE_KEYWORDS, "|") + `)\s`),
		Type:    TYPE_KEYWORD,
	},
	// Avoiding SQL confusion with "SELECT"
	{
		Pattern: regexp.MustCompile(`(?i)SELECT\s`),
		Type:    TYPE_NOT,
	},
}
