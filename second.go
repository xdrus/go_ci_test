package second

import (
	"fmt"
        "golang.org/x/text/language"
)

func main() {
	ru := language.Make("ru")
	fmt.Println(ru.Region())
}
