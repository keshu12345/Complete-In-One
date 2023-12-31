// Copyright 2020 New Relic Corporation. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package jsonx

import (
	"bytes"
	"fmt"
	"math"
	"testing"
)

func TestAppendFloat(t *testing.T) {
	buf := &bytes.Buffer{}

	err := AppendFloat(buf, math.NaN())
	if err == nil {
		t.Error("AppendFloat(NaN) should return an error")
	}

	err = AppendFloat(buf, math.Inf(1))
	if err == nil {
		t.Error("AppendFloat(+Inf) should return an error")
	}

	err = AppendFloat(buf, math.Inf(-1))
	if err == nil {
		t.Error("AppendFloat(-Inf) should return an error")
	}
}

func TestAppendFloat32(t *testing.T) {
	buf := &bytes.Buffer{}

	err := AppendFloat32(buf, float32(math.NaN()))
	if err == nil {
		t.Error("AppendFloat(NaN) should return an error")
	}

	err = AppendFloat32(buf, float32(math.Inf(1)))
	if err == nil {
		t.Error("AppendFloat(+Inf) should return an error")
	}

	err = AppendFloat32(buf, float32(math.Inf(-1)))
	if err == nil {
		t.Error("AppendFloat(-Inf) should return an error")
	}

	err = AppendFloat32(buf, float32(12.5))
	if err != nil {
		t.Error("AppendFloat(12.5) should not return an error")
	}
}

func TestAppendFloats(t *testing.T) {
	buf := &bytes.Buffer{}

	AppendFloatArray(buf)
	if want, got := "[]", buf.String(); want != got {
		t.Errorf("AppendFloatArray(buf)=%q want=%q", got, want)
	}

	buf.Reset()
	AppendFloatArray(buf, 3.14)
	if want, got := "[3.14]", buf.String(); want != got {
		t.Errorf("AppendFloatArray(buf)=%q want=%q", got, want)
	}

	buf.Reset()
	AppendFloatArray(buf, 1, 2)
	if want, got := "[1,2]", buf.String(); want != got {
		t.Errorf("AppendFloatArray(buf)=%q want=%q", got, want)
	}
}

func TestAppendInt(t *testing.T) {
	buf := &bytes.Buffer{}

	AppendInt(buf, 42)
	if got := buf.String(); got != "42" {
		t.Errorf("AppendUint(42) = %#q want %#q", got, "42")
	}

	buf.Reset()
	AppendInt(buf, -42)
	if got := buf.String(); got != "-42" {
		t.Errorf("AppendUint(-42) = %#q want %#q", got, "-42")
	}
}

func TestAppendIntArray(t *testing.T) {
	buf := &bytes.Buffer{}

	AppendIntArray(buf)
	if want, got := "[]", buf.String(); want != got {
		t.Errorf("AppendIntArray(buf)=%q want=%q", got, want)
	}

	buf.Reset()
	AppendIntArray(buf, 42)
	if want, got := "[42]", buf.String(); want != got {
		t.Errorf("AppendIntArray(buf)=%q want=%q", got, want)
	}

	buf.Reset()
	AppendIntArray(buf, 1, -2)
	if want, got := "[1,-2]", buf.String(); want != got {
		t.Errorf("AppendIntArray(buf)=%q want=%q", got, want)
	}

	buf.Reset()
	AppendIntArray(buf, 1, -2, 0)
	if want, got := "[1,-2,0]", buf.String(); want != got {
		t.Errorf("AppendIntArray(buf)=%q want=%q", got, want)
	}
}

func TestAppendUint(t *testing.T) {
	buf := &bytes.Buffer{}

	AppendUint(buf, 42)
	if got := buf.String(); got != "42" {
		t.Errorf("AppendUint(42) = %#q want %#q", got, "42")
	}
}

func TestAppendUintArray(t *testing.T) {
	buf := &bytes.Buffer{}

	AppendUintArray(buf)
	if want, got := "[]", buf.String(); want != got {
		t.Errorf("AppendUintArray(buf)=%q want=%q", got, want)
	}

	buf.Reset()
	AppendUintArray(buf, 42)
	if want, got := "[42]", buf.String(); want != got {
		t.Errorf("AppendUintArray(buf)=%q want=%q", got, want)
	}

	buf.Reset()
	AppendUintArray(buf, 1, 2)
	if want, got := "[1,2]", buf.String(); want != got {
		t.Errorf("AppendUintArray(buf)=%q want=%q", got, want)
	}

	buf.Reset()
	AppendUintArray(buf, 1, 2, 3)
	if want, got := "[1,2,3]", buf.String(); want != got {
		t.Errorf("AppendUintArray(buf)=%q want=%q", got, want)
	}
}

