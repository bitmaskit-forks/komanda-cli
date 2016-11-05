// Copyright 2014 The gocui Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gocui

import (
	"errors"
	"strconv"
)

type escapeInterpreter struct {
	state                  escapeState
	curch                  rune
	csiParam               []string
	curFgColor, curBgColor Attribute
	mode                   OutputMode
}

type escapeState int

const (
	stateNone escapeState = iota
	stateEscape
	stateCSI
	stateParams
)

var (
	errNotCSI        = errors.New("Not a CSI escape sequence")
	errCSINotANumber = errors.New("CSI escape sequence was expecting a number or a ;")
	errCSIParseError = errors.New("CSI escape sequence parsing error")
	errCSITooLong    = errors.New("CSI escape sequence is too long")
)

// runes in case of error will output the non-parsed runes as a string.
func (ei *escapeInterpreter) runes() []rune {
	switch ei.state {
	case stateNone:
		return []rune{0x1b}
	case stateEscape:
		return []rune{0x1b, ei.curch}
	case stateCSI:
		return []rune{0x1b, '[', ei.curch}
	case stateParams:
		ret := []rune{0x1b, '['}
		for _, s := range ei.csiParam {
			ret = append(ret, []rune(s)...)
			ret = append(ret, ';')
		}
		return append(ret, ei.curch)
	}
	return nil
}

// newEscapeInterpreter returns an escapeInterpreter that will be able to parse
// terminal escape sequences.
func newEscapeInterpreter(mode OutputMode) *escapeInterpreter {
	ei := &escapeInterpreter{
		state:      stateNone,
		curFgColor: ColorDefault,
		curBgColor: ColorDefault,
		mode:       mode,
	}
	return ei
}

// reset sets the escapeInterpreter in initial state.
func (ei *escapeInterpreter) reset() {
	ei.state = stateNone
	ei.curFgColor = ColorDefault
	ei.curBgColor = ColorDefault
	ei.csiParam = nil
}

// paramToColor returns an attribute given a terminfo coloring.
func paramToColor(p int) Attribute {
	switch p {
	case 0:
		return ColorBlack
	case 1:
		return ColorRed
	case 2:
		return ColorGreen
	case 3:
		return ColorYellow
	case 4:
		return ColorBlue
	case 5:
		return ColorMagenta
	case 6:
		return ColorCyan
	case 7:
		return ColorWhite
	}
	return ColorDefault
}

// parseOne parses a rune. If isEscape is true, it means that the rune is part
// of an escape sequence, and as such should not be printed verbatim. Otherwise,
// it's not an escape sequence.
func (ei *escapeInterpreter) parseOne(ch rune) (isEscape bool, err error) {
	// Sanity checks to make sure we're not parsing something totally bogus.
	if len(ei.csiParam) > 20 {
		return false, errCSITooLong
	}
	if len(ei.csiParam) > 0 && len(ei.csiParam[len(ei.csiParam)-1]) > 255 {
		return false, errCSITooLong
	}
	ei.curch = ch
	switch ei.state {
	case stateNone:
		if ch == 0x1b {
			ei.state = stateEscape
			return true, nil
		}
		return false, nil
	case stateEscape:
		if ch == '[' {
			ei.state = stateCSI
			return true, nil
		}
		return false, errNotCSI
	case stateCSI:
		if ch >= '0' && ch <= '9' {
			ei.state = stateParams
			ei.csiParam = append(ei.csiParam, string(ch))
			return true, nil
		}
		return false, errCSINotANumber
	case stateParams:
		switch {
		case ch >= '0' && ch <= '9':
			ei.csiParam[len(ei.csiParam)-1] += string(ch)
			return true, nil
		case ch == ';':
			ei.csiParam = append(ei.csiParam, "")
			return true, nil
		case ch == 'm':
			if len(ei.csiParam) < 1 {
				return false, errCSIParseError
			}

			var err error

			switch ei.mode {
			case OutputMode(OutputNormal):
				err = ei.outputNormal()
			case OutputMode(Output256):
				err = ei.output256()
			}

			ei.state = stateNone
			ei.csiParam = nil
			return true, err
		}
	}
	return false, nil
}

// outputNormal  provides 8 different colors:
// black, red, green, yellow, blue, magenta, cyan, white
func (ei *escapeInterpreter) outputNormal() error {

	for _, param := range ei.csiParam {
		p, err := strconv.Atoi(param)

		if err != nil {
			return errCSIParseError
		}

		switch {
		case p >= 30 && p <= 37:
			ei.curFgColor = paramToColor(p - 30)
		case p >= 40 && p <= 47:
			ei.curBgColor = paramToColor(p - 40)
		case p == 1:
			ei.curFgColor |= AttrBold
		case p == 4:
			ei.curFgColor |= AttrUnderline
		case p == 7:
			ei.curFgColor |= AttrReverse
		case p == 0 || p == 39:
			ei.curFgColor = ColorDefault
			ei.curBgColor = ColorDefault
		}
	}

	return nil
}

// output256 allows you to leverage the 256 terminal mode
//    0x01 - 0x08: the 8 colors as in OutputNormal
//    0x09 - 0x10: Color* | AttrBold
//    0x11 - 0xe8: 216 different colors
//    0xe9 - 0x1ff: 24 different shades of grey
func (ei *escapeInterpreter) output256() error {

	if len(ei.csiParam) < 3 {
		ei.curFgColor = ColorDefault
		ei.curBgColor = ColorDefault

		return nil
	}

	fgbgParam, err := strconv.Atoi(ei.csiParam[0])

	if err != nil {
		return errCSIParseError
	}

	if fgbgParam == 38 {
		ei.curFgColor = 38

		for _, param := range ei.csiParam[1:] {
			p, err := strconv.Atoi(param)

			if err != nil {
				return errCSIParseError
			}

			ei.curFgColor |= Attribute(p)
		}

	} else if fgbgParam == 48 {
		ei.curBgColor = 48

		for _, param := range ei.csiParam[1:] {
			p, err := strconv.Atoi(param)

			if err != nil {
				return errCSIParseError
			}

			ei.curFgColor |= Attribute(p)
		}

	} else {
		return errCSIParseError
	}

	return nil
}
