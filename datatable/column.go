package datatable

import (
	"fmt"

	"github.com/datasweet/datatable/serie"
)

type column struct {
	name   string
	typ    string
	hidden bool
	serie.Serie
}

func main() {
	c := &column{}
	Sum(c)
}

func Sum(s serie.Serie) {
	fmt.Println(s.Error())
}
