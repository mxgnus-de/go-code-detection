package detection

import "github.com/mxgnus-de/go-code-detection/language"

type LanguagePoints struct {
	Language language.Language
	Points   int
}

type detectionOptions struct {
	noUnknown       bool
	heuristic       bool
	lowerCaseOutput bool
}

type DetectionResult struct {
	Language    language.Language
	Statistics  map[language.Language]float64
	LinesOfCode int
}

type LanguageChecker struct {
	Language language.Language
	Patterns []language.LanguagePattern
}
