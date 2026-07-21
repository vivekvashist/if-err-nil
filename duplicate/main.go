package main

import "fmt"

func RemoveDuplicates(s []string) []string {
	seen := make(map[string]struct{}, len(s))
	result := make([]string, 0, len(s))
	for _, v := range s {
		if _, ok := seen[v]; !ok {
			seen[v] = struct{}{}
			result = append(result, v)
		}
	}
	return result
}

func main() {
	ips := []string{
		"10.1.1.1",
		"10.1.1.2",
		"10.1.1.1",
		"192.168.3.2",
		"192.168.3.12",
		"192.168.3.26",
		"192.168.3.12",
	}

	visited := []string{"api.example.com", "cdn.example.com", "api.example.com", "www.example.com", "api.example.com"}

	fmt.Printf("%v\n", RemoveDuplicates(ips))
	fmt.Printf("%v\n", RemoveDuplicates(visited))
}
