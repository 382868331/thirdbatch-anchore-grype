package rootio

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixGrype015SourceContract(t *testing.T) {
    source, err := os.ReadFile("rootio.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "return false") {
        t.Fatalf("expected source contract is missing")
    }
}
