package errors_test

import (
	"io"
	"log"
	"testing"

	"github.com/streamwest-1629/go-utilpkgs/errors"
)

func TestNew(t *testing.T) {
	log.Println(errors.New("test error").Error())
	log.Println(errors.New("test error containing key/value",
		errors.NewKV("arg1", "value1"),
		errors.NewKV("arg2", "value2"),
	).Error())
	log.Println(errors.NewWith("test error containing error (io.EOF)", io.EOF).Error())
	log.Println(errors.NewWith(
		"test error containing error (io.EOF) and key/value",
		io.EOF,
		errors.NewKV("arg1", "value1"),
		errors.NewKV("arg2", "value2"),
	).Error())
}
