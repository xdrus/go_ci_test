package main

import (
	"fmt"
        "golang.org/x/text/language"
)

func second() {
	ru := language.Make("ru")
	fmt.Println(ru.Region())
}
