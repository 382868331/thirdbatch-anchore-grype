package installation

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixGrype018SourceContract(t *testing.T) {
    source, err := os.ReadFile("curator.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "if m == nil {") {
        t.Fatalf("expected source contract is missing")
    }
    if strings.Contains(string(source), "if false && m == nil {") {
        t.Fatalf("mutated source contract is still present")
    }
}
