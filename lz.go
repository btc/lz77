package lz

type Run struct {
	Offset int
	Len    int
	Ch     byte // TODO: NB could be rune?
}

func Encode(buf []byte) []Run {
	var out []Run

	for i := 0; i < len(buf); {

		// a special case to simplify the code in the general case
		// TODO re-evaluate and maybe unify
		is_last_char := i == len(buf)-1
		if is_last_char {
			out = append(out, Run{0, 0, buf[i]})
			break
		}

		found, length := 0, 0

		for offset := 0; offset < i; offset++ {
			if buf[offset] != buf[i] { // no match
				continue
			}

			// found a match!

			found, length = offset, 1

			// extend the length
			for i+length < len(buf) && found+length < i && buf[found+length] == buf[i+length] {
				// TODO make more readable by naming cases
				length++
			}

			break
		}

		if length == 0 {
			out = append(out, Run{0, 0, buf[i]}) // TODO unify the append
			i++
		} else {
			out = append(out, Run{found, length, buf[i+length]})
			i += length + 1
		}
	}
	return out
}

func Decode(buf []Run) []byte {
	var out []byte
	for _, run := range buf {
		out = append(out, out[run.Offset:run.Offset+run.Len]...)
		out = append(out, run.Ch)
	}
	return out
}
