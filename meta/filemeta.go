package meta

// FileMeta: File meta information structure
type FileMeta struct {
	FileSha1 string
	FileName string
	FileSize int64
	Location string
	UploadAt string
}

var fileMetas map[string]FileMeta

func init() {
	fileMetas = make(map[string]FileMeta)
}

// Add/update File metas
func UpdateFileMeta(fmeta FileMeta) {
	fileMetas[fmeta.FileSha1] = fmeta
}

// Get File Meta from sha1
func GetFileMeta(fileSha1 string) FileMeta {
	return fileMetas[fileSha1]
}
