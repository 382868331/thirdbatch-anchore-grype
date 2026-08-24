package version

import (
    "os"
    "strings"
    "testing"
)

func TestTaskDiagnosisGrype004SourceContract(t *testing.T) {
    source, err := os.ReadFile("set.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "if out[j] == nil {") {
        t.Fatalf("expected source contract is missing")
    }
    if strings.Contains(string(source), "if out[j] != nil {") {
        t.Fatalf("mutated source contract is still present")
    }
}
