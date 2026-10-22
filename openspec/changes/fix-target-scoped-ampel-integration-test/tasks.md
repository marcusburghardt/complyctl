# Tasks

## 1. Target-Scoped Bundle Resolution

- [x] 1.1 Confirm the `generate-index.json` schema from complytime-providers#207 and update `tests/cross-repo/cross_repo_integration_test.sh` to resolve the `complyctl-repo` bundle through a present index; reject absent target entries and paths outside the generated-artifact root.
- [x] 1.2 Preserve the legacy shared-path assertion only when no index exists; use the resolved bundle with the existing non-empty and `block-force-push` content assertions, and verify `shellcheck tests/cross-repo/cross_repo_integration_test.sh` passes.

## 2. Cross-Repository Regression Coverage

- [x] 2.1 Run `make test-cross-repo PROVIDERS_BIN_DIR=/path/to/providers/bin` against current provider main and the #207 provider build; verify both runs retain the existing Generate and Scan checks and pass.
- [x] 2.2 Exercise an index fixture or provider case with a missing target entry and a missing or empty indexed bundle; verify the generation assertion fails without falling back to the legacy path.

<!-- spec-review: passed -->
<!-- code-review: passed -->
