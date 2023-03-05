package language

import "regexp"

var LUA_LANGUAGE_PATTERN = []LanguagePattern{
	// multiline string
	{
		Pattern: regexp.MustCompile(`(\[\[.*\]\])`),
		Type:    TYPE_CONSTANT_STRING,
	},
	// local definition
	{
		Pattern: regexp.MustCompile(`local\s([a-zA-Z0-9_]+)(\s*=)?`),
		Type:    TYPE_KEYWORD_VARIABLE,
	},
	// function definition
	{
		Pattern: regexp.MustCompile(`(local\s)?function\s*([a-zA-Z0-9_]*)?\(\)`),
		Type:    TYPE_KEYWORD_FUNCTION,
	},
	// for loop
	{
		Pattern: regexp.MustCompile(`for\s+([a-zA-Z]+)\s*=\s*([a-zA-Z0-9_]+),\s*([a-zA-Z0-9_]+)\s+do`),
		Type:    TYPE_KEYWORD_CONTROL,
	},
	// while loop
	{
		Pattern: regexp.MustCompile(`while\s(.*)\sdo`),
		Type:    TYPE_KEYWORD_CONTROL,
	},
	// keywords
	{
		Pattern: regexp.MustCompile(`\s+(and|break|do|else|elseif|end|false|function|if|in|not|or|local|repeat|return|then|true|until|pairs|ipairs|in|yield)`),
		Type:    TYPE_KEYWORD_OTHER,
	},
	{
		Pattern: regexp.MustCompile(`nil`),
		Type:    TYPE_CONSTANT_NULL,
	},
	// length operator
	{
		Pattern: regexp.MustCompile(`#([a-zA-Z_{}]+)`),
		Type:    TYPE_KEYWORD_OPERATOR,
	},
	// metatables
	{
		Pattern: regexp.MustCompile(`((get|set)metatable|raw(get|set|equal))\(.*\)`),
		Type:    TYPE_KEYWORD_OTHER,
	},
	// metamethods
	{
		Pattern: regexp.MustCompile(`__(index|newindex|call|sub|mul|div|mod|pow|unm|eq|le|lt)`),
		Type:    TYPE_KEYWORD_OTHER,
	},
	// method invocation
	{
		Pattern: regexp.MustCompile(`(\(.+\)|([a-zA-Z_]+)):([a-zA-Z_])\(.*\)`),
		Type:    TYPE_KEYWORD_OTHER,
	},
	// array-like table
	{
		Pattern: regexp.MustCompile(`{\s*[^\s;,]+([;,]\s*[^\s;,]+)*,?\s*}`),
		Type:    TYPE_CONSTANT_ARRAY,
	},
	// map-like table
	{
		Pattern: regexp.MustCompile(`{\s*([^\s;,=]+\s*=\s*[^\s;,=]+)(\s*[;,=]\s*[^\s;,=]+\s*=\s*[^\s;,=]+)*\s*,?\s*}`),
		Type:    TYPE_CONSTANT_DICTIONARY,
	},
	// builtin math methods
	{
		Pattern: regexp.MustCompile(`math\.(.*)\([0-9]*\)`),
		Type:    TYPE_MACRO,
	},
	// builtin table methods
	{
		Pattern: regexp.MustCompile(`table\.(.*)\(.*\)`),
		Type:    TYPE_MACRO,
	},
	// builtin io methods
	{
		Pattern: regexp.MustCompile(`io\.(.*)\(.*\)`),
		Type:    TYPE_MACRO,
	},
	// builtin functions
	{
		Pattern: regexp.MustCompile(`(require|dofile)\((.*)\)`),
		Type:    TYPE_META_IMPORT,
	},
	{
		Pattern: regexp.MustCompile(`(pcall|xpcall|unpack|pack|coroutine)`),
		Type:    TYPE_KEYWORD_OTHER,
	},
	// comments
	{
		Pattern: regexp.MustCompile(`--(\[\[)?.*`),
		Type:    TYPE_COMMENT_LINE,
	},
	// rest arguments
	{
		Pattern: regexp.MustCompile(`\.\.\.`),
		Type:    TYPE_KEYWORD_OTHER,
	},
	// module usage
	{
		Pattern: regexp.MustCompile(`\bmodule\s*\(.*\)`),
		Type:    TYPE_KEYWORD_OTHER,
	},

	// invalid comments
	{
		Pattern: regexp.MustCompile(`(\/\/|\/\*)`),
		Type:    TYPE_NOT,
	},
	// avoid confusion with C
	{
		Pattern: regexp.MustCompile(`(#(include|define)|printf|\s+int\s+)/`),
		Type:    TYPE_NOT,
	},
	// avoid confusion with javascript
	{
		Pattern: regexp.MustCompile(`\s+(let|const|var)\s+`),
		Type:    TYPE_NOT,
	},
	// avoid confusion with PHP & Python
	{
		Pattern: regexp.MustCompile(`\s+(echo|die|\$(.*))\s+`),
		Type:    TYPE_NOT,
	},
	// avoid confusion with Python
	{
		Pattern: regexp.MustCompile(`(def|len|from|import)`),
		Type:    TYPE_NOT,
	},
	// avoid confusion with SQL
	{
		Pattern: regexp.MustCompile(`(SELECT|FROM|INSERT|ALTER)`),
		Type:    TYPE_NOT,
	},
	// avoid confusion with Ruby
	{
		Pattern: regexp.MustCompile(`(puts)`),
		Type:    TYPE_NOT,
	},
	{
		Pattern: regexp.MustCompile(`\bmodule\s\S`),
		Type:    TYPE_NOT,
	},
	// avoid confusion Julia
	{
		Pattern: regexp.MustCompile(`(([a-zA-Z0-9]+)::([a-zA-Z0-9]+)|using|(.*)!\(.*\)|(\|\|))`),
		Type:    TYPE_NOT,
	},
}
