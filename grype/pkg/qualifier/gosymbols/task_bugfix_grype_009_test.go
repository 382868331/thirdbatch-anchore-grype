package gosymbols

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixGrype009SourceContract(t *testing.T) {
    source, err := os.ReadFile("qualifier.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "return true, nil") {
        t.Fatalf("expected source contract is missing")
    }
}
