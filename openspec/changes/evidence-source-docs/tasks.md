# Tasks

## 1. Create docs/EVIDENCE.md

- [ ] 1.1 Create `docs/EVIDENCE.md` with overview section explaining evidence source provenance and its value for auditors, CI pipelines, and SREs. Verify: file exists and opens with a clear multi-audience motivation.
- [ ] 1.2 Add "Anatomy of an EvaluationLog" section with the real Ampel branch-protection example (one control: `force-push-restriction`, two evidence entries). Use the output from complytime-providers PR #192. Verify: YAML block parses without syntax errors.
- [ ] 1.3 Add "Tracing the Evidence Trail" section with five narrative steps: (1) control evaluation, (2) assessment evidence, (3) source provenance with coordinate/reference-id, (4) integrity and host identification with digest/remarks, (5) mapping-references cross-reference with ASCII diagram showing reference-id resolution. Verify: each step references concrete fields from the example in 1.2.
- [ ] 1.4 Add "Field Reference" table covering the five EvidenceSource fields (reference-id, coordinate, entry-id, digest, remarks) with purpose and example values from the real output. Verify: all five fields listed with non-empty purpose and example columns.
- [ ] 1.5 Add "Report Formats" section documenting field visibility per format: EvaluationLog (all five fields), Markdown (reference-id + coordinate only), OSCAL/SARIF (via gemara conversions). Verify: Markdown rendering example matches `formatEvidenceMeta()` output format (`source: <ref> @ <coord>`).
- [ ] 1.6 Add "Provider Support" section noting that evidence source requires provider support, and the `source:` block is omitted when absent. Keep provider version references generic. Verify: section does not reference specific provider versions.

## 2. Update docs/QUICK_START.md

- [ ] 2.1 Add "Understanding Scan Output" subsection in Step 7 after line 364 (`Output written to ./.complytime/scan/.`) and before the exit codes callout (line 366). Include one annotated evidence example, Markdown rendering line, provider support note, and link to `EVIDENCE.md`. Verify: subsection appears between scan output path and exit codes; link to EVIDENCE.md resolves correctly (`[Evidence Source Provenance](EVIDENCE.md)`).

## 3. Update docs/man/complyctl.md

- [ ] 3.1 Add OUTPUT section between EXIT CODES (line 452) and FILES (line 454). Include scan report directory, report filename pattern, and evidence source provenance field descriptions. Verify: section renders correctly as man page structure (# OUTPUT heading, ## subsections).
- [ ] 3.2 Add report filename pattern note to the `.complytime/scan/` FILES entry (line 478-479). Verify: filename example matches `BuildReportFilename()` output format in `internal/output/filename.go`.

## 4. Documentation gate

- [ ] 4.1 Add CHANGELOG.md entry under Unreleased for the documentation addition. Verify: entry references #876 and describes the three documentation deliverables.
- [ ] 4.2 Verify all internal links resolve: QUICK_START.md -> EVIDENCE.md, man page -> project repository EVIDENCE.md URL. Run: check that referenced files exist at the linked paths.

## Workflow follow-up

- Archive the change after the project's review requirements are satisfied.
