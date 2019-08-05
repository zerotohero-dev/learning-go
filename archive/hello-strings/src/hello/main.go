package main

import (
	"fmt"
	"os"
	s "strings"
)

type point struct {
	x int
	y int
}

func main() {
	atoz := "the quick brown fox jumps over the lazy dog\n"

	fmt.Printf("%s\n", atoz[0:9])
	fmt.Printf("%s\n", atoz[:9])
	fmt.Printf("%s\n", atoz[15:])
	fmt.Printf("%s\n", atoz[15:19])
	fmt.Printf("%s\n", atoz[:])

	for i, r := range atoz {
		fmt.Printf("%d %c\n", i, r)
	}

	fmt.Printf("len %d\n", len(atoz))

	backquotes := `
		this is taken
		completely literally"";''\n
	`

	fmt.Println(backquotes)

	p := fmt.Println

	p("Contains :", s.Contains("test", "es"))
	p("Count    :", s.Count("test", "t"))
	p("HasPrefix:", s.HasPrefix("test", "te"))
	p("HasSuffix:", s.HasSuffix("test", "st"))
	p("Index    :", s.Index("test", "e"))
	p("Join     :", s.Join([]string{"a", "b"}, "-"))
	p("Repeat   :", s.Repeat("a", 5))
	p("Replace  :", s.Replace("FooBar", "o", "0", -1))
	p("Split    :", s.Split("a-b-c-d-e", "-"))
	p("ToLower  :", s.ToLower("BAZINGA"))
	p("ToUpper  :", s.ToUpper("bazinga"))
	p("Len      :", len("Test"))
	p("Char     :", "Hello"[1])

	o := point{3, 4}

	fmt.Printf("%v\n", o)
	fmt.Printf("%+v\n", o)
	fmt.Printf("%#v\n", o)
	fmt.Printf("%T\n", o)
	fmt.Printf("%t\n", false)
	fmt.Printf("%d\n", 14412)
	fmt.Printf("%b\n", 14412)
	fmt.Printf("%c\n", 14412)
	fmt.Printf("%x\n", 14412)
	fmt.Printf("%f\n", 14412.1)
	fmt.Printf("%e\n", 14412.1)
	fmt.Printf("%E\n", 14412.1)
	fmt.Printf("%s\n", "14412.1")
	fmt.Printf("%q\n", 14412)
	fmt.Printf("%p\n", &o)
	fmt.Printf("|%6d|%6d|\n", 12, 345)
	fmt.Printf("|%-6.2f|%-6.2f|\n", 1.2, 3.45)
	fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 3.45)
	fmt.Printf("|%-6s|%-6s|\n", "foo", "b")

	baz := fmt.Sprintf("a %s", "string")

	fmt.Println(baz)

	fmt.Fprintf(os.Stderr, "an %s\n", "error")
}
