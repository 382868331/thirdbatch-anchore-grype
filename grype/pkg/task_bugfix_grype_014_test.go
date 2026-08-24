package pkg

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixGrype014SourceContract(t *testing.T) {
    source, err := os.ReadFile("provider.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "if err != nil {") {
        t.Fatalf("expected source contract is missing")
    }
}
