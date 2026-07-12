# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Overview

Sokki is a markup language for note-taking ("速記" = shorthand). It compiles `.sk` files into a three-column HTML layout (CSS grid). Lines are placed in the left column by default; a ` >` suffix moves a line to the middle column and ` >>` to the right column. Blank lines separate sections (grid rows). See `resource/test.sk` and `resource/result.html` for a sample input/output pair.

## Commands

The repository has no `go.mod` (it predates Go modules; the import path is `github.com/roronya/sokki`). To build or test locally, first run:

```shell
go mod init github.com/roronya/sokki
```

Do not commit the generated `go.mod` unless asked to.

- Run all tests: `go test ./...`
- Run one package's tests: `go test ./lexer`
- Run a single test: `go test ./parser -run TestParseDocument`
- Build/run: `go build` then `./sokki input.sk output.html` (or `go run . input.sk output.html`)

## Architecture

The code follows the interpreter structure from "Writing An Interpreter In Go" (Monkey language), adapted to a line-oriented markup language. The pipeline in `main.go` is:

```
source text → lexer → parser → AST → evaluator → HTML string
```

- **token**: token types. Only five exist: `STRING` (a line's text), `SHIFT` (` >`), `MORESHIFT` (` >>`), `NEWLINE`, `EOD` (end of document).
- **lexer**: line-oriented, not character-oriented. `NextToken` grabs everything up to the next `\n`, then checks for a trailing ` >`/` >>` suffix (emitting the `STRING` first, then the shift token on the next call). Input is trimmed and handled as `[]rune` — position arithmetic must count runes, not bytes, because input is typically Japanese text.
- **ast**: `Document` → `[]*Section` → each Section holds `Left`/`Middle`/`Right` slices of `*Paragraph`. A `Section` corresponds to a blank-line-separated block and carries an `Id` used as the grid row.
- **parser**: `ParseDocument` loops over sections; `parseSection` reads a paragraph, then looks at the following token to decide which column (`SHIFT` → Middle, `MORESHIFT` → Right, otherwise Left). Maintains the `curToken`/`peekToken` two-token window. Malformed input is skipped rather than reported (the `errors` field is currently unused).
- **evaluator**: `Eval` renders the AST into an HTML page (template with inline CSS grid in `evaluator.go`). Paragraphs of all sections are grouped by column, with `grid-row: Id+1` placing each section on its row.
- **object**: vestigial from the Monkey interpreter scaffold; not used in the pipeline.

Tests are table-driven and live alongside each package (`lexer`, `parser`, `evaluator`). Comments and commit messages are in Japanese; follow that convention.
