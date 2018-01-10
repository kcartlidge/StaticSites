package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
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
	di := flag.String("sites", ".", "folder containing sites")
	pt := flag.Int("port", 8000, "port to serve sites on")
	si := flag.String("local", "", "optional site to serve as localhost")
	flag.Usage()
	fmt.Println()
	flag.Parse()

	// Create a new server.
	s, err := NewServer(*pt)
	if err != nil {
		log.Fatalln(err)
	}

	// Scan the working folder and treat subfolders as sites.
	dirs, err := ioutil.ReadDir(*di)
	if err != nil {
		log.Fatalln(err)
	}

	// Register and serve the sites.
	fmt.Println("SITES:")
	for _, fi := range dirs {
		nm := fi.Name()
		if fi.IsDir() && nm != "letsencrypt" {
			fo := filepath.Join(*di, nm)
			fmt.Println(" ", nm, " =>", fo)
			s.AddSite(nm, fo)
			if nm == *si {
				fmt.Println(" ", nm, " => localhost:"+strconv.Itoa(*pt))
				s.AddSite("localhost", fo)
			}
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
