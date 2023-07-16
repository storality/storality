package flags

import "flag"

var Port *int
var Headless *bool

func init() {
	portFlag := flag.Int("port", 3000, "The port to run the server on")
	headlessFlag := flag.Bool("headless", false, "Run in headless mode")
	flag.Parse()
	Port = portFlag
	Headless = headlessFlag
}

func IsPassed(name string) bool {
	passed := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			passed = true
		}
	})
	return passed
}