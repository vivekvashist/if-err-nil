package main

import (
	"fmt"
	"reflect"
	"sort"
)

func main() {
	hellos := map[string]string{
		"English":  "Hello",
		"Mandarin": "您好",
		"Hindi":    "नमस्कार",
	}

	fmt.Printf("%#v\n", hellos)
	fmt.Printf("%#T\n", hellos)                // same as reflect.TypeOf => map[string]string
	fmt.Printf("%s\n", reflect.TypeOf(hellos)) // => map[string]string

	// langs := make([]string, 0, 0) // not nil slice
	var langs []string // nil slice , lenght 0

	for lang, hello := range hellos {
		fmt.Printf("%s %s World \n", lang, hello)
		langs = append(langs, lang)
	}

	sort.Strings(langs)
	fmt.Printf("%v = %d\n", langs, len(langs))
}
