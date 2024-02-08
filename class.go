/*
 * Text spanners for GOPL
 * Copyright 2024 John Douglas Pritchard, Syntelos
 */
package span
/*
 * Character class operator.
 */
type CC func (byte)(bool)
/*
 * Character class spanner.
 */
func Class (src []byte, ofs, len int, clop CC) (spx int) {
	/*
	 * Clamp to relationship
	 */
        spx = -1

        for ; ofs < len; ofs++ {

                if clop(src[ofs]) {

                        spx = ofs
                } else {
                        return spx
                }
        }
        return spx
}
/*
 * XML Attribute character class.
 */
func XA (ch byte) (bool) {
        switch {
        case 'a' <= ch && 'z' >= ch :
                return true
        case 'A' <= ch && 'Z' >= ch :
                return true
        case '0' <= ch && '9' >= ch :
                return true
        case '_' == ch || '-' == ch || '+' == ch || '.' == ch || '=' == ch || ':' == ch :
                return true
        case '/' == ch || '\'' == ch || '"' == ch:
                return true
	case '?' == ch || '%' == ch || '!' == ch || '#' == ch || '$' == ch :
		return true
	case '(' == ch || ')' == ch || '[' == ch || ']' == ch || '*' == ch :
		return true
        default:
                return false
        }
}
/*
 * Generic identifier character class.
 */
func GI (ch byte) (bool) {
        switch {
        case 'a' <= ch && 'z' >= ch :
                return true
        case 'A' <= ch && 'Z' >= ch :
                return true
        case '0' <= ch && '9' >= ch :
                return true
        case '_' == ch || '-' == ch || '+' == ch || '.' == ch :
                return true
        default:
                return false
        }
}
/*
 * XML Identifier character class.
 */
func XI (ch byte) (bool) {
        switch {
        case 'a' <= ch && 'z' >= ch :
                return true
        case 'A' <= ch && 'Z' >= ch :
                return true
        case '0' <= ch && '9' >= ch :
                return true
        case '_' == ch || '-' == ch || '+' == ch || '.' == ch || ':' == ch :
                return true
        default:
                return false
        }
}
/*
 * XML Code character class.
 */
func XC (ch byte) (bool) {
        switch {
	case '<' == ch || '>' == ch:
		return true
	case '?' == ch || '!' == ch:
		return true
        default:
                return false
        }
}
/*
 * Whitespace character class.
 */
func WS (ch byte) (bool) {
        switch {
        case '\r' == ch || '\n' == ch || '\t' == ch || ' ' == ch:
                return true
        default:
                return false
        }
}
