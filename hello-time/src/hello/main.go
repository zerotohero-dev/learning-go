package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println

	now := time.Now()
	p(now)

	then := time.Date(
		2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	p(then)

	diff := now.Sub(then)
	p(diff)

	p(then.Add(diff))
	p(then.Add(-diff))

	secs := now.Unix()
	nanos := now.UnixNano()

	fmt.Println(time.Unix(secs, 0))
	fmt.Println(time.Unix(0, nanos))

	t := time.Now()

	p(t.Format(time.RFC3339))

	t1, _ := time.Parse(
		time.RFC3339,
		"2012-11-01T22:08:41+00:00")
	p(t1)
	p("----")
	p(t.Format("3:04PM"))
	p(t.Format("Mon Jan _2 15:04:05 2006"))
	p(t.Format("2006-01-02T15:04:05.999999-07:00"))
	form := "3 04 PM"
	t2, _ := time.Parse(form, "8 41 PM")
	p(t2)

}
