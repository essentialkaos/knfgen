package main

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2017 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"os"
	"strconv"
	"strings"

	"pkg.re/essentialkaos/ek.v8/arg"
	"pkg.re/essentialkaos/ek.v8/env"
	"pkg.re/essentialkaos/ek.v8/fmtc"
	"pkg.re/essentialkaos/ek.v8/fmtutil"
	"pkg.re/essentialkaos/ek.v8/fsutil"
	"pkg.re/essentialkaos/ek.v8/knf"
	"pkg.re/essentialkaos/ek.v8/mathutil"
	"pkg.re/essentialkaos/ek.v8/usage"
)

// ////////////////////////////////////////////////////////////////////////////////// //

const (
	APP  = "KNFGen"
	VER  = "0.5.0"
	DESC = "Utility for generating Golang const code for KNF configs"
)

const (
	ARG_SEPARATORS = "s:separators"
	ARG_NO_COLOR   = "nc:no-color"
	ARG_HELP       = "h:help"
	ARG_VER        = "v:version"
)

// ////////////////////////////////////////////////////////////////////////////////// //

var argMap = arg.Map{
	ARG_SEPARATORS: {Type: arg.BOOL},
	ARG_NO_COLOR:   {Type: arg.BOOL},
	ARG_HELP:       {Type: arg.BOOL, Alias: "u:usage"},
	ARG_VER:        {Type: arg.BOOL, Alias: "ver"},
}

var rawOutput = false

// ////////////////////////////////////////////////////////////////////////////////// //

func main() {
	args, errs := arg.Parse(argMap)

	if len(errs) != 0 {
		for _, err := range errs {
			printError(err.Error())
		}

		os.Exit(1)
	}

	configureUI()

	if arg.GetB(ARG_VER) {
		showAbout()
		return
	}

	if arg.GetB(ARG_HELP) || len(args) == 0 {
		showUsage()
		return
	}

	process(args[0])
}

// configureUI configure user interface
func configureUI() {
	envVars := env.Get()
	term := envVars.GetS("TERM")

	fmtc.DisableColors = true
	rawOutput = true

	if term != "" {
		switch {
		case strings.Contains(term, "xterm"),
			strings.Contains(term, "color"),
			term == "screen":
			fmtc.DisableColors = false
			rawOutput = false
		}
	}

	if arg.GetB(ARG_NO_COLOR) {
		fmtc.DisableColors = true
	}

	if !fsutil.IsCharacterDevice("/dev/stdout") && envVars.GetS("FAKETTY") == "" {
		fmtc.DisableColors = true
		rawOutput = true
	}
}

// process start config processing
func process(file string) {
	config, err := knf.Read(file)

	if err != nil {
		if !rawOutput {
			printError(err.Error())
		}

		os.Exit(1)
	}

	renderConfig(config)
}

// renderConfig render config data
func renderConfig(config *knf.Config) {
	if !rawOutput {
		fmtutil.Separator(false)
	}

	var maxPropSize int

	for _, section := range config.Sections() {
		for _, prop := range config.Props(section) {
			maxPropSize = mathutil.Max(maxPropSize, len(formatConstName(section, prop)))
		}
	}

	formatString := getFormatString(maxPropSize)
	sectionsTotal := len(config.Sections())

	fmtc.Println("{*}const ({!}")

	for sectionIndex, section := range config.Sections() {
		for _, prop := range config.Props(section) {
			fmtc.Printf(
				formatString,
				formatConstName(section, prop),
				section, prop,
			)
		}

		if arg.GetB(ARG_SEPARATORS) && sectionIndex < sectionsTotal-1 {
			fmtc.NewLine()
		}
	}

	fmtc.Println("{*}){!}")

	if !rawOutput {
		fmtutil.Separator(false)
	}
}

// formatConstName return const name
func formatConstName(section, prop string) string {
	fs := strings.ToUpper(section)
	fp := strings.ToUpper(prop)

	fs = strings.Replace(fs, "-", "_", -1)
	fp = strings.Replace(fp, "-", "_", -1)

	return fs + "_" + fp
}

// getFormatString return format string
func getFormatString(maxSize int) string {
	return "\t%-" + strconv.Itoa(maxSize) + "s = {y}\"%s:%s\"{!}\n"
}

// printError prints error message to console
func printError(f string, a ...interface{}) {
	fmtc.Printf("{r}"+f+"{!}\n", a...)
}

// printError prints warning message to console
func printWarn(f string, a ...interface{}) {
	fmtc.Printf("{y}"+f+"{!}\n", a...)
}

// ////////////////////////////////////////////////////////////////////////////////// //

func showUsage() {
	info := usage.NewInfo("", "config-file")

	info.AddOption(ARG_SEPARATORS, "Add new lines between sections")
	info.AddOption(ARG_NO_COLOR, "Disable colors in output")
	info.AddOption(ARG_HELP, "Show this help message")
	info.AddOption(ARG_VER, "Show version")

	info.AddExample("app.conf", "Generate copy-paste code for app.conf")

	info.Render()
}

func showAbout() {
	about := &usage.About{
		App:     APP,
		Version: VER,
		Desc:    DESC,
		Year:    2006,
		Owner:   "ESSENTIAL KAOS",
		License: "Essential Kaos Open Source License <https://essentialkaos.com/ekol>",
	}

	about.Render()
}
