Text spanners for GOPL

  Point, group, and class sequence operators.


Group

  A group span is exemplified by the XML code.

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


References

  [XML] https://www.w3.org/TR/REC-xml/REC-xml-20081126.xml
  [URI] https://datatracker.ietf.org/doc/rfc3986/

