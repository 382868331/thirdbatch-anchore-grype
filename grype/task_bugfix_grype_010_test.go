package grype

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixGrype010SourceContract(t *testing.T) {
    source, err := os.ReadFile("vulnerability_matcher.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "if rule.Package.Location != \"\" && !strings.ContainsRune(rule.Package.Location, '*') {") {
        t.Fatalf("expected source contract is missing")
    }
    if strings.Contains(string(source), "if rule.Package.Location == \"\" && !strings.ContainsRune(rule.Package.Location, '*') {") {
        t.Fatalf("mutated source contract is still present")
    }
}
