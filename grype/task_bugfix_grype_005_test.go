package grype

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixGrype005SourceContract(t *testing.T) {
    source, err := os.ReadFile("vulnerability_matcher.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "return false") {
        t.Fatalf("expected source contract is missing")
    }
}
