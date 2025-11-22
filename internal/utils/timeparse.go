package utils

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

func ParseMessage(msg string) (string, time.Duration, error) {
	words := strings.Fields(msg)
	if words[0] != "/remind" {
		return "",0, errors.New("unkown command")
	}

	dur, err := ParseDuration(words[1])
	if err != nil {
		return "",0,err
	} 
	
	text := ""

	for i := 2;i < len(words);i++ {
		text = text + " " + words[i]
	}

	return text,dur,nil
}

func ParseDuration(str string) (time.Duration, error) {
    var total time.Duration
    var number string

	if len(str) == 0 {
		return 0, errors.New("string cannot be empty")
	}

    for _, r := range str {
        if r >= '0' && r <= '9' {
            number += string(r)
            continue
        }

        if number == "" {
            return 0, errors.New("invalid duration format")
        }

        value, _ := strconv.Atoi(number)

        switch r {
        case 's':
            total += time.Duration(value) * time.Second
        case 'm':
            total += time.Duration(value) * time.Minute
        case 'h':
            total += time.Duration(value) * time.Hour
        case 'd':
            total += time.Duration(value) * 24 * time.Hour
        default:
            return 0, errors.New("unknown duration suffix")
        }

        number = ""
    }

    return total, nil
}
