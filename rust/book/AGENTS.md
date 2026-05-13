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
| `chapter_01` | bin  | `saying_hello`   | --      | `saying_hello.rs` | `io` (path `../io`) | `assert_cmd` (ws) |
| `io`         | lib  | --               | `prompt`| none  | none | none |

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

## `io` — Shared Library

- **Exports**: `fn prompt(output: &mut impl Write, msg: &str)`.
- **Used by**: `chapter_01` (path dependency).

## Key Patterns

- **Workspace deps**: `assert_cmd` shared via `[workspace.dependencies]`
- **Binary naming**: file stem = binary name (auto-discovered)
- **Testing**: `assert_cmd::Command::cargo_bin()` in `tests/`
- **I/O abstraction**: `impl BufRead` + `impl Write` for testability
- **Terminal detection**: `std::io::stdin().is_terminal()` for pipe vs interactive
- **Zero-cost**: `name.as_str()` for `match` patterns (no allocation)

## Commands

```sh
cargo build -p chapter_XX            # build specific chapter
cargo run --bin saying_hello -p chapter_01  # run a binary
cargo nextest run -p chapter_01      # test specific chapter
cargo nextest run --workspace        # test everything
```
