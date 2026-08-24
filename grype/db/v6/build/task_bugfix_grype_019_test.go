package v6

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixGrype019SourceContract(t *testing.T) {
    source, err := os.ReadFile("govulndb_merge.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "if !ok || aph.Package == nil {") {
        t.Fatalf("expected source contract is missing")
    }
    if strings.Contains(string(source), "if !ok && aph.Package == nil {") {
        t.Fatalf("mutated source contract is still present")
    }
}
