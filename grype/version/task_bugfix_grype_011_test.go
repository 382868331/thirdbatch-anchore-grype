package version

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixGrype011SourceContract(t *testing.T) {
    source, err := os.ReadFile("rpm_version.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "if other == nil {") {
        t.Fatalf("expected source contract is missing")
    }
    if strings.Contains(string(source), "if false && other == nil {") {
        t.Fatalf("mutated source contract is still present")
    }
}
