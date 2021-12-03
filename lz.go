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
	return out
}
