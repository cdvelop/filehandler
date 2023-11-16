package filehandler

func (f *FileSetting) SetMaximumFileSize() {
	f.maximumFileSize = int64(float64(f.MaximumFilesAllowed*f.MaximumKbSize*1024) * 1.05)
}
func (f FileSetting) GetMaximumFileSize() int64 {
	return f.maximumFileSize
}

func (f *FileSetting) SetAllowedExtensions() {

	switch f.FileType {

	case "imagen":
		f.Legend = "Imágenes"
		f.allowedExtensions = []string{".jpg", ".png", ".jpeg"}
		f.FileType = "i"

	case "video":
		f.Legend = "Videos"
		f.allowedExtensions = []string{".avi, .mkv, .mp4"}
		f.FileType = "v"

	case "document":
		f.Legend = "Documentos"
		f.allowedExtensions = []string{".doc, .xlsx, .txt"}
		f.FileType = "d"

	case "pdf":
		f.Legend = "Documentos PDF"
		f.allowedExtensions = []string{".pdf"}
		f.FileType = "p"
	}
}

func (f FileSetting) GetAllowedExtensions() []string {
	return f.allowedExtensions
}
