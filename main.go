package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"strings"
)

const VideoUrl = "https://hanime.tv/videos/hentai/"
const ApiUrl = "https://search.htv-services.com"

var search string

func initFlags() {
	flag.StringVar(&search, "search", "", "Search for animes")
	flag.Parse()
}

func main() {
	initFlags()

	a := fiber.AcquireAgent()
	req := a.Request()
	req.Header.SetMethod(fiber.MethodPost)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/102.0.5005.61 Mobile Safari/537.36")
	req.Header.Set("Content-Type", "application/json")
	req.SetRequestURI(ApiUrl)

	payload := Payload{
		SearchText: search,
		Tags:       []string{},
		TagsMode:   "AND",
		Brands:     []string{},
		Blacklist:  []string{},
		OrderBy:    "created_at_unix",
		Ordering:   "desc",
		Page:       0,
	}

	b, err := json.Marshal(payload)
	HandleError(err)

	req.SetBodyString(string(b))

	err = a.Parse()
	HandleError(err)

	code, body, errs := a.Bytes()

	HandleErrors(errs)

	fmt.Println("Searching: " + search)

	if code == 200 {
		response := &Response{}
		err = json.Unmarshal(body, response)
		HandleError(err)

		hits := &[]Hit{}
		json.Unmarshal([]byte(response.Hits), hits)
		output := "\n"
		for i := 0; i < len(*hits); i++ {
			hit := (*hits)[i]
			url := VideoUrl + strings.Replace(strings.ToLower(hit.Name), " ", "-", -1)
			output += fmt.Sprintf(" [X] %s (%s)\n", hit.Name, url)
		}
		fmt.Printf("Found: %s\n", output)
	} else {
		fmt.Println("There was an error while trying to search")
	}

}

func HandleError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func HandleErrors(err []error) {
	if err != nil {
		log.Fatalln(err)
	}
}
