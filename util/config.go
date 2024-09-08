package util

import (
	"fmt"
	"regexp"

	"github.com/pterm/pterm"
	"github.com/pterm/pterm/putils"
)

type Config struct {
	Addr            string
	Port            int
	DnsAddr         string
	DnsPort         int
	DnsIPv4Only     bool
	EnableDoh       bool
	Debug           bool
	Banner          bool
	SystemProxy     bool
	Timeout         int
	WindowSize      int
	AllowedPatterns []*regexp.Regexp
}

var config *Config

func GetConfig() *Config {
	if config == nil {
		config = new(Config)
	}
	return config
}

func (c *Config) Load(args *Args) {
	c.Addr = args.Addr
	c.Port = int(args.Port)
	c.DnsAddr = args.DnsAddr
	c.DnsPort = int(args.DnsPort)
	c.DnsIPv4Only = args.DnsIPv4Only
	c.Debug = args.Debug
	c.EnableDoh = args.EnableDoh
	c.Banner = args.Banner
	c.SystemProxy = args.SystemProxy
	c.Timeout = int(args.Timeout)
	c.AllowedPatterns = parseAllowedPattern(args.AllowedPattern)
	c.WindowSize = int(args.WindowSize)
}

func parseAllowedPattern(patterns StringArray) []*regexp.Regexp {
	var allowedPatterns []*regexp.Regexp

	for _, pattern := range patterns {
		allowedPatterns = append(allowedPatterns, regexp.MustCompile(pattern))
	}

	return allowedPatterns
}

func PrintColoredBanner() {
	cyan := putils.LettersFromStringWithStyle("Spoof", pterm.NewStyle(pterm.FgCyan))
	purple := putils.LettersFromStringWithStyle("DPI", pterm.NewStyle(pterm.FgLightMagenta))
	pterm.DefaultBigText.WithLetters(cyan, purple).Render()

	pterm.DefaultBulletList.WithItems([]pterm.BulletListItem{
		{Level: 0, Text: "ADDR    : " + fmt.Sprint(config.Addr)},
		{Level: 0, Text: "PORT    : " + fmt.Sprint(config.Port)},
		{Level: 0, Text: "DNS     : " + fmt.Sprint(config.DnsAddr)},
		{Level: 0, Text: "DEBUG   : " + fmt.Sprint(config.Debug)},
	}).Render()

	pterm.DefaultBasicText.Println("Press 'CTRL + c' to quit")
}

func PrintSimpleInfo() {
	fmt.Println("")
	fmt.Println("- ADDR    : ", config.Addr)
	fmt.Println("- PORT    : ", config.Port)
	fmt.Println("- DNS     : ", config.DnsAddr)
	fmt.Println("- DEBUG   : ", config.Debug)
	fmt.Println("")
	fmt.Println("Press 'CTRL + c to quit'")
	fmt.Println("")
}
