package protocol

type TransferFile struct {
	FileName string
}

type FilePart struct {
	FileName    string
	Buffer      []byte
	MoreContent bool
}
