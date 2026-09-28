# Tasks

## 1. Constants and Flag Type

- [ ] 1.1 Add `OutputFormatAll = "all"` constant to
  `internal/complytime/consts.go` alongside the existing
  `OutputFormat*` constants. Verify `make vet` passes.
- [ ] 1.2 Change `scanOptions.format` field from `string` to
  `[]string` in `cmd/complyctl/cli/scan.go`. Change flag
  registration from `StringVarP` to `StringSliceVarP`. Update
  flag help text to `"Output format(s): oscal, pretty, sarif,
  all (comma-separated or repeated)"`. Add `all` to shell
  completion values. Verify `make build` compiles without
  errors.

## 2. Validation

- [ ] 2.1 Rewrite `(o *scanOptions).validate()` to iterate over
  the `[]string` slice: reject unknown values, reject duplicates,
  and expand `all` to all concrete formats. Reject `all` combined
  with specific values. Verify with unit tests covering: single
  valid format, multiple valid formats, `all` alone, `all` with
  specific format (error), unknown format (error), duplicate
  format (error), empty slice (valid, no secondary reports).

## 3. Call Chain Signature Update

- [ ] 3.1 Update `runScanAndReport` signature to accept
  `formats []string` instead of `format string`. Thread the
  slice to `processScanOutput`. Verify `make build` compiles.
- [ ] 3.2 Update `processScanOutput` signature to accept
  `formats []string`. Thread to `writeScanReports`. Verify
  `make build` compiles.
- [ ] 3.3 Update `writeScanReports` signature to accept
  `formats []string`. Thread to `writeFormatReports`. Verify
  `make build` compiles.

## 4. Multi-Format Dispatch with Partial Success

- [ ] 4.1 Replace `writeFormatReport` (single format) with
  `writeFormatReports` (format slice). Loop over each format,
  call the existing per-format helpers (`writePrettyReport`,
  `writeSARIFReport`, `writeOSCALReport`). On error, collect
  a warning string and continue. After the loop, print any
  collected warnings to stderr. Return nil (never fail the
  scan due to secondary formatter errors). Verify with unit
  tests covering: multiple formats all succeed, one format
  fails with warning emitted, all formats fail with all
  warnings emitted.

## 5. Integration Verification

- [ ] 5.1 Run `make test-unit` and verify all existing tests
  pass with the refactored signatures.
- [ ] 5.2 Run `make lint` and verify zero lint issues.
- [ ] 5.3 Run `make build` and manually verify:
  `./bin/complyctl scan --help` shows the updated `--format`
  flag description.
