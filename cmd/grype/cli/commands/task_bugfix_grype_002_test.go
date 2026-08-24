package commands

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixGrype002SourceContract(t *testing.T) {
    source, err := os.ReadFile("update.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "if err != nil {") {
        t.Fatalf("expected source contract is missing")
    }
}
