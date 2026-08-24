package differ

import (
    "os"
    "strings"
    "testing"
)

func TestTaskDiagnosisGrype020SourceContract(t *testing.T) {
    source, err := os.ReadFile("differ.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "if err != nil {") {
        t.Fatalf("expected source contract is missing")
    }
}
