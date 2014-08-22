package main

import (
	"github.com/ghthor/gospec"
	. "github.com/ghthor/gospec"
	"testing"
)

func TestUnitSpecs(t *testing.T) {
	r := gospec.NewRunner()

	r.AddSpec(DescribeGenPermaLink)

	gospec.MainGoTest(r, t)
}

func DescribeGenPermaLink(c gospec.Context) {
	c.Specify("a permalink will be generated", func() {
		cases := []struct {
			url, permalink string
		}{{
			"http://golang.org/src/pkg/bytes/bytes.go?s=8002:8039#L326",
			"https://code.google.com/p/go/source/browse/src/pkg/bytes/bytes.go?name=go1.3.1#326",
		}, {
			"http://golang.org/src/pkg/bufio/scan.go?s=1184:1667#L20",
			"https://code.google.com/p/go/source/browse/src/pkg/bufio/scan.go?name=go1.3.1#20",
		}}

		for _, ca := range cases {
			link, err := genPermalink(ca.url, "go1.3.1")
			c.Assume(err, IsNil)
			c.Expect(link, Equals, ca.permalink)
		}
	})
}
