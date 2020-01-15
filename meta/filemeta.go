package meta

import "sort"

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

//Get multiple file meta information
func GetLastFileMetas(count int) []FileMeta {
	var fMetaArray []FileMeta
	for _, v := range fileMetas {
		fMetaArray = append(fMetaArray, v)
	}
	sort.Sort(ByUploadTime(fMetaArray))
	return fMetaArray[0:count]
}
