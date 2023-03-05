package detection

import (
	"fmt"

	"github.com/mxgnus-de/go-code-detection/language"
)

func parseLanguagePoints(pointType language.Type) int {
	switch pointType {
	case language.TYPE_KEYWORD_PRINT,
		language.TYPE_META_IMPORT,
		language.TYPE_META_MODULE:
		return 5
	case language.TYPE_KEYWORD_FUNCTION,
		language.TYPE_CONSTANT_NULL:
		return 4
	case language.TYPE_CONSTANT_TYPE,
		language.TYPE_CONSTANT_STRING,
		language.TYPE_CONSTANT_NUMERIC,
		language.TYPE_CONSTANT_BOOLEAN,
		language.TYPE_CONSTANT_DICTIONARY,
		language.TYPE_CONSTANT_ARRAY,
		language.TYPE_KEYWORD_VARIABLE:
		return 3
	case language.TYPE_SECTION_SCOPE,
		language.TYPE_KEYWORD_OTHER,
		language.TYPE_KEYWORD_OPERATOR,
		language.TYPE_KEYWORD_CONTROL,
		language.TYPE_KEYWORD_VISIBILITY,
		language.TYPE_KEYWORD:
		return 2
	case language.TYPE_COMMENT_BLOCK,
		language.TYPE_COMMENT_LINE,
		language.TYPE_COMMENT_DOCUMENTATION,
		language.TYPE_MACRO:
		return 1
	case language.TYPE_NOT:
		return -20
	default:
		fmt.Println("Unknown type: ", pointType)
		return 0
	}
}

func getLanguagePoints(line string, patterns []language.LanguagePattern) int {
	points := []int{}

	for _, pattern := range patterns {
		if pattern.Pattern.MatchString(line) {
			points = append(points, parseLanguagePoints(pattern.Type))
		}
	}

	totalPoints := 0

	for _, point := range points {
		totalPoints += point
	}

	return totalPoints
}

func isNearTop(index int, lineCount int) bool {
	if lineCount <= 10 {
		return true
	}

	return index < lineCount/10
}
