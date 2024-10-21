package cli

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2024 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/essentialkaos/ek/v13/fmtc"
	"github.com/essentialkaos/ek/v13/knf"
	"github.com/essentialkaos/ek/v13/mathutil"
	"github.com/essentialkaos/ek/v13/options"
	"github.com/essentialkaos/ek/v13/pager"
	"github.com/essentialkaos/ek/v13/support"
	"github.com/essentialkaos/ek/v13/support/apps"
	"github.com/essentialkaos/ek/v13/support/deps"
	"github.com/essentialkaos/ek/v13/terminal"
	"github.com/essentialkaos/ek/v13/terminal/tty"
	"github.com/essentialkaos/ek/v13/usage"
	"github.com/essentialkaos/ek/v13/usage/completion/bash"
	"github.com/essentialkaos/ek/v13/usage/completion/fish"
	"github.com/essentialkaos/ek/v13/usage/completion/zsh"
	"github.com/essentialkaos/ek/v13/usage/man"
	"github.com/essentialkaos/ek/v13/usage/update"
)

// ////////////////////////////////////////////////////////////////////////////////// //

const (
	APP  = "knfgen"
	VER  = "1.0.2"
	DESC = "Utility for generating Golang const code for KNF configuration files"
)

const (
	OPT_SEPARATORS = "S:separators"
	OPT_UNITED     = "U:united"
	OPT_NO_COLOR   = "nc:no-color"
	OPT_HELP       = "h:help"
	OPT_VER        = "v:version"

	OPT_VERB_VER     = "vv:verbose-version"
	OPT_COMPLETION   = "completion"
	OPT_GENERATE_MAN = "generate-man"
)

// ////////////////////////////////////////////////////////////////////////////////// //

var optMap = options.Map{
	OPT_SEPARATORS: {Type: options.BOOL},
	OPT_UNITED:     {Type: options.BOOL},
	OPT_NO_COLOR:   {Type: options.BOOL},
	OPT_HELP:       {Type: options.BOOL},
	OPT_VER:        {Type: options.MIXED},

	OPT_VERB_VER:     {Type: options.BOOL},
	OPT_COMPLETION:   {},
	OPT_GENERATE_MAN: {Type: options.BOOL},
}

// color tags for app name and version
var colorTagApp, colorTagVer string

// tabSymbol contains tab symbol
var tabSymbol = "\t"

// ////////////////////////////////////////////////////////////////////////////////// //

// Run is main utility function
func Run(gitRev string, gomod []byte) {
	preConfigureUI()

	args, errs := options.Parse(optMap)

	if !errs.IsEmpty() {
		terminal.Error("Options parsing errors:")
		terminal.Error(errs.Error("- "))
		os.Exit(1)
	}

	configureUI()

	switch {
	case options.Has(OPT_COMPLETION):
		os.Exit(printCompletion())
	case options.Has(OPT_GENERATE_MAN):
		printMan()
		os.Exit(0)
	case options.GetB(OPT_VER):
		genAbout(gitRev).Print(options.GetS(OPT_VER))
		os.Exit(0)
	case options.GetB(OPT_VERB_VER):
		support.Collect(APP, VER).
			WithRevision(gitRev).
			WithDeps(deps.Extract(gomod)).
			WithApps(apps.Golang()).
			Print()
		os.Exit(0)
	case options.GetB(OPT_HELP) || len(args) == 0:
		genUsage().Print()
		os.Exit(0)
	}

	process(args.Get(0).Clean().String())
}

// preConfigureUI preconfigures UI based on information about user terminal
func preConfigureUI() {
	if !tty.IsTTY() {
		fmtc.DisableColors = true
	}

	switch {
	case fmtc.IsTrueColorSupported():
		colorTagApp, colorTagVer = "{*}{#784FFF}", "{#784FFF}"
	case fmtc.Is256ColorsSupported():
		colorTagApp, colorTagVer = "{*}{#99}", "{#99}"
	default:
		colorTagApp, colorTagVer = "{*}{c}", "{c}"
	}
}

// configureUI configures user interface
func configureUI() {
	if options.GetB(OPT_NO_COLOR) {
		fmtc.DisableColors = true
	}

	isTmux, _ := tty.IsTMUX()

	if isTmux {
		tabSymbol = "  "
	}
}

