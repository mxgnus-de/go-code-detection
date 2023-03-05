# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [12.2.0] - 2021-08-02

### Added

-  Ordered lists: add order value to token info.

### Fixed

-  Always suffix indented code block with a newline, #799.

## [12.1.0] - 2021-07-01

### Changed

-  Updated CM spec compatibility to 0.30.

## [12.0.6] - 2021-04-16

### Fixed

-  Newline in \`alt\` should be rendered, #775.

## [12.0.5] - 2021-04-15

### Fixed

-  HTML block tags with \`===\` inside are no longer incorrectly interpreted as headers, #772.
-  Fix table/list parsing ambiguity, #767.

## [12.0.4] - 2020-12-20

### Fixed

-  Fix crash introduced in \`12.0.3\` when processing strikethrough (\`~~\`) and similar plugins, #742.
-  Avoid fenced token mutation, #745.

## [12.0.3] - 2020-12-07

### Fixed

-  \`[](foo<bar)\` is no longer a valid link.
-  \`[](url 'xxx(')\` is no longer a valid link.
-  \`[](url xxx)\` is no longer a valid link.
-  Fix performance issues when parsing links (#732, #734), backticks, (#733, #736),
   emphases (#735), and autolinks (#737).
-  Allow newline in \`<? ... ?>\` in an inline context.
-  Allow \`<meta>\` html tag to appear in an inline context.

## [12.0.2] - 2020-10-23

### Fixed

-  Three pipes (\`|\n|\n|\`) are no longer rendered as a table with no columns, #724.

## [12.0.1] - 2020-10-19

### Fixed

-  Fix tables inside lists indented with tabs, #721.

## [12.0.0] - 2020-10-14

### Added

-  \`.gitattributes\`, force unix eol under windows, for development.

### Changed

-  Added 3rd argument to \`highlight(code, lang, attrs)\`, #626.
-  Rewrite tables according to latest GFM spec, #697.
-  Use \`rollup.js\` to browserify sources.
-  Drop \`bower.json\` (bower reached EOL).
-  Deps bump.
-  Tune \`specsplit.js\` options.
-  Drop \`Makefile\` in favour of npm scrips.

### Fixed

-  Fix mappings for table rows (amended fix made in 11.0.1), #705.
-  \`%25\` is no longer decoded in beautified urls, #720.

[12.2.0]: https://github.com/markdown-it/markdown-it/compare/12.1.0...12.2.0
[12.1.0]: https://github.com/markdown-it/markdown-it/compare/12.0.6...12.1.0
[12.0.6]: https://github.com/markdown-it/markdown-it/compare/12.0.5...12.0.6
[12.0.5]: https://github.com/markdown-it/markdown-it/compare/12.0.4...12.0.5
[12.0.4]: https://github.com/markdown-it/markdown-it/compare/12.0.3...12.0.4
[12.0.3]: https://github.com/markdown-it/markdown-it/compare/12.0.2...12.0.3
[12.0.2]: https://github.com/markdown-it/markdown-it/compare/12.0.1...12.0.2
[12.0.1]: https://github.com/markdown-it/markdown-it/compare/12.0.0...12.0.1
[12.0.0]: https://github.com/markdown-it/markdown-it/compare/11.0.1...12.0.0
