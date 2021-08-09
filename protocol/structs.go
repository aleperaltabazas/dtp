package protocol

type SendFile struct {
	FileName string
}

type FilePart struct {
	FileName    string
	Buffer      []byte
	MoreContent bool
}
