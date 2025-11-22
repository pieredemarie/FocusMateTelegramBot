package utils

import (
	"testing"
	"time"
)

func TestParseDuration(t *testing.T) {
    tests := []struct {
        input    string
        expected time.Duration
        wantErr  bool
    }{
        {"10s", 10 * time.Second, false},
        {"5m", 5 * time.Minute, false},
        {"2h", 2 * time.Hour, false},
        {"1d", 24 * time.Hour, false},

        {"1h30m", 1*time.Hour + 30*time.Minute, false},
        {"2h10m30s", 2*time.Hour + 10*time.Minute + 30*time.Second, false},
        {"1d4h30m", 24*time.Hour + 4*time.Hour + 30*time.Minute, false},

        {"abc", 0, true},     // неизвестный формат
        {"10x", 0, true},     // неправильный суффикс
        {"h10m", 0, true},    // начинается не с числа
        {"", 0, true},        // пустая строка
    }

    for _, tt := range tests {
        got, err := ParseDuration(tt.input)

        if tt.wantErr && err == nil {
            t.Errorf("ParseDuration(%q) expected error, got none", tt.input)
            continue
        }

        if !tt.wantErr && err != nil {
            t.Errorf("ParseDuration(%q) unexpected error: %v", tt.input, err)
            continue
        }

        if got != tt.expected {
            t.Errorf("ParseDuration(%q) = %v; want %v", tt.input, got, tt.expected)
        }
    }
}