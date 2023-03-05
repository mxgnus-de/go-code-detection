# Language detector for Go

[Flourite language detector](https://github.com/teknologi-umum/flourite) implementation in Go

It detects a programming language from the supported languages from a given string.

-  No external dependencies
-  Over 250 test cases

### Detectable languages

| Languages |            |            |        |      |
| --------- | ---------- | ---------- | ------ | ---- |
| C         | Dockerfile | Javascript | Pascal | SQL  |
| C++       | Elixir     | Julia      | PHP    | YAML |
| C#        | Go         | Kotlin     | Python |      |
| Clojure   | HTML       | Lua        | Ruby   |      |
| CSS       | Java       | Markdown   | Rust   |      |

## Installation

```bash
$ go get -u github.com/mxgnus-de/go-code-detection
```

## Usage

```go
package main

import (
   "fmt"
   "github.com/mxgnus-de/go-code-detection/detection"
)

func main() {
   code := `
   package main

   import "fmt"

   func main() {
      fmt.Println("Hello world!")
   }`
	result := detection.DetectLanguage(code, detection.NewDetectionOptions())
   fmt.Printf("Detected language: %s\n", result.Language)
   // => Detected language: Go
   fmt.Printf("Result: %v\n", result)
   /*
   => Result: {
         Go                // <- Detected language
         map[              // <- Detected language stats
            C:0
            C#:0
            C++:0
            CSS:0
            Clojure:0
            Dart:0
            Dockerfile:0
            Elixir:0
            Go:15
            HTML:0
            JSON:0
            Java:0
            JavaScript:-20
            Julia:0
            Kotlin:0
            Lua:-18
            Markdown:0
            PHP:0
            Pascal:0
            Python:5
            Ruby:0
            Rust:-20
            SQL:0
            Unknown:1
            YAML:0
         ]
         8                // <- Total lines of code
      }
   */
}
```

If you want to disallow the `Unknown` value:

```go
package main

import (
   "fmt"
   "github.com/mxgnus-de/go-code-detection/detection"
)

func main() {
   code := `System.out.println('Hello world!')`
	result := detection.DetectLanguage(code, detection.NewDetectionOptions().WithNoUnknown())
   // => Java
}
```

By default only the top 500 lines of an given input get checked. To disable this feature:

```go
package main

import (
   "fmt"
   "github.com/mxgnus-de/go-code-detection/detection"
)

func main() {
   code := `console.log('Hello world!')`
	result := detection.DetectLanguage(code, detection.NewDetectionOptions().WithoutHeuristic())
   // => Javascript
}
```

## Available Options

| Option    | Type   | Default | Description                                                                                      |
| --------- | ------ | ------- | ------------------------------------------------------------------------------------------------ |
| Heuristic | `bool` | `true`  | Checks for codes on the top of the given input. Only checks when the lines of code is above 500. |
| NoUnknown | `bool` | `false` | If `true`, will not output `Unknown` on detected and statistics result                           |

## Credits

Special thanks to [Teknologi Umum ](https://github.com/teknologi-umum) for providing the original package [Flourite](https://github.com/teknologi-umum/flourite)

## License

[MIT](./LICENSE)
