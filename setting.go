package filehandler

func (f *FileSetting) SetMaximumFileSize() {
	f.maximumFileSize = int64(float64(f.MaximumFilesAllowed*f.MaximumKbSize*1024) * 1.05)
}
func (f FileSetting) GetMaximumFileSize() int64 {
	return f.maximumFileSize
}
