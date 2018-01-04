package main


import (
	"fmt"
	s "strings"
)

var p = fmt.Println

func main() {
	p("Contains:",s.Contains("test","es"))
	p("Count:",s.Contains("test","es"))
	p("HasPrefix:",s.HasPrefix("test","te"))
	p("Index:",s.Index("test","e"))
	p("Join:      ", s.Join([]string{"a", "b"}, "-"))
	p("ToLower:",s.ToLower("TEST"))
	p("Len:",len("hello"))
}



