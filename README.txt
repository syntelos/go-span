Text spanners for GOPL

  Point, group, and class sequence operators.


Point

  A character code point in the ASCII domain is typically
  synactically relative to a syntactic frame or scope.


Group

  A group span is exemplified by the XML code.  The XML
  specification assures the syntactic uniqueness of its
  principal character codes, requiring that syntactic
  violations are encoded alternatively.

	var xml []byte = []byte(...)
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

  Another group code example is SEXP.

  These are (both) demonstrated in "span_test".


Class

  A character class code is exemplified by the URI code.
  The density of syntactic freedom is relatively maximized
  with respect to the ASCII code.  Conversely, the
  distinction of class is ambiguated in size.  The larger
  the character class membership set, the less freedom
  remains available to any alternative class of distinct
  uniqueness.


References

  [XML] https://www.w3.org/TR/REC-xml/REC-xml-20081126.xml
  [SXP] https://datatracker.ietf.org/doc/draft-rivest-sexp/
  [URI] https://datatracker.ietf.org/doc/rfc3986/
  [ASC] https://datatracker.ietf.org/doc/html/rfc20

