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
	var xml []byte = []byte(`<?xml version="1.0" standalone="no"?>
<?xml-stylesheet href="https://cdn.jsdelivr.net/gh/aaaakshat/cm-web-fonts@latest/fonts.css" type="text/css"?>
<!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN" "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd">
<svg xmlns="http://www.w3.org/2000/svg" version="1.1" width="500" height="500" viewBox="0 0 500 500">
  <defs>
    <style type="text/css">
      text {
	  font-family: "Computer Modern Serif", serif;
	  font-size: 11;
	  fill: black;
      }
      text.title {
	  font-size: 22;
	  font-weight: bold
      }
    </style>
  </defs>
  <text class="title" x="20" y="34">Hello, world.</text>
</svg>
`)
	var ofs, len = 0, len(xml)
	for ofs < len {
		var first, last = ofs, Forward(xml,ofs,len,'<','>')
		if first < last {
			var begin, end int = first, (last+1)
			var span []byte = xml[begin:end]
			fmt.Printf("\t[%05o,%05o]\t%s\n",first,last,string(span))
			ofs = end
		} else {
			ofs += 1
		}
	}

}
