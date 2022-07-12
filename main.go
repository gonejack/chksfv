package main

import (
	"fmt"
	"log"

	"github.com/alecthomas/kong"
	"github.com/mpolden/sfv"
)

func main() {
	var opt struct {
		SFV   string `name:"sfv file" arg:"" help:"path to .sfv file"`
		About bool   `help:"About."`
	}
	kong.Parse(&opt,
		kong.Name("chksfv"),
		kong.Description("This command line do read and check of .sfv file."),
		kong.UsageOnError(),
	)
	if opt.About {
		fmt.Println("Visit https://github.com/gonejack/chksfv")
		return
	}
	s, err := sfv.Read(opt.SFV)
	if err != nil {
		log.Fatal(err)
	}
	ok, err := s.Verify()
	if err != nil {
		log.Fatal(err)
	}
	if ok {
		log.Print("All files are OK!")
	}
	for _, c := range s.Checksums {
		log.Printf("%+v", c)
	}
}
