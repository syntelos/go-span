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
