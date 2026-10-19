# Tasks

## 1. Create consolidated workflow

- [ ] 1.1 Create `.github/workflows/ci_local.yml` with
  `name: Local CI`, triggers `push: branches: [main]` and
  `pull_request:` (default types), `permissions: contents: read`,
  and four jobs: `buf-lint`, `unit-test`, `e2e-test`,
  `integration-test`. Each job preserves its existing steps
  from the source workflow. Verify the file passes
  `yamllint` and `actionlint` (if available).

## 2. Remove old workflows

- [ ] 2.1 Delete `.github/workflows/unit_test.yml`,
  `.github/workflows/e2e_test.yml`, and
  `.github/workflows/integration_test.yml`. Verify the files
  no longer exist and no other workflow references them.

## 3. Remove preflight override

- [ ] 3.1 Remove the `ci_checks` input override
  (`ci_checks: '["unit-test", "e2e-test", "integration-test"]'`)
  from the `preflight` job in `.github/workflows/release.yml`.
  Verify the `preflight` job `with:` block no longer contains
  `ci_checks`.

## 4. Update documentation

- [ ] 4.1 Update the CI Workflow Structure table in `AGENTS.md`:
  remove the Unit Test, E2E Test, and Integration Test rows;
  add a Local CI row for `ci_local.yml` with purpose
  describing the consolidated local test jobs. Verify the
  table renders correctly in markdown preview.

## 5. Verification

- [ ] 5.1 Run `make test-unit` locally to confirm tests still
  pass (no workflow logic affects test execution, but confirms
  nothing was inadvertently changed).
- [ ] 5.2 Inspect the consolidated `ci_local.yml` to confirm
  the preflight would auto-discover four check names:
  `Local CI / buf-lint`, `Local CI / unit-test`,
  `Local CI / e2e-test`, `Local CI / integration-test`.

## Workflow follow-up

- After PR merge, update GitHub branch protection rules to
  reference the new check names (e.g., `Local CI / unit-test`
  instead of `unit-test`) if any are configured.
- Archive the change after the project's review requirements
  are satisfied.
