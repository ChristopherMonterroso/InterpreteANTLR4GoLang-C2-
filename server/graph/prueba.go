package graph

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func CreateGraph(input string) string {
	// replace the " with \"
	input = strings.ReplaceAll(input, "\"", "\\\"")

	url := "http://lab.antlr.org/parse/"

	// payload
	payload := []byte(`{"grammar": "Control",
"input": "` + input + `", 
"lexgrammar": "Control",
"start": "prog"}`)

	// request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))

	if err != nil {
		fmt.Println("Error creating request:", err)
		return ""
	}
	req.Header.Set("Content-Type", "application/json") // Set the appropriate content type

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return ""
	}
	defer resp.Body.Close()

	fmt.Println("Response Status:", resp.Status)

	// parse the response body to json
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading body:", err)
		return ""
	}

	// create a map to store the json
	var data map[string]interface{}

	// unmarshal the json
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println("Error unmarshalling json:", err)
		return ""
	}

	// fmt.Println("Response Body:", data)

	result := data["result"].(map[string]interface{})

	// fmt.Println("Response Body:", result["svgtree"])

	// create the file
	err = os.WriteFile("cst.svg", []byte(result["svgtree"].(string)), 0644)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return ""
	}

	fmt.Println("File created successfully")

	return result["svgtree"].(string)

}
