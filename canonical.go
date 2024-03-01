package hexpalette

const (
	byte_HT = 0x09
	byte_LF = 0x0A
	byte_VT = 0x0B
	byte_FF = 0x0C
	byte_CR = 0x0D
	byte_FS = 0x1C
	byte_GS = 0x1D
	byte_RS = 0x1E
	byte_US = 0x1F
)

// Canonical takes a series of color hex codes and returns a canonical color palette hex code.
//
// For example:
//
//	"#4e3806 #f0A501  #EAC016 f07400   4e3809 " -> "4e3806f0a501eac016f074004e3809"
func Canonical(str string) (string, error) {

	var p []byte = make([]byte, 0, len(str))

	var length int = len(str)
	for i:=0; i<length; i++ {
		var b byte = str[i]

		switch {
		case '0' <= b && b <= '9' ||
		     'a' <= b && b <= 'z':

			p = append(p, b)

		case 'A' <= b && b <= 'Z':
			var lowercase byte = b | 0b00100000

			p = append(p, lowercase)

		default:
			switch b {
			case byte_HT, byte_LF, byte_VT, byte_FF, byte_CR, byte_FS, byte_GS, byte_RS, byte_US, ' ', '#':
				// Nothing here.
			default:
				return "", errBadColors
			}
		}
	}

	return string(p), nil
}
