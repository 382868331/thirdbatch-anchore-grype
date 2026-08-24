package internal

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixGrype017SourceContract(t *testing.T) {
    source, err := os.ReadFile("only_vulnerable_targets.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "return false, reason") {
        t.Fatalf("expected source contract is missing")
    }
    if strings.Contains(string(source), "return true, reason") {
        t.Fatalf("mutated source contract is still present")
    }
}