var encodeStringTests = []struct {
	in  string
	out string
}{
	{"\x00", `"\u0000"`},
	{"\x01", `"\u0001"`},
	{"\x02", `"\u0002"`},
	{"\x03", `"\u0003"`},
	{"\x04", `"\u0004"`},
	{"\x05", `"\u0005"`},
	{"\x06", `"\u0006"`},
	{"\x07", `"\u0007"`},
	{"\x08", `"\u0008"`},
	{"\x09", `"\t"`},
	{"\x0a", `"\n"`},
	{"\x0b", `"\u000b"`},
	{"\x0c", `"\u000c"`},
	{"\x0d", `"\r"`},
	{"\x0e", `"\u000e"`},
	{"\x0f", `"\u000f"`},
	{"\x10", `"\u0010"`},
	{"\x11", `"\u0011"`},
	{"\x12", `"\u0012"`},
	{"\x13", `"\u0013"`},
	{"\x14", `"\u0014"`},
	{"\x15", `"\u0015"`},
	{"\x16", `"\u0016"`},
	{"\x17", `"\u0017"`},
	{"\x18", `"\u0018"`},
	{"\x19", `"\u0019"`},
	{"\x1a", `"\u001a"`},
	{"\x1b", `"\u001b"`},
	{"\x1c", `"\u001c"`},
	{"\x1d", `"\u001d"`},
	{"\x1e", `"\u001e"`},
	{"\x1f", `"\u001f"`},
	{"\\", `"\\"`},
	{`"`, `"\""`},
	{"the\u2028quick\t\nbrown\u2029fox", `"the\u2028quick\t\nbrown\u2029fox"`},

	//extra edge cases
	{string([]byte{237, 159, 193}), `"\ufffd\ufffd\ufffd"`},                            // invalid utf8
	{string([]byte{55, 237, 159, 193, 55}), `"7\ufffd\ufffd\ufffd7"`},                  // invalid utf8 surrounded by valid utf8
	{`abcdefghijklmnopqrstuvwxyz1234567890`, `"abcdefghijklmnopqrstuvwxyz1234567890"`}, // alphanumeric
	{"'", `"'"`},
	{``, `""`},
	{`\`, `"\\"`},
	{fmt.Sprintf("%c", rune(65533)), fmt.Sprintf("\"%c\"", rune(65533))}, // invalid rune utf8 symbol (valid utf8)
}

func TestAppendString(t *testing.T) {
	buf := &bytes.Buffer{}

	for _, tt := range encodeStringTests {
		buf.Reset()

		AppendString(buf, tt.in)
		if got := buf.String(); got != tt.out {
			t.Errorf("AppendString(%q) = %#q, want %#q", tt.in, got, tt.out)
		}
	}
}

func TestAppendStringArray(t *testing.T) {
	buf := &bytes.Buffer{}

	var encodeStringArrayTests = []struct {
		in  []string
		out string
	}{
		{
			in: []string{
				"hi",
				"foo",
			},
			out: `["hi","foo"]`,
		},
		{
			in: []string{
				"foo",
			},
			out: `["foo"]`,
		},
		{
			in:  []string{},
			out: `[]`,
		},
	}

	for _, tt := range encodeStringArrayTests {
		buf.Reset()

		AppendStringArray(buf, tt.in...)
		if got := buf.String(); got != tt.out {
			t.Errorf("AppendString(%q) = %#q, want %#q", tt.in, got, tt.out)
		}
	}
}
func BenchmarkAppendString(b *testing.B) {
	buf := &bytes.Buffer{}

	for i := 0; i < b.N; i++ {
		AppendString(buf, "s")
	}
}

func BenchmarkAppendString10(b *testing.B) {
	buf := &bytes.Buffer{}

	for i := 0; i < b.N; i++ {
		AppendString(buf, "qwertyuiop")
	}
}
