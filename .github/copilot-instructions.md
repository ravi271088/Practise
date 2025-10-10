<!-- .github/copilot-instructions.md - guidance for AI coding agents -->
# Quick agent instructions for this repo

Purpose: short, practical notes to help an AI agent be productive editing and extending this small Go practice project.

## Quick facts / run
- This is a tiny Go workspace with a single runnable program at `interface/main.go`.
- No `go.mod` file is present in the repo root. Typical local run commands (Windows PowerShell):

```powershell
# from repo root
go run .\interface\main.go
# or change into the folder then run
cd interface; go run .
```

## Project layout (relevant files)
- `interface/main.go` — the only program. Package `main` with global variables and method receivers.
- `README.md` — minimal project README (no build instructions).

## Important, discoverable code patterns
- The code declares an interface as a variable, not a named type:

```go
var shape interface {
    area() float64
    perimeter() float64
}
```

- Two anonymous struct variables are used as concrete values: `square` and `circle` (both declared as `var <name> struct { ... }`). Methods for those are defined with value receivers:

```go
var square struct { side float64 }
func (s square) area() float64 { return s.side * s.side }
```

- The program uses the built-in `println` for output rather than `fmt.Println` or `fmt.Printf`. Keep this in mind when adding logging or tests.
- The code uses a hard-coded `3.14` for π in the circle calculations.

## Editing guidance / conventions
- Keep changes minimal and mechanical unless asked to redesign. This is a small
project used for practice; avoid large refactors without confirmation.

## Small contract for edits
- Inputs: edits should preserve the program's visible behavior (stdout) unless the task explicitly asks for behavior change.
- Outputs: compiled `main` program that prints the square and circle area/perimeter.
- Error modes: there's no error handling in current code; adding IO or external dependencies must include small safety checks and documentation.

## Edge cases to consider when changing code
- The interface is declared as a variable and assigned concrete anonymous structs. If you convert the interface to a named type or the structs to named types, adjust method receivers and assignments accordingly.
- Methods use value receivers; if you change to pointer receivers, ensure assignments use pointers (e.g., `&square`).
- The program assumes non-negative dimensions; adding validation should be explicit and tested.

## Examples of safe edits an agent might perform
- Replace `println` with `fmt.Printf` for formatted output: update imports and calls in `interface/main.go`.
- Introduce `go.mod` to make builds reproducible: run `go mod init github.com/yourname/Practise` (ask user for module path) and `go run` will then work as a module.
- Replace hard-coded `3.14` with `math.Pi` from the standard library (remember to import `math`).

## Tests and verification
- There are no tests. Add small unit tests in `interface/main_test.go` when changing calculation logic.
- Quick smoke run (PowerShell):

```powershell
cd interface; go run .
```

## Integration and external dependencies
- Currently none. The repo is self-contained and uses only Go standard library primitives when needed.

## Files to reference when making changes
- `interface/main.go` — primary source to edit.
- `README.md` — update if you add build instructions or module information.

## What to ask the user before larger changes
- If adding a `go.mod` ask for the desired module path.
- If changing the public API (named types/method receivers), confirm expected external usage.

If anything in this summary looks incomplete or you want the agent to enforce different conventions (for example, always use named types or include unit tests), tell me which direction to take and I will update this file.
