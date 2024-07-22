package main

import (
	"fmt"
	"html/template"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"

	"github.com/BurntSushi/toml"
	"golang.org/x/exp/rand"
)

func main() {
	log.Println(`
 ______  _                  
(____  \| |                 
 ____)  ) | ____  ____ ____ 
|  __  (| |/ _  )/ ___) _  |
| |__)  ) ( (/ /| |  ( ( | |
|______/|_|\____)_|   \_|| |
                     (_____|
Blog generator
--`)
	generateOutput()
}

func generateOutput() {
	// Wipe and recreate the output folder
	os.RemoveAll("./output")
	os.Mkdir("./output", 0775)
	var err error
	// Grab the pages and header images from the index.toml and create an index.html
	if _, err := os.Stat("./content/index.toml"); err != nil {
		log.Fatalln("Couldn't find ./content/index.toml - have you created it?")
	}

	var config BlergConfig
	_, err = toml.DecodeFile("./content/index.toml", &config)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	funcMap := template.FuncMap{
		"replace":      replace,
		"randomEffect": randomEffect,
	}

	// Load our index.tmpl.html
	indexFile := "./templates/index.tmpl.html"
	tmpl := template.Must(template.New("index.tmpl.html").Funcs(funcMap).ParseFiles(indexFile))

	// create a new file
	file, _ := os.Create("./output/index.html")
	defer file.Close()

	// apply the template to the vars map and write the result to file.
	if err := tmpl.Execute(file, config.Post); err != nil {
		panic(err)
	}

	// Also add a random hover effect to each 'hover:scale-105', 'hover:-rotate-1', 'hover:rotate-1'
	// Run through the pages and create the raw HTML
	// Sexy example of customising the AST here
	// https://github.com/gomarkdown/markdown/blob/master/examples/modify_ast.go

	// Run the tailwind binary to create the CSS
	echo := exec.Command("tailwindcss", "-i", "./tailwind.css", "-o", "./output/main.css", "--minify")
	log.Println("Running tailwind from", echo.Path, "...")
	output, err := echo.CombinedOutput()
	if err != nil {
		log.Fatal("Error running tailwindcsscli from path:", string(output), err)
	}

	// Function to generate new markdown and add to index.toml
	// Generate GenAI image from random keywords
	// Generate sitemap.xml
	// inject in the ol OpenGraph stuff too
	// Create the output in friendly format
	log.Println("Your blog is ready in ./output")
}

// HTML template functions
func replace(input, from, to string) string {
	return strings.Replace(input, from, to, -1)
}

func randomEffect() string {
	effects := []string{
		"hover:-rotate-1 hover:scale-105",
		"hover:rotate-1 hover:scale-105",
		"hover:scale-105",
	}
	return effects[rand.Intn(len(effects))]
}

var (
	// Remove parenthesis and things inside
	noparens     = regexp.MustCompile(`\(.*\)`)
	nosquares    = regexp.MustCompile(`\[.*\]`)
	onlyAlphaNum = regexp.MustCompile(`[^a-zA-Z0-9 ]+`)
)

// generateSexySlug
// https://backlinko.com/hub/seo/urls
func generateSexySlug(title string) string {
	// Lower case and
	title = strings.ToLower(title)
	// Remove all apostrophes and commas
	title = strings.ReplaceAll(title, "'", "")
	title = strings.ReplaceAll(title, ",", "-")
	// remove all spaces for hyphens
	title = strings.ReplaceAll(title, " ", "-")

	// Remove parens then brackets
	title = noparens.ReplaceAllString(title, "")
	title = nosquares.ReplaceAllString(title, "")

	// Strip weird shit last and trails last
	title = onlyAlphaNum.ReplaceAllString(title, "-")
	title = strings.Trim(title, "- ")

	return title
}
