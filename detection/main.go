package detection

import (
	"bufio"
	"math"
	"regexp"
	"sort"
	"strings"

	"github.com/mxgnus-de/go-code-detection/language"
)

func NewDetectionOptions() *detectionOptions {
	return &detectionOptions{
		noUnknown: false,
		heuristic: false,
	}
}

func (o *detectionOptions) WithNoUnknown() *detectionOptions {
	o.noUnknown = true
	return o
}

func (o *detectionOptions) WithoutHeuristic() *detectionOptions {
	o.heuristic = true
	return o
}

func DetectLanguage(code string, options *detectionOptions) DetectionResult {
	lines, lineCount := getLines(code)

	if options.heuristic && lineCount > 500 {
		newLines := []string{}

		for i, line := range lines {
			if isNearTop(i, lineCount) {
				newLines = append(newLines, line)
			}

			if math.Mod(float64(i), float64(lineCount/500)) == 0 {
				newLines = append(newLines, line)
			}
		}

		lines = newLines
		lineCount = len(lines)
	}

	result := DetectionResult{
		Language:    language.LANGUAGE_UNKNOWN,
		Statistics:  map[language.Language]float64{},
		LinesOfCode: lineCount,
	}

	// if the first line is a shebang, check if it is a language specific shebang
	if isShebangPresent(lines) {
		shebangLanguage := hasLanguageSpecificShebang(lines)

		if shebangLanguage != "" {
			result.Language = shebangLanguage
			return result
		}
	}

	checkers := []LanguageChecker{}

	for language, languagePattern := range language.LANGUAGES {
		checker := LanguageChecker{
			Language: language,
			Patterns: languagePattern,
		}

		checkers = append(checkers, checker)
	}

	emptyFileRegexp := regexp.MustCompile(`^\s*$`)
	languagePoints := []LanguagePoints{}

	for _, checker := range checkers {
		checkerPatterns := checker.Patterns
		points := 0

		for j, line := range lines {
			// early return if line is empty or only contains spaces
			if match := emptyFileRegexp.Match([]byte(line)); match {
				continue
			}

			if !isNearTop(j, lineCount) {
				patterns := []language.LanguagePattern{}

				for _, checker := range checkerPatterns {
					if !checker.IsNearTop {
						patterns = append(patterns, checker)
					}
				}

				points += getLanguagePoints(line, patterns)
			} else {
				points += getLanguagePoints(line, checker.Patterns)
			}
		}

		languagePoint := LanguagePoints{
			Language: checker.Language,
			Points:   points,
		}
		languagePoints = append(languagePoints, languagePoint)
	}

	sort.Slice(languagePoints, func(i, j int) bool {
		return languagePoints[i].Language < languagePoints[j].Language
	})

	if !options.noUnknown {
		unknownLanguagePoints := LanguagePoints{
			Language: "Unknown",
			Points:   1,
		}
		languagePoints = append(languagePoints, unknownLanguagePoints)
	}

	bestLanguagePointResult := getBestLanguagePointsResult(languagePoints)
	languagePointStatisticMap := getLanguagePointStatisticMap(languagePoints)

	result.Language = bestLanguagePointResult.Language
	result.Statistics = languagePointStatisticMap

	return result
}

func getLines(code string) ([]string, int) {
	scanner := bufio.NewScanner(strings.NewReader(code))
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, len(lines)
}

func getBestLanguagePointsResult(languagePoints []LanguagePoints) LanguagePoints {
	bestResult := LanguagePoints{}

	for _, languagePoint := range languagePoints {
		if languagePoint.Points > bestResult.Points {
			bestResult = languagePoint
		}
	}

	return bestResult
}

func getLanguagePointStatisticMap(languagePoints []LanguagePoints) map[language.Language]float64 {
	languagePointStatisticMap := map[language.Language]float64{}

	for _, languagePoint := range languagePoints {
		languagePointStatisticMap[languagePoint.Language] = float64(languagePoint.Points)
	}

	return languagePointStatisticMap
}
