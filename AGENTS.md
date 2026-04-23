# AGENTS.md - Exercises for Programmers Book

## Quick Commands

- `task` - List tasks (from go/ directory)
- `task test` - Run tests (gotestsum --format=testname)
- `task lint` - Run golangci-lint
- `task test-taxcalc` - Test specific package

## Project Structure

- Go module in go/ subdirectory (not workspace root)
- Always run task commands from go/ directory

## Key Conventions

- gotestsum uses --format=testname (not testdox)
- "exercices" misspell excluded in .golangci.yml - intentional for book title

## Testing

- Use t.Parallel() per tparallel linter rules
- Prefer dependency injection (io.Reader/io.Writer) over globals for testability
