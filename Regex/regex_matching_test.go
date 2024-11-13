package regexmatching

import "testing"

func TestIsMatch(t *testing.T) {
    tests := []struct {
        name    string
        s       string
        p       string
        want    bool
    }{
        // Basic matching without special characters
        {
            name: "exact match",
            s:    "abc",
            p:    "abc",
            want: true,
        },
        {
            name: "simple mismatch",
            s:    "abc",
            p:    "def",
            want: false,
        },

        // Single character matching with '.'
        {
            name: "single dot match",
            s:    "abc",
            p:    "a.c",
            want: true,
        },
        {
            name: "multiple dots",
            s:    "abc",
            p:    "...",
            want: true,
        },
        {
            name: "all dots with different length",
            s:    "abc",
            p:    "....",
            want: false,
        },

        // Star pattern tests
        {
            name: "simple star pattern",
            s:    "aaa",
            p:    "a*",
            want: true,
        },
        {
            name: "zero occurrence star",
            s:    "bcd",
            p:    "a*bcd",
            want: true,
        },
        {
            name: "multiple occurrences star",
            s:    "aaabcd",
            p:    "a*bcd",
            want: true,
        },
        {
            name: "multiple stars",
            s:    "aaabbbccc",
            p:    "a*b*c*",
            want: true,
        },

        // Combination of dot and star
        {
            name: "dot star pattern",
            s:    "abcdef",
            p:    ".*",
            want: true,
        },
        {
            name: "complex dot star pattern",
            s:    "abcdef",
            p:    "a.*ef",
            want: true,
        },
        {
            name: "dot star with specific chars",
            s:    "aabcd",
            p:    "a.*d",
            want: true,
        },

        // Edge cases
        {
            name: "empty string and pattern",
            s:    "",
            p:    "",
            want: true,
        },
        {
            name: "empty string with star pattern",
            s:    "",
            p:    "a*",
            want: true,
        },
        {
            name: "empty string with complex pattern",
            s:    "",
            p:    ".*",
            want: true,
        },
        {
            name: "empty pattern",
            s:    "abc",
            p:    "",
            want: false,
        },

        // Complex patterns
        {
            name: "alternating star pattern",
            s:    "aabbcc",
            p:    "a*b*c*",
            want: true,
        },
        {
            name: "nested star patterns",
            s:    "aaa",
            p:    "a*a*a*",
            want: true,
        },
        {
            name: "complex mix of dot and star",
            s:    "abcabc",
            p:    ".*abc",
            want: true,
        },
        {
            name: "star after dot",
            s:    "abcabc",
            p:    ".*c",
            want: true,
        },

        // Tricky cases
        {
            name: "consecutive stars",
            s:    "aaa",
            p:    "aa*a",
            want: true,
        },
        {
            name: "star with no match",
            s:    "bbb",
            p:    "a*bbb",
            want: true,
        },
        {
            name: "partial match",
            s:    "abc",
            p:    "abc.*def",
            want: false,
        },

        // Special character combinations
        {
            name: "mix of dot and char before star",
            s:    "abcdef",
            p:    "abc.*f",
            want: true,
        },
        {
            name: "multiple patterns",
            s:    "abcdef",
            p:    "a.*c.*f",
            want: true,
        },
        {
            name: "complex failure case",
            s:    "abcdef",
            p:    "a.*c.*e$",
            want: false,
        },

        // Long string cases
        {
            name: "long repeating pattern",
            s:    "aaaaaaaaaaaaab",
            p:    "a*b",
            want: true,
        },
        {
            name: "long mixed pattern",
            s:    "aaabbbccc",
            p:    "a*b*c*d*",
            want: true,
        },

        // Invalid patterns
        {
            name: "star without preceding char",
            s:    "abc",
            p:    "*abc",
            want: false,
        },
        {
            name: "consecutive stars without chars",
            s:    "abc",
            p:    "a**bc",
            want: false,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := IsMatch(tt.s, tt.p)
            if got != tt.want {
                t.Errorf("IsMatch(%q, %q) = %v, want %v",
                    tt.s, tt.p, got, tt.want)
            }
        })
    }
}

// Additional test for concurrent execution
func TestConcurrentIsMatch(t *testing.T) {
    // Test concurrent execution of IsMatch
    done := make(chan bool)
    for i := 0; i < 10; i++ {
        go func() {
            result := IsMatch("aaaa", "a*")
            if !result {
                t.Errorf("Concurrent test failed")
            }
            done <- true
        }()
    }

    // Wait for all goroutines to complete
    for i := 0; i < 10; i++ {
        <-done
    }
}

// Benchmark the IsMatch function
func BenchmarkIsMatch(b *testing.B) {
    for i := 0; i < b.N; i++ {
        IsMatch("aaaaaaaaaaaaab", "a*b")
    }
}