// process starts config processing
func process(file string) {
	config, err := knf.Read(file)

	if err != nil {
		terminal.Error(err)
		os.Exit(1)
	}

	if pager.Setup() == nil {
		defer pager.Complete()
	}

	renderConfig(config)

	if options.GetB(OPT_UNITED) {
		renderUnitedConfig(config)
	}
}

// renderConfig renders config data
func renderConfig(config *knf.Config) {
	var maxPropSize int

	for _, section := range config.Sections() {
		for _, prop := range config.Props(section) {
			maxPropSize = mathutil.Max(maxPropSize, len(formatConstName(section, prop)))
		}
	}

	formatString := getFormatString(maxPropSize)
	sectionsTotal := len(config.Sections())

	printSeparator()

	fmtc.Println("{*}const{!} (")

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

	fmt.Println(")")

	printSeparator()
}

// renderUnitedConfig renders united config code
func renderUnitedConfig(config *knf.Config) {
	fmtc.Println(`{s-}// addExtraOptions adds extra options{!}`)
	fmtc.Println(`{*}func{!} {b}addExtraOptions{!}({&}m{!} {*}options.Map{!}) {`)

	for _, section := range config.Sections() {
		for _, prop := range config.Props(section) {
			fmtc.Printf(
				tabSymbol+"m.{r*}Set{!}({r*}knfu.O{!}(%s), &options.{*}V{!}{s}{}{!})\n",
				formatConstName(section, prop),
			)
		}
	}

	fmt.Println("}\n")

	fmtc.Println(`{s-}// combineConfigs combines knf configuration with options and environment variables{!}`)
	fmtc.Println(`{*}func{!} {b}combineConfigs{!}() {`)
	fmtc.Println(tabSymbol + "knfu.{r*}Combine{!}(")

	for _, section := range config.Sections() {
		for _, prop := range config.Props(section) {
			fmtc.Printf(
				tabSymbol+tabSymbol+"knfu.{r*}Simple{!}(%s),\n",
				formatConstName(section, prop),
			)
		}
	}

	fmt.Println(tabSymbol + ")")
	fmt.Println("}")

	printSeparator()
}

// formatConstName returns const name
func formatConstName(section, prop string) string {
	fs := strings.ToUpper(section)
	fp := strings.ToUpper(prop)

	fs = strings.ReplaceAll(fs, "-", "_")
	fp = strings.ReplaceAll(fp, "-", "_")

	return fs + "_" + fp
}

// printSeparator prints separator
func printSeparator() {
	fmtc.Printf("\n{s-}// %s //{!}\n\n", strings.Repeat("/", 72))
}

// getFormatString returns format string
func getFormatString(maxSize int) string {
	return tabSymbol + "%-" + strconv.Itoa(maxSize) + "s = {y}\"%s:%s\"{!}\n"
}

// ////////////////////////////////////////////////////////////////////////////////// //

// printCompletion prints completion for given shell
func printCompletion() int {
	info := genUsage()

	switch options.GetS(OPT_COMPLETION) {
	case "bash":
		fmt.Print(bash.Generate(info, APP))
	case "fish":
		fmt.Print(fish.Generate(info, APP))
	case "zsh":
		fmt.Print(zsh.Generate(info, optMap, APP))
	default:
		return 1
	}

	return 0
}

// printMan prints man page
func printMan() {
	fmt.Println(man.Generate(genUsage(), genAbout("")))
}

// genUsage generates usage info
func genUsage() *usage.Info {
	info := usage.NewInfo("", "file")

	info.AddOption(OPT_SEPARATORS, "Add new lines between sections")
	info.AddOption(OPT_UNITED, "Generate code for united configuration")
	info.AddOption(OPT_NO_COLOR, "Disable colors in output")
	info.AddOption(OPT_HELP, "Show this help message")
	info.AddOption(OPT_VER, "Show version")

	info.AddExample("app.knf", "Generate copy-paste code for app.knf")

	return info
}

// genAbout generates info about version
func genAbout(gitRev string) *usage.About {
	about := &usage.About{
		App:     APP,
		Version: VER,
		Desc:    DESC,
		Year:    2006,
		Owner:   "ESSENTIAL KAOS",

		AppNameColorTag: colorTagApp,
		VersionColorTag: colorTagVer,
		DescSeparator:   "{s}—{!}",

		License: "Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>",
	}

	if gitRev != "" {
		about.Build = "git:" + gitRev
		about.UpdateChecker = usage.UpdateChecker{"essentialkaos/knfgen", update.GitHubChecker}
	}

	return about
}
