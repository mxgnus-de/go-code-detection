package language

import "regexp"

var SQL_LANGUAGE_PATTERN = []LanguagePattern{
	{
		Pattern:   regexp.MustCompile(`CREATE (TABLE|DATABASE)`),
		Type:      TYPE_KEYWORD,
		IsNearTop: true,
	},
	{
		Pattern:   regexp.MustCompile(`DROP (TABLE|DATABASE)`),
		Type:      TYPE_KEYWORD,
		IsNearTop: true,
	},
	{
		Pattern:   regexp.MustCompile(`SHOW DATABASES`),
		Type:      TYPE_KEYWORD,
		IsNearTop: true,
	},
	{
		Pattern: regexp.MustCompile(`INSERT INTO`),
		Type:    TYPE_KEYWORD,
	},
	{
		Pattern: regexp.MustCompile(`(SELECT|SELECT DISTINCT)\s`),
		Type:    TYPE_KEYWORD,
	},
	{
		Pattern: regexp.MustCompile(`INNER JOIN`),
		Type:    TYPE_KEYWORD,
	},
	{
		Pattern: regexp.MustCompile(`(GROUP|ORDER) BY`),
		Type:    TYPE_KEYWORD,
	},
	{
		Pattern: regexp.MustCompile(`(END;|COMMIT;)`),
		Type:    TYPE_KEYWORD,
	},

	{
		Pattern: regexp.MustCompile(`UPDATE\s+\w+\sSET`),
		Type:    TYPE_KEYWORD,
	},
	{
		Pattern: regexp.MustCompile(`VALUES+(\s+\(\w|\(\w)`),
		Type:    TYPE_KEYWORD,
	},
	// Comments
	{
		Pattern: regexp.MustCompile(`--\s\w`),
		Type:    TYPE_COMMENT_LINE,
	},
	// Data types
	{
		Pattern: regexp.MustCompile(`(VARCHAR|CHAR|BINARY|VARBINARY|BLOB|TEXT)\([0-9]+\)`),
		Type:    TYPE_CONSTANT_TYPE,
	},
	{
		Pattern: regexp.MustCompile(`(BIT|TINYINT|SMALLINT|MEDIUMINT|INT|INTEGER|BIGINT|DOUBLE)\([0-9]+\)`),
		Type:    TYPE_CONSTANT_TYPE,
	},
	{
		Pattern: regexp.MustCompile(`(TINYBLOB|TINYTEXT|MEDIUMTEXT|MEDIUMBLOB|LONGTEXT|LONGBLOB)`),
		Type:    TYPE_CONSTANT_TYPE,
	},
	{
		Pattern: regexp.MustCompile(`(BOOLEAN|BOOL|DATE|YEAR)`),
		Type:    TYPE_CONSTANT_TYPE,
	},
	// Math
	{
		Pattern: regexp.MustCompile(`(EXP|SUM|SQRT|MIN|MAX)`),
		Type:    TYPE_KEYWORD_OPERATOR,
	},
	// Avoiding Lua
	{
		Pattern: regexp.MustCompile(`local\s(function|\w+)?\s=\s`),
		Type:    TYPE_NOT,
	},
	{
		Pattern: regexp.MustCompile(`(require|dofile)\((.*)\)`),
		Type:    TYPE_NOT,
	},
}
