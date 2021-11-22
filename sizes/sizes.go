package sizes

import "fmt"

const (
	KB = 1024
	MB = KB * 1024
	GB = MB * 1024
)

func ParseBytes(input string) (uint64, error) {
	var res uint64
	var done bool
	for _, c := range []rune(input) {
		switch c {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			res = res*10 + uint64(c-'0')
		case 'g':
			if done {
				return 0, fmt.Errorf("unrecognized suffix in %q", input)
			}
			res = res * 1024 * 1024 * 1024
			done = true
		case 'm':
			if done {
				return 0, fmt.Errorf("unrecognized suffix in %q", input)
			}
			res = res * 1024 * 1024
			done = true
		case 'k':
			if done {
				return 0, fmt.Errorf("unrecognized suffix in %q", input)
			}
			res = res * 1024
			done = true
		default:
			return 0, fmt.Errorf("unparsable size %q", input)
		}
	}
	return res, nil
}

func Format(size uint64) (uint64, string) {
	if size < 10*KB {
		return size, "B"
	}
	if size < 10*MB {
		return size / KB, "KB"
	}
	if size < 10*GB {
		return size / MB, "MB"
	}
	return size / GB, "GB"
}
