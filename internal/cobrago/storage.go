package cobrago

type RemoteStorage interface {
	List(bucket string) []RemoteFile
}

type SystemWriter interface {
	Print(files []RemoteFile)
}

type RemoteFile struct {
	Key  string
	Size int64
}

func ListFilesFromStorage(rstorage RemoteStorage, bucket string, writer SystemWriter) {
	files := rstorage.List(bucket)
	writer.Print(files)
}

// This is invalid and wil throw an error on the static analyzers
// func ListFilesInStorage(bucket string, storage RemoteStorage, writer SystemWriter) {
// 	files := storage.List(bucket)
// 	writer.Print(files)
// }
