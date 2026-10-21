# Tasks

## 1. Bump Gemara CUE schema pin

- [x] 1.1 Update `GEMARA_CUE_MODULE` default in `tests/schema-validation/validate.sh` from `v0.23.0` to `v1.6.0` and revise the comment on lines 27-29 to reflect alignment with `go-gemara v0.11.0`. Verify: `grep 'v1.6.0' tests/schema-validation/validate.sh` returns the updated line.

## 2. Add Renovate custom manager

- [x] 2.1 Create `renovate.json` in the repo root with: `$schema` reference, `enabledManagers: ["gomod", "custom.gemara-cue"]`, and a `customManagers` entry that matches the `GEMARA_CUE_MODULE` version in `validate.sh` using `github-releases` datasource for `gemaraproj/gemara`. Include a comment explaining the `enabledManagers` override. Verify: `cat renovate.json | python3 -m json.tool` validates as well-formed JSON.

## 3. Verification

- [x] 3.1 Run `make test-schema-validation` against existing scan output (or confirm the script parses the updated module reference correctly with `bash -n tests/schema-validation/validate.sh`). Verify: script syntax check passes with exit code 0.
- [x] 3.2 Run `openspec validate --change fix-gemara-cue-schema-pin` and confirm no validation errors. Verify: exit code 0.

## Workflow follow-up

- Archive the change after the project's review requirements are satisfied.

<!-- spec-review: passed -->
<!-- code-review: passed -->
