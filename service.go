package main

import (
	"errors"
	"strings"
)

// StringService provides operations on strings.
type StringService interface {
	Uppercase(string) (string, error)
	Downcase(string) (string, error)
	Count(string) int
	Palindrome(string) (bool, error)
}

type stringService struct{}

func (stringService) Uppercase(s string) (string, error) {
	if s == "" {
		return "", ErrEmpty
	}
	return strings.ToUpper(s), nil
}

func (stringService) Count(s string) int {
	return len(s)
}

func (stringService) Downcase(s string) (string, error) {
	if s == "" {
		return "", ErrEmpty
	}
	return strings.ToLower(s), nil
}

func (stringService) Palindrome(s string) (bool, error) {
	if s == "" {
		return false, ErrEmpty
	}
	return isPalindrome(s), nil
}

// ErrEmpty is returned when an input string is empty.
var ErrEmpty = errors.New("empty string")

// ServiceMiddleware is a chainable behavior modifier for StringService.
type ServiceMiddleware func(StringService) StringService

// Actual implementation of isPalindrome
func isPalindrome(input string) bool {
	for i := 0; i < len(input)/2; i++ {
		if input[i] != input[len(input)-i-1] {
			return false
		}
	}
	return true
}
