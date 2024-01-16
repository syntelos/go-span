/*
 * Text spanners for GOPL
 * Copyright 2024 John Douglas Pritchard, Syntelos
 */
package span
/*
 * Sequential group iterator.
 */
func Forward(src []byte, ofs int, len int, lhs byte, rhs byte) int {
	if ofs < len {
		/*
		 * Clamp to last
		 */
		var y int = (len-1)
		if ofs < y {
			if lhs == src[ofs] {

				return First(src,ofs,len,rhs)

			} else {

				return First(src,ofs,len,lhs)-1
			}
			return y
		}
	}
	return -1
}
/*
 * Sequential group iterator.
 */
func Reverse(src []byte, ofs int, len int, lhs byte, rhs byte) int {
	/*
	 * Clamp to last
	 */
	if 0 < ofs {

		if rhs == src[ofs] {

			return Last(src,ofs,len,lhs)

		} else {

			return Last(src,ofs,len,rhs)-1
		}
		return ofs
	}
	return -1
}
/*
 * Group code resolver.
 */
func GroupRhs(lhs byte) byte {
	switch lhs {
	case '(':
		return ')'
	case '<':
		return '>'
	case '[':
		return ']'
	case '{':
		return '}'
	default:
		return 0
	}
}
/*
 * Nested group operator.
 */
func Group(src []byte, ofs int, len int, lhs byte) int {

	var rhs byte = GroupRhs(lhs)
	if 0 != rhs && lhs == src[ofs] {
		var index, depth int = ofs, 0

		for ; index < len; index++ {
			var ch byte = src[index]

			if lhs == ch {

				depth += 1

			} else if rhs == ch {

				if 1 < depth {

					depth -= 1

				} else if 1 == depth {

					return index
				}
			}
		}
	}
	return -1
}
