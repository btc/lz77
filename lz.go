package lz

type Run struct {
	Offset int
	Len    int
	Ch     byte // TODO: NB could be rune?
}

func Encode(buf []byte) []Run {
	return nil
}

func Decode(buf []Run) []byte {
	var out []byte
	for _, run := range buf {
		out = append(out, out[run.Offset:run.Offset+run.Len]...)
		out = append(out, run.Ch)
	}
	return out
}
