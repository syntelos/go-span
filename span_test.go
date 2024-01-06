/*
 * Text spanners for GOPL
 * Copyright 2024 John Douglas Pritchard, Syntelos
 */
package span

import (
	"fmt"
	"testing"
)

func TestXMLP(t *testing.T){
	var xml []byte = []byte("<svg><rect x=\"0\" y=\"0\" width=\"10\" height=\"10\" fill=\"#ABC\" /></svg>")
	var ofs, len = 0, len(xml)
	for ofs < len {
		var first, last = ofs, Forward(xml,ofs,len,'<','>')
		if first < last {
			var begin, end int = first, (last+1)
			var span []byte = xml[begin:end]
			fmt.Printf("\t[%03o,%03o]\t%s\n",first,last,string(span))
			ofs = end
		} else {
			ofs += 1
		}
	}

}
