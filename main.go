package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"
	"unsafe"
)

var (
	runConfig = new(mainConfig)
	build     string
	commands  [][]string
)

func init() {
	addFlag(&runConfig.clone, []string{"-clone", "-c"}, false, "Clone project (optional)")
	addFlag(&runConfig.debugMode, []string{"-verbose", "-v", "--verbose"}, false, "Verbose Mode")
	addFlag(&runConfig.version, []string{"-version", "--version"}, false, "Print version and exit")

	flag.Usage = printUsage
	flag.Parse()
}

func main() {
	urls := flag.Args()

	if runConfig.version {
		printVersion()
		return
	}

	if runConfig.debugMode {
		log.Printf("config = %+v", runConfig)
		log.Printf("url = %s", urls)
	}
	if len(urls) == 0 {
		fmt.Printf("missing url(s)\n")
		printUsage()
		return
	}

	var f []string
	for _, v := range urls {
		var err error
		if strings.HasPrefix(v, "https://") || strings.HasPrefix(v, "http://") {
			// Download Mode
			err = download(v)
		} else {
			f = append(f, v)
		}
		if err != nil {
			fmt.Printf("Error: %v", err)
		}
	}
}

func printUsage() {
	fmt.Printf("\nUsage:\n\n  %s [options] file(s)/url(s)\n\n", os.Args[0])
	fmt.Printf("Options:\n\n")
	for _, val := range commands {
		s := fmt.Sprintf("  %s %s", val[0], val[1])
		block := strings.Repeat(" ", 30-len(s))
		fmt.Printf("%s%s%s\n", s, block, val[2])
	}
	fmt.Printf("\n")
}

func printVersion() {
	version := fmt.Sprintf("\ncowTransfer-uploader\n"+
		"Source: https://github.com/Mikubill/cowtransfer-uploader\n"+
		"Build: %s\n", build)
	fmt.Println(version)
}

func addFlag(p interface{}, cmd []string, val interface{}, usage string) {
	s := []string{strings.Join(cmd[1:], ", "), "", usage}
	ptr := unsafe.Pointer(reflect.ValueOf(p).Pointer())
	for _, item := range cmd {
		switch val := val.(type) {
		case int:
			s[1] = "int"
			flag.IntVar((*int)(ptr), item[1:], val, usage)
		case string:
			s[1] = "string"
			flag.StringVar((*string)(ptr), item[1:], val, usage)
		case bool:
			flag.BoolVar((*bool)(ptr), item[1:], val, usage)
		}
	}
	commands = append(commands, s)
}
