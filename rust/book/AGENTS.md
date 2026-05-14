# book/ — Rust Workspace

## Overview

Workspace for exercises from "Exercises for Programmers" book, implemented in Rust.

## Workspace Structure

```
book/
├── Cargo.toml                    # workspace root, resolver "3"
├── io/                           # shared library crate
├── chapter_00/                   # template/minimal example
└── chapter_01/                   # first real exercise
```

## Members

| Member       | Type | Binaries         | Library | Tests | Runtime Deps | Dev Deps |
|-------------|------|------------------|---------|-------|-------------|----------|
| `chapter_00` | bin  | `hello`          | --      | `hello.rs` | none | `assert_cmd` (ws) |
| `chapter_01` | bin  | `saying_hello`, `counting_number_chars` | -- | `saying_hello.rs`, `counting_number_chars.rs` | `io` (path `../io`) | `assert_cmd` (ws) |
| `io`         | lib  | --               | `write`, `prompt`, `read_once` | none  | none | none |

## `chapter_00` — Template

- **Binary**: `hello` — prints `"hello"`. Template for new chapters.
- **Tests**: single assertion on stdout.

## `chapter_01` — Saying Hello

- **Binary**: `saying_hello` — reads name from stdin, prints greeting.
- **Architecture**: directory-based binary (`src/bin/saying_hello/main.rs` + `run.rs`).
- **I/O**: decoupled via `impl BufRead`/`impl Write` (DI for testability).
- **Logic**: prompt → read → echo (pipe only) → greet.
- **Greeting variants**: default, `"Parker"`, `"Mechenyi"`.
- **Tests**: 2 integration test functions, 3 paths via `assert_cmd`.

## `chapter_01` — Counting Number of Characters

- **Binary**: `counting_number_chars` — reads string, prints char count.
- **Architecture**: directory-based binary (`src/bin/counting_number_chars/main.rs` + `run.rs`).
- **Logic**: prompt → read → if empty, print message and `process::exit(2)`, else count.
- **Return type**: `run() -> io::Result<()>` — uses `?` for error propagation.
- **Tests**: 2 integration tests (normal input + empty input), via `assert_cmd`.

## `io` — Shared Library

- **Exports**: `fn write(msg: &str)`, `fn prompt(msg: &str) -> io::Result<String>`, `fn read_once() -> io::Result<String>`.
- **Design**: global I/O (`io::stdin()`/`io::stdout()`), not DI-based.
- **`prompt`**: writes `"{msg} "`, calls `read_once()`, returns `Err(InvalidInput)` on empty.
- **`read_once`**: reads line, trims, echoes on pipe, returns `Ok(String)`.
- **`write`**: fire-and-forget write to stdout (panics on error).
- **Used by**: `chapter_01` (path dependency).

## Key Patterns

- **Workspace deps**: `assert_cmd` shared via `[workspace.dependencies]`
- **Binary naming**: file stem = binary name (auto-discovered)
- **Testing**: `assert_cmd::Command::cargo_bin()` in `tests/`
- **I/O approaches**: two styles coexist — DI with `impl BufRead`/`impl Write` (`saying_hello`) vs global I/O + `io::Result` (`counting_number_chars`)
- **Terminal detection**: `std::io::stdin().is_terminal()` for pipe vs interactive
- **Error pattern**: match on `Err(e) if e.kind() == ErrorKind::InvalidInput` for recoverable errors
- **String against `&str`**: use `.as_str()` or `if res.is_empty()` instead of matching `String` against `""`

## Commands

```sh
cargo build -p chapter_XX            # build specific chapter
cargo run --bin saying_hello -p chapter_01  # run a binary
cargo run --bin counting_number_chars -p chapter_01  # run a binary
cargo nextest run -p chapter_01      # test specific chapter
cargo nextest run --workspace        # test everything
```
