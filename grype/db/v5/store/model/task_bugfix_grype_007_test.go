package model

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixGrype007SourceContract(t *testing.T) {
    source, err := os.ReadFile("vulnerability_metadata.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "if err := json.Unmarshal(m.URLs.ToByteSlice(), &links); err != nil {") {
        t.Fatalf("expected source contract is missing")
    }
}
