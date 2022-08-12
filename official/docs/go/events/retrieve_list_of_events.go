package example

import (
    "os"

    "github.com/EasyPost/easypost-go/v2"
)

func main() {
	apiKey := os.Getenv("EASYPOST_API_KEY")
    client := easypost.New(apiKey)

	events, err := client.ListEvents(
        &easypost.ListOptions{
            PageSize: 5,
        },
    )
	
	fmt.Println(events)
}
