package dbtest

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixGrype013SourceContract(t *testing.T) {
    source, err := os.ReadFile("workspace.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "if states[j].Provider == \"eol\" {") {
        t.Fatalf("expected source contract is missing")
    }
}
