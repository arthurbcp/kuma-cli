package main

import "github.com/arthurbcp/kuma-cli/cmd"

func main() {
	test :
	map[string]interface{} {
		"items": map[string]interface{} {
	   		"subItems": []string { "item1", "item2", "item3"}
		} 
	}
	cmd.Execute()
}
