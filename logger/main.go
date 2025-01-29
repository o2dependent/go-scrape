package logger

import "github.com/fatih/color"

var Err = color.New(color.FgRed).Add(color.Underline)
var Warn = color.New(color.FgHiYellow)
var Info = color.New(color.FgCyan)
var InfoAccent = color.New(color.FgHiCyan).Add(color.Underline)
var DevHighlight = color.New(color.BgHiRed).Add(color.FgBlack)
