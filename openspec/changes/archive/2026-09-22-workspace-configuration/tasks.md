# Implementation Tasks

## Phase 1: Core Workspace Resolution

### Task 1.1: Add workspace resolution function
- [x] Add `WorkspaceEnvVar` constant to `internal/complytime/consts.go`
- [x] Implement `ResolveWorkspaceDir(flagValue string) (string, error)` in `internal/complytime/workspace.go`
  - [x] Check precedence: flag > env var > cwd
  - [x] Expand `~/` prefix using existing `ExpandPath()`
  - [x] Convert to absolute path via `filepath.Abs()`
  - [x] Validate path exists and is directory
  - [x] Return error with clear message for invalid paths

### Task 1.2: Add config detection function
- [x] Implement `DetectConfigPath(baseDir string) (string, bool, error)` in `internal/complytime/workspace.go`
  - [x] Check `.complytime/complytime.yaml` first
  - [x] Fall back to `complytime.yaml` at root
  - [x] Return `isLegacy=true` for root location
  - [x] Return error if neither exists

### Task 1.3: Add deprecation warning
- [x] Implement `printDeprecationWarning()` helper
  - [x] Print to stderr
  - [x] Include migration command in warning text
  - [x] Call from `NewWorkspace()` when legacy location detected

### Task 1.4: Update Workspace struct and constructor
- [x] Add `baseDir string` field to `Workspace` struct
- [x] Change `NewWorkspace()` signature to `NewWorkspace(baseDir string) *Workspace`
- [x] Update constructor to call `DetectConfigPath(baseDir)`
- [x] Store both `baseDir` and `configPath` in struct
- [x] Add `BaseDir() string` method to expose workspace root

## Phase 2: CLI Integration

### Task 2.1: Update Common struct and flag binding
- [x] Add `Workspace string` field to `Common` struct in `cmd/complyctl/cli/options.go`
- [x] Update `BindFlags()` to register `--workspace` / `-w` flag
- [x] Add `ResolveWorkspace() (string, error)` method to `Common`

### Task 2.2: Update command files to use resolved workspace
- [x] Update `init.go`: resolve workspace and pass to `NewWorkspace(baseDir)`
- [x] Update `list.go`: resolve workspace and pass to `NewWorkspace(baseDir)`
- [x] Update `scan.go`: resolve workspace and pass to `NewWorkspace(baseDir)`
- [x] Update `generate.go`: resolve workspace and pass to `NewWorkspace(baseDir)` (if applicable)
- [x] Update any other commands using `NewWorkspace()`

### Task 2.3: Update scan output path construction
- [x] Update `processScanOutput()` in `scan.go` to use baseDir parameter
- [x] Change `outDir := filepath.Join(".", complytime.WorkspaceDir, complytime.ScanOutputDir)`
  to `outDir := filepath.Join(baseDir, complytime.WorkspaceDir, complytime.ScanOutputDir)`
- [x] Verify `writeScanReports()` receives correct output directory

## Phase 3: Log Writer Update

### Task 3.1: Modify lazyLogWriter
- [x] Add `baseDir string` field to `lazyLogWriter` struct in `root.go`
- [x] Add `SetWorkspace(baseDir string)` method
- [x] Update `Write()` to use `w.baseDir` for log path construction
- [x] Change `logDir := complytime.WorkspaceDir` to `logDir := filepath.Join(w.baseDir, complytime.WorkspaceDir)`

### Task 3.2: Set workspace in PersistentPreRun
- [x] Update `PersistentPreRun` in `New()` to resolve workspace
- [x] Call `lw.SetWorkspace(baseDir)` after resolution
- [x] Handle resolution error gracefully (fall back to ".")

## Phase 4: Testing

### Task 4.1: Unit tests for workspace resolution
- [x] `TestResolveWorkspaceDir_FlagPrecedence`
- [x] `TestResolveWorkspaceDir_EnvVarFallback`
- [x] `TestResolveWorkspaceDir_DefaultToCwd`
- [x] `TestResolveWorkspaceDir_TildeExpansion`
- [x] `TestResolveWorkspaceDir_RelativeToAbsolute`
- [x] `TestResolveWorkspaceDir_InvalidPath`
- [x] `TestResolveWorkspaceDir_NotDirectory`

### Task 4.2: Unit tests for config detection
- [x] `TestDetectConfigPath_NewLocation`
- [x] `TestDetectConfigPath_LegacyFallback`
- [x] `TestDetectConfigPath_BothExist`
- [x] `TestDetectConfigPath_NeitherExists`

### Task 4.3: Unit tests for Workspace constructor
- [x] `TestNewWorkspace_WithBaseDir`
- [x] `TestNewWorkspace_BaseDir`
- [x] `TestNewWorkspace_ConfigPath`
- [x] `TestNewWorkspace_DeprecationWarning` (capture stderr)

### Task 4.4: Integration tests
- [x] Add test for `--workspace` flag with scan command
- [x] Add test for `COMPLYTIME_WORKSPACE` env var
- [x] Add test for flag overriding env var
- [x] Add test for relative workspace path
- [x] Add test for tilde expansion
- [x] Add test for legacy config deprecation warning
- [x] Add test for new location preferred when both exist
- [x] Add test for scan output in correct directory
- [x] Add test for log file in correct directory
- [x] Add test for error on invalid workspace path

## Phase 5: Documentation

### Task 5.1: Update user documentation
- [x] Update README.md with `--workspace` flag examples
- [x] Add migration guide section to README.md
- [x] Update CHANGELOG.md with feature description and migration instructions

### Task 5.2: Update project documentation
- [x] Update AGENTS.md "Recent Changes" section
- [x] Add entry for workspace-configuration OpenSpec

### Task 5.3: Update command help text
- [x] Verify `--workspace` flag appears in `complyctl --help`
- [x] Add workspace env var mention to `scan` command help text
- [x] Update examples in command help text if needed

## Phase 6: Final Verification

### Task 6.1: Manual testing
- [x] Test `complyctl init` creates `.complytime/complytime.yaml`
- [x] Test `complyctl scan --workspace /path` from different directory
- [x] Test `COMPLYTIME_WORKSPACE=/path complyctl scan`
- [x] Test legacy config shows deprecation warning
- [x] Test both locations exist (new location used, no warning)
- [x] Test invalid workspace path shows error

### Task 6.2: CI verification
- [x] Run full test suite: `make test-unit`
- [x] Run integration tests: `make test-integration`
- [x] Run E2E tests: `make test-e2e`
- [x] Run linter: `make lint`
- [x] Verify CRAP scores: `make crapload-check`

## Notes

- All tests must use `t.TempDir()` for filesystem isolation
- All error messages must be clear and actionable
- All path construction must use constants from `consts.go`
- All code must follow Go conventions from `go.md` convention pack
- Tasks marked `[P]` can be executed in parallel if needed
