/*
 * Text spanners for GOPL
 * Copyright 2024 John Douglas Pritchard, Syntelos
 */
package span
/*
 * Sequential point sequence operator.
 */
func Last(src []byte, ofs int, len int, ch byte) (spx int) {
	/*
	 * Clamp to last
	 */
	spx = ofs

	for ; 0 < ofs; ofs-- {

		if ch == src[ofs] {
			/*
			 * Found object next
			 */
			return ofs
		}
	}
	return spx
}
/*
 * Sequential point sequence operator.
 */
func First(src []byte, ofs int, len int, ch byte) (spx int) {
	/*
	 * Clamp to first
	 */
	spx = ofs

	for ; ofs < len; ofs++ {

		if ch == src[ofs] {
			/*
			 * Found object next
			 */
			return ofs
		}
	}
	return spx
}
