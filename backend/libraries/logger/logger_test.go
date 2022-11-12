package logger

import (
	"bytes"
	"errors"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func BenchmarkDebugWithDebugLevel(b *testing.B) {
	SetLogLevel("debug")
	Log.Out = ioutil.Discard // Discard log output
	for n := 0; n < b.N; n++ {
		Debugf("Some message %s %s %s %s", "1", "2", "3", "4")
	}
	Log.Out = os.Stdout
}

func BenchmarkDebugWithInfoLevel(b *testing.B) {
	SetLogLevel("info")
	Log.Out = ioutil.Discard // Discard log output
	for n := 0; n < b.N; n++ {
		Debugf("Some message %s %s %s %s", "1", "2", "3", "4")
	}
	Log.Out = os.Stdout
}

func BenchmarkErrorWithErrorLevel(b *testing.B) {
	SetLogLevel("error")
	Log.Out = ioutil.Discard // Discard log output
	for n := 0; n < b.N; n++ {
		Errorf("Some message %s %s %s %s", errors.New("some error"), "1", "2", "3", "4")
	}
	Log.Out = os.Stdout
}

func BenchmarkErrorWithPanicLevel(b *testing.B) {
	SetLogLevel("panic")
	Log.Out = ioutil.Discard // Discard log output
	for n := 0; n < b.N; n++ {
		Errorf("Some message %s %s %s %s", errors.New("some error"), "1", "2", "3", "4")
	}
	Log.Out = os.Stdout
}

func TestTagsShouldWorkWithSimpleValues(t *testing.T) {
	buffer := &bytes.Buffer{}

	SetLogLevel("info")
	Log.Out = buffer
	Info(
		"Message",
		"tag1:foo",
	)

	out := buffer.String()
	if !strings.Contains(out, "[tag1:foo]") {
		t.Fail()
	}

	Log.Out = os.Stdout
}

func TestDebugHasLevelDebugTag(t *testing.T) {
	buffer := &bytes.Buffer{}

	SetLogLevel("debug")
	Log.Out = buffer
	Debug(
		"Message",
		"tag1:foo",
	)

	out := buffer.String()
	if !strings.Contains(out, "[level:debug]") {
		t.Fail()
	}

	Log.Out = os.Stdout
}

func TestInfoHasLevelInfoTag(t *testing.T) {
	buffer := &bytes.Buffer{}

	SetLogLevel("info")
	Log.Out = buffer
	Info(
		"Message",
		"tag1:foo",
	)

	out := buffer.String()
	if !strings.Contains(out, "[level:info]") {
		t.Fail()
	}

	Log.Out = os.Stdout
}

func TestWarnHasLevelWarnTag(t *testing.T) {
	buffer := &bytes.Buffer{}

	SetLogLevel("info")
	Log.Out = buffer
	Warn(
		"Message",
		"tag1:foo",
	)

	out := buffer.String()
	if !strings.Contains(out, "[level:warn]") {
		t.Fail()
	}

	Log.Out = os.Stdout
}

func TestErrorHasLevelErrorTag(t *testing.T) {
	buffer := &bytes.Buffer{}

	SetLogLevel("info")
	Log.Out = buffer
	err := errors.New("Some Error")
	Error(
		"Message",
		err,
		"tag1:foo",
	)

	out := buffer.String()
	if !strings.Contains(out, "[level:error]") {
		t.Fail()
	}

	if !strings.Contains(out, err.Error()) {
		t.Fatalf("expected to find '%s' in log message", err.Error())
	}

	Log.Out = os.Stdout
}

func TestTagsShouldWorkWithComplexValues(t *testing.T) {
	buffer := &bytes.Buffer{}

	SetLogLevel("info")
	Log.Out = buffer

	tag := "tag1:{\"foo\":\"bar\"}"
	Info("Message", tag)

	expected := "[tag1:\"{\\\"foo\\\":\\\"bar\\\"}\"]"
	if !strings.Contains(buffer.String(), expected) {
		t.Fail()
	}

	Log.Out = os.Stdout
}

func TestWrongTagsShouldBePartOfTheLogMessage(t *testing.T) {
	buffer := &bytes.Buffer{}

	SetLogLevel("info")
	Log.Out = buffer

	wrongTag := "wrongtag"
	validTag := "validtag:value"
	Info("Message", wrongTag, validTag)

	expected := `Message - Error parsing tags (wrongtag)`
	value := buffer.String()

	if !strings.Contains(value, expected) {
		t.Fail()
	}

	Log.Out = os.Stdout
}
