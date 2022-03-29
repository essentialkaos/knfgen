package main

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2022 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"os"
	"strconv"
	"strings"

	"github.com/essentialkaos/ek/v12/env"
	"github.com/essentialkaos/ek/v12/fmtc"
	"github.com/essentialkaos/ek/v12/fmtutil"
	"github.com/essentialkaos/ek/v12/fsutil"
	"github.com/essentialkaos/ek/v12/knf"
	"github.com/essentialkaos/ek/v12/mathutil"
	"github.com/essentialkaos/ek/v12/options"
	"github.com/essentialkaos/ek/v12/usage"
	"github.com/essentialkaos/ek/v12/usage/update"
)

// ////////////////////////////////////////////////////////////////////////////////// //

const (
	APP  = "KNFGen"
	VER  = "0.7.2"
	DESC = "Utility for generating Golang const code for KNF configs"
)

const (
	OPT_SEPARATORS = "s:separators"
	OPT_NO_COLOR   = "nc:no-color"
	OPT_HELP       = "h:help"
	OPT_VER        = "v:version"
)

// ////////////////////////////////////////////////////////////////////////////////// //

var optMap = options.Map{
	OPT_SEPARATORS: {Type: options.BOOL},
	OPT_NO_COLOR:   {Type: options.BOOL},
	OPT_HELP:       {Type: options.BOOL, Alias: "u:usage"},
	OPT_VER:        {Type: options.BOOL, Alias: "ver"},
}

var rawOutput = false

// ////////////////////////////////////////////////////////////////////////////////// //

func main() {
	args, errs := options.Parse(optMap)

	if len(errs) != 0 {
		for _, err := range errs {
			printError(err.Error())
		}

		os.Exit(1)
	}

	configureUI()

	if options.GetB(OPT_VER) {
		showAbout()
		return
	}

	if options.GetB(OPT_HELP) || len(args) == 0 {
		showUsage()
		return
	}

	process(args[0])
}

// configureUI configures user interface
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

	if options.GetB(OPT_NO_COLOR) {
		fmtc.DisableColors = true
	}

	if !fsutil.IsCharacterDevice("/dev/stdout") && envVars.GetS("FAKETTY") == "" {
		fmtc.DisableColors = true
		rawOutput = true
	}
}

// process starts config processing
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

// renderConfig renders config data
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

		if options.GetB(OPT_SEPARATORS) && sectionIndex < sectionsTotal-1 {
			fmtc.NewLine()
		}
	}

	fmtc.Println("{*}){!}")

	if !rawOutput {
		fmtutil.Separator(false)
	}
}

// formatConstName returns const name
func formatConstName(section, prop string) string {
	fs := strings.ToUpper(section)
	fp := strings.ToUpper(prop)

	fs = strings.Replace(fs, "-", "_", -1)
	fp = strings.Replace(fp, "-", "_", -1)

	return fs + "_" + fp
}

// getFormatString returns format string
func getFormatString(maxSize int) string {
	return "\t%-" + strconv.Itoa(maxSize) + "s = {y}\"%s:%s\"{!}\n"
}

// printError prints error message to console
func printError(f string, a ...interface{}) {
	fmtc.Fprintf(os.Stderr, "{r}"+f+"{!}\n", a...)
}

// printError prints warning message to console
func printWarn(f string, a ...interface{}) {
	fmtc.Fprintf(os.Stderr, "{y}"+f+"{!}\n", a...)
}

// ////////////////////////////////////////////////////////////////////////////////// //

// showUsage shows usage info
func showUsage() {
	info := usage.NewInfo("", "config-file")

	info.AddOption(OPT_SEPARATORS, "Add new lines between sections")
	info.AddOption(OPT_NO_COLOR, "Disable colors in output")
	info.AddOption(OPT_HELP, "Show this help message")
	info.AddOption(OPT_VER, "Show version")

	info.AddExample("app.conf", "Generate copy-paste code for app.conf")

	info.Render()
}

// showAbout shows info about version
func showAbout() {
	about := &usage.About{
		App:           APP,
		Version:       VER,
		Desc:          DESC,
		Year:          2006,
		Owner:         "ESSENTIAL KAOS",
		License:       "Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>",
		UpdateChecker: usage.UpdateChecker{"essentialkaos/knfgen", update.GitHubChecker},
	}

	about.Render()
}
