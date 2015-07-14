package ansi

import (
	"fmt"
	"os"

	"golang.org/x/crypto/ssh/terminal"
)

// Printf prints a formatted output to the terminal
func Printf(f string, v ...interface{}) {
	fmt.Printf(f, v)
}

// Color represents a color that can be set
type Color int

// Constant colors definition
const (
	Black   Color = 0
	Red     Color = 1
	Green   Color = 2
	Yellow  Color = 3
	Blue    Color = 4
	Magenta Color = 5
	Cyan    Color = 6
	White   Color = 7
	Default Color = 9
)

type formatter interface {
	ResetColor()
	SetForeground(c Color)
	SetBackground(c Color)
	SetForeAndBackground(fore, back Color)
	CursorSave()
	CursorRestore()
	ClearToEndOfLine()
	ClearToBeginOfLine()
	ClearLine()
	ClearScreenAndMoveOrigin()
	MoveCursor(row, col int)
	CursorUp(rows int)
	CursorDown(rows int)
}

type nullFormater struct{}

func (f *nullFormater) ResetColor()                           {}
func (f *nullFormater) SetForeground(c Color)                 {}
func (f *nullFormater) SetBackground(c Color)                 {}
func (f *nullFormater) SetForeAndBackground(fore, back Color) {}
func (f *nullFormater) CursorSave()                           {}
func (f *nullFormater) CursorRestore()                        {}
func (f *nullFormater) ClearToEndOfLine()                     {}
func (f *nullFormater) ClearToBeginOfLine()                   {}
func (f *nullFormater) ClearLine()                            {}
func (f *nullFormater) ClearScreenAndMoveOrigin()             {}
func (f *nullFormater) MoveCursor(row, col int)               {}
func (f *nullFormater) CursorUp(rows int)                     {}
func (f *nullFormater) CursorDown(rows int)                   {}

type escapeFormatter struct {
	csi string
}

func (f *escapeFormatter) ResetColor() {
	fmt.Printf("%s%dm", f.csi, 0)
}

func (f *escapeFormatter) SetForeground(c Color) {
	fmt.Printf("%s%dm", f.csi, 30+c)
}

func (f *escapeFormatter) SetBackground(c Color) {
	fmt.Printf("%s%dm", f.csi, 40+c)
}

func (f *escapeFormatter) SetForeAndBackground(fore, back Color) {
	fmt.Printf("%s%d;%dm", f.csi, 30+fore, 40+back)
}

func (f *escapeFormatter) CursorSave() {
	fmt.Printf("%ss", f.csi)
}

func (f *escapeFormatter) CursorRestore() {
	fmt.Printf("%su", f.csi)
}

func (f *escapeFormatter) ClearToEndOfLine() {
	fmt.Printf("%sK", f.csi)
}
func (f *escapeFormatter) ClearToBeginOfLine() {
	fmt.Printf("%s1K", f.csi)
}
func (f *escapeFormatter) ClearLine() {
	fmt.Printf("%s2K", f.csi)
}
func (f *escapeFormatter) ClearScreenAndMoveOrigin() {
	fmt.Printf("%s2J", f.csi)
}
func (f *escapeFormatter) MoveCursor(row, col int) {
	fmt.Printf("%s%d;%dH", f.csi, row, col)
}
func (f *escapeFormatter) CursorUp(rows int) {
	fmt.Printf("%s%dA", f.csi, rows)
}
func (f *escapeFormatter) CursorDown(rows int) {
	fmt.Printf("%s%dB", f.csi, rows)
}

var f formatter

func init() {
	if terminal.IsTerminal(int(os.Stdout.Fd())) == true {
		f = &escapeFormatter{
			csi: `\033[`,
		}
	} else {
		f = &nullFormater{}
	}
}

// ResetColor sets the color and format to terminal default
func ResetColor() {
	f.ResetColor()
}

// SetForeground sets the foreground color
func SetForeground(c Color) {
	f.SetForeground(c)
}

// SetBackground sets the background color
func SetBackground(c Color) {
	f.SetBackground(c)
}

// SetForeAndBackground sets the fore and background color
func SetForeAndBackground(fore, back Color) {
	f.SetForeAndBackground(fore, back)
}

// CursorSave saves the current cursor position
func CursorSave() {
	f.CursorSave()
}

// CursorRestore puts cursor position to saved position
func CursorRestore() {
	f.CursorRestore()
}

// ClearToEndOfLine removes all character until end of current line
func ClearToEndOfLine() {
	f.ClearToEndOfLine()
}

// ClearToBeginOfLine removes all character until end of current
// line. Cursor position does not change
func ClearToBeginOfLine() {
	f.ClearToBeginOfLine()
}

// ClearLine clears the current line. Cursors position does not change
func ClearLine() {
	f.ClearLine()
}

// ClearScreenAndMoveOrigin clear the whole screen and moves cursor to
// origin
func ClearScreenAndMoveOrigin() {
	f.ClearScreenAndMoveOrigin()
}

// MoveCursor sets the cursor position
func MoveCursor(row, col int) {
	f.MoveCursor(row, col)
}

// CursorUp moves the cursor rows rows up
func CursorUp(rows int) {
	f.CursorUp(rows)
}

// CursorDown moves the cursors rows rows down
func CursorDown(rows int) {
	f.CursorDown(rows)
}
