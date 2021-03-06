package main

import (
	"time"

	"github.com/go-kit/kit/log"
)

func loggingMiddleware(logger log.Logger) ServiceMiddleware {
	return func(next StringService) StringService {
		return logmw{logger, next}
	}
}

type logmw struct {
	logger log.Logger
	StringService
}

func (mw logmw) Uppercase(s string) (output string, err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "uppercase",
			"input", s,
			"output", output,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	output, err = mw.StringService.Uppercase(s)
	return
}

func (mw logmw) Downcase(s string) (output string, err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "downcase",
			"input", s,
			"output", output,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	output, err = mw.StringService.Downcase(s)
	return
}

func (mw logmw) Count(s string) (n int) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "count",
			"input", s,
			"n", n,
			"took", time.Since(begin),
		)
	}(time.Now())

	n = mw.StringService.Count(s)
	return
}

func (mw logmw) Palindrome(s string) (p bool, err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "palindrome",
			"input", s,
			"p", p,
			"took", time.Since(begin),
		)
	}(time.Now())

	p, err = mw.StringService.Palindrome(s)
	return
}
