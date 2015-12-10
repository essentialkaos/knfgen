package main

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2015 Essential Kaos                         //
//      Essential Kaos Open Source License <http://essentialkaos.com/ekol?en>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"os"
	"strconv"
	"strings"

	"github.com/essentialkaos/ek/arg"
	"github.com/essentialkaos/ek/fmtc"
	"github.com/essentialkaos/ek/fmtutil"
	"github.com/essentialkaos/ek/knf"
	"github.com/essentialkaos/ek/usage"
)

// ////////////////////////////////////////////////////////////////////////////////// //

const (
	APP = "KNFGen"
	VER = "0.1"
)

const (
	ARG_NO_COLOR = "nc:no-color"
	ARG_HELP     = "h:help"
	ARG_VER      = "v:version"
)

// ////////////////////////////////////////////////////////////////////////////////// //

var argList = arg.Map{
	ARG_NO_COLOR: &arg.V{Type: arg.BOOL},
	ARG_HELP:     &arg.V{Type: arg.BOOL, Alias: "u:usage"},
	ARG_VER:      &arg.V{Type: arg.BOOL, Alias: "ver"},
}

// ////////////////////////////////////////////////////////////////////////////////// //

func main() {
	args, errs := arg.Parse(argList)

	if len(errs) != 0 {
		fmtc.NewLine()

		for _, err := range errs {
			fmtc.Printf("{r}%s{!}\n", err.Error())
		}

		os.Exit(1)
	}

	if arg.GetB(ARG_NO_COLOR) {
		fmtc.DisableColors = true
	}

	if arg.GetB(ARG_VER) {
		showAbout()
		return
	}

	if arg.GetB(ARG_HELP) || len(args) == 0 {
		showUsage()
		return
	}

	config, err := knf.Read(args[0])

	if err != nil {
		fmtc.Printf("{r}%s{!}\n", err.Error())
		os.Exit(1)
	}

	renderConfig(config)
}

// renderConfig render config data
func renderConfig(config *knf.Config) {
	fmtutil.Separator(false)

	var maxPropSize int

	for _, section := range config.Sections() {
		for _, prop := range config.Props(section) {
			maxPropSize = len(formatConstName(section, prop))
		}
	}

	formatString := getFormatString(maxPropSize)

	fmtc.Println("const (")

	for _, section := range config.Sections() {
		for _, prop := range config.Props(section) {
			fmtc.Printf(
				formatString,
				formatConstName(section, prop),
				section, prop,
			)
		}
	}

	fmtc.Println(")")

	fmtutil.Separator(false)
}

func formatConstName(section, prop string) string {
	fs := strings.ToUpper(section)
	fp := strings.ToUpper(prop)

	fp = strings.Replace(fp, "-", "_", -1)

	return fs + "_" + fp
}

func getFormatString(maxSize int) string {
	return "\t%" + strconv.Itoa(maxSize) + "s = \"%s:%s\"\n"
}

// ////////////////////////////////////////////////////////////////////////////////// //

func showUsage() {
	info := usage.NewInfo("knfgen", "config-file")

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
		Desc:    "Utility for generating go const code for KNF configs",
		Year:    2006,
		Owner:   "ESSENTIAL KAOS",
		License: "Essential Kaos Open Source License <https://essentialkaos.com/ekol?en>",
	}

	about.Render()
}
