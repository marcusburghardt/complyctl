// SPDX-License-Identifier: Apache-2.0

package output_test

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/complytime/complyctl/internal/output"
)

func TestToSARIF_ProducesValidJSON(t *testing.T) {
	outDir := t.TempDir()
	log := mockGemaraEvalLog()

	path, err := output.ToSARIF(log, "file:///scan", outDir)
	require.NoError(t, err)
	assert.FileExists(t, path)

	data, err := os.ReadFile(path)
	require.NoError(t, err)

	var sarifDoc map[string]interface{}
	require.NoError(t, json.Unmarshal(data, &sarifDoc))

	assert.Contains(t, sarifDoc, "$schema")
	assert.Contains(t, sarifDoc, "version")

	runs, ok := sarifDoc["runs"].([]interface{})
	assert.True(t, ok)
	assert.NotEmpty(t, runs)
}

func TestToSARIF_OutputFileNaming(t *testing.T) {
	outDir := t.TempDir()
	log := mockGemaraEvalLog()

	path, err := output.ToSARIF(log, "file:///scan", outDir)
	require.NoError(t, err)

	filename := filepath.Base(path)
	assert.Contains(t, filename, "scan-test-policy-")
	assert.Contains(t, filename, "web-server")
	assert.Contains(t, filename, ".sarif.json")
}

func TestToSARIF_ExcludesPassed(t *testing.T) {
	outDir := t.TempDir()
	log := mockGemaraEvalLogWithFindings()

	path, err := output.ToSARIF(log, "file:///scan", outDir)
	require.NoError(t, err)

	data, err := os.ReadFile(path)
	require.NoError(t, err)

	var sarifDoc struct {
		Runs []struct {
			Results []struct {
				Kind       string         `json:"kind"`
				RuleID     string         `json:"ruleId"`
				Properties map[string]any `json:"properties"`
			} `json:"results"`
		} `json:"runs"`
	}
	require.NoError(t, json.Unmarshal(data, &sarifDoc))
	require.Len(t, sarifDoc.Runs, 1)

	results := sarifDoc.Runs[0].Results
	// Only the Failed evaluation should remain:
	// Passed is excluded via WithExcludedStatuses, NotApplicable is
	// always skipped by gemaraconv.ToSARIF.
	require.Len(t, results, 1)
	assert.Equal(t, "req-2", results[0].RuleID)
	assert.Equal(t, "fail", results[0].Kind)
	assert.NotContains(t, string(data), `"kind":"pass"`)
}
