/*
 * Text spanners for GOPL
 * Copyright 2024 John Douglas Pritchard, Syntelos
 */
package span
/*
 * Byte string concatenation operator.
 */
func Cat(a, b []byte) (c []byte) {
	var len_a int = len(a)
	if 0 != len_a {
		var len_b int = len(b)
		if 0 != len_b {
			var len_c int = (len_a+len_b)
			c = make([]byte,len_c)
			var tgt, src int = 0, 0
			for tgt = 0; tgt < len_c; tgt++ {

				if tgt < len_a {
					src = tgt

					c[tgt] = a[src]

				} else {
					src = (tgt-len_a)

					c[tgt] = b[src]
				}
			}
		} else {
			c = a
		}
	} else {
		var len_b int = len(b)
		if 0 != len_b {
			c = b
		}
	}
	return c
}
