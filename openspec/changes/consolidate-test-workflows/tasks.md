# Tasks

## 1. Create consolidated workflow

- [x] 1.1 Create `.github/workflows/ci_local.yml` with
  `name: Local CI`, triggers `push: branches: [main]` and
  `pull_request:` (default types), `permissions: contents: read`,
  and four jobs: `buf-lint`, `unit-test`, `e2e-test`,
  `integration-test`. Each job preserves its existing steps
  from the source workflow (including pinned action SHA
  references). Verify the file passes `yamllint` and
  `actionlint` (if available).

## 2. Remove old workflows

- [x] 2.1 Delete `.github/workflows/unit_test.yml`,
  `.github/workflows/e2e_test.yml`, and
  `.github/workflows/integration_test.yml`. Verify the files
  no longer exist and no other workflow references them.

## 3. Update preflight override

- [x] 3.1 Update the `ci_checks` input override in the
  `preflight` job in `.github/workflows/release.yml` to
  `'["buf-lint", "unit-test", "e2e-test", "integration-test"]'`.
  Auto-discovery removal deferred: org-infra preflight
  (v0.8.0) constructs check names as `<workflow name> / <job>`
  but the Checks API returns bare job keys, causing exact-match
  failures.

## 4. Update documentation

- [x] 4.1 Update the CI Workflow Structure table in `AGENTS.md`:
  remove the Unit Test, E2E Test, and Integration Test rows;
  add a Local CI row for `ci_local.yml` with purpose
  describing the consolidated local test jobs. Verify the
  table renders correctly in markdown preview.

## 5. Verification

- [x] 5.1 Run `make test-unit` locally to confirm tests still
  pass (no workflow logic affects test execution, but confirms
  nothing was inadvertently changed).
- [x] 5.2 Run the preflight `yq` expression against
  `ci_local.yml` and verify it produces exactly four check
  names:
  ```bash
  yq '.name as $wf | .jobs | keys[] | $wf + " / " + .' \
    .github/workflows/ci_local.yml
  ```
  Expected output:
  ```
  Local CI / buf-lint
  Local CI / e2e-test
  Local CI / integration-test
  Local CI / unit-test
  ```

## Pre-merge prerequisite

- [ ] Verify current branch protection required status checks
  (Settings → Branches → main → require status checks).
  Document which checks are required and coordinate with
  a repo admin to update them immediately after merge.
  If required checks reference the old names, the admin
  MUST update branch protection within minutes of merge
  to avoid blocking other PRs.

## Workflow follow-up

- After PR merge, update GitHub branch protection rules to
  reference the new check names (e.g., `Local CI / unit-test`
  instead of `unit-test`). Transition sequence: temporarily
  remove required status checks → merge → wait for new
  workflow to run once → add new check names → re-enable
  requirements.
- Archive the change after the project's review requirements
  are satisfied.
<!-- spec-review: passed -->
<!-- code-review: passed -->
