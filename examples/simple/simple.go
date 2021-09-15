package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/KonstantinGasser/humanci"
)

func main() {
	cli := humanci.New()

	cli.RootNOP("search").
		WithRegex(func(meta humanci.MetaData) error {
			switch meta.LastToken() {
			case "google":
				meta.Value("kind", 0)
			case "simple.go":
				meta.Value("kind", 1)
			}
			return nil
		}, "google", "simple.go").
		WithNOP("for").
		WithRegex(func(meta humanci.MetaData) error {

			switch meta.Return("kind").(int) {
			case 0:
				resp, err := http.Get("https://www.google.com?q=" + meta.LastToken())
				if err != nil {
					log.Fatal(err)
				}
				b, err := io.ReadAll(resp.Body)
				if err != nil {
					log.Fatal(err)
				}
				defer resp.Body.Close()
				fmt.Println("Here are the google results for " + meta.LastToken())
				fmt.Println(string(b))
			case 1:
				f, err := os.Open("simple.go")
				if err != nil {
					log.Fatal(err)
				}
				defer f.Close()

				scanner := bufio.NewScanner(f)
				var i int
				for scanner.Scan() {
					text := scanner.Text()
					if strings.Contains(text, meta.LastToken()) {
						for j := 0; j < 5; j++ {
							for scanner.Scan() {
								fmt.Println(scanner.Text())
								i++
								break
							}
						}
						break
					}

				}

				if err := scanner.Err(); err != nil {
					log.Fatal(err)
				}

			}
			return nil
		}, "[a-z]+")

	cli.RootNOP("what").
		WithRegex(func(meta humanci.MetaData) error {
			fmt.Println("it is: ", time.Now().Format("15:04:00"))
			return nil
		}, "time")

	if err := cli.Execute(); err != nil {
		log.Fatal(err)
	}

}
