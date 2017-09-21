package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	simples "github.com/kcartlidge/simples-config"
)

func main() {
	// Capture ctrl-c etc.
	done := make(chan bool)
	HandleShutdown(done)
	defer fmt.Println()

	// Banner text.
	fmt.Println()
	fmt.Println(" ____  _        _   _        ____  _ _")
	fmt.Println("/ ___|| |_ __ _| |_(_) ___  / ___|(_| |_ ___ ___")
	fmt.Println("\\___ \\| __/ _` | __| |/ __| \\___ \\| | __/ _ / __|")
	fmt.Println(" ___) | || (_| | |_| | (__   ___) | | ||  __\\__ \\")
	fmt.Println("|____/ \\__\\__,_|\\__|_|\\___| |____/|_|\\__\\___|___/")
	fmt.Println()
	fmt.Println("https://github.com/kcartlidge/StaticSites")
	fmt.Println()
	fmt.Println()

	// Handle the args.
	ini := flag.String("sites", "sites.ini", "file containing site details")
	pt := flag.Int("port", 8000, "port to serve sites on")
	si := flag.String("local", "", "optional site to serve as localhost")
	flag.Usage()
	fmt.Println()
	flag.Parse()

	// Load the sites config.
	c, err := simples.CreateConfig(*ini)
	if err != nil {
		log.Fatalln(err.Error())
	}

	// Create a new server.
	s, err := NewServer(*pt)
	if err != nil {
		log.Fatalln(err)
	}

	// Register and serve the sites.
	fmt.Println("SITES:")
	for _, e := range c.GetSection("DEFAULT") {
		fmt.Println(" ", e.Key)
		s.AddSite(e.Key, e.Value)
		if e.Key == *si {
			fmt.Println(" ", e.Key, " => localhost:"+strconv.Itoa(*pt))
			s.AddSite("localhost", e.Value)
		}
	}
	go s.Serve()

	// Wait for enter/return.
	go func() {
		bufio.NewReader(os.Stdin).ReadBytes('\n')
		done <- true
	}()

	// Wait until finished.
	<-done
	fmt.Println()
	fmt.Println("Stopped serving")
}
