package filehandler

import "github.com/cdvelop/model"

type FileSettingAdapter interface {
	GetFileSettings(object_name string) (*FileSetting, error)
}

type FolderPathAdapter interface {
	BuildFilePath(data map[string]string) (file_path, area string)
	UploadFolderPath(data map[string]string) string
}

func (f *FileHandler) AddNewFileSetting(o *model.Object, fs *FileSetting) {
	f.file_settings[o.ObjectName] = fs
}

func (f FileHandler) GetFileSettings(object_name string) (s *FileSetting, err string) {
	if config, exist := f.file_settings[object_name]; exist {
		return config, ""
	}

	return nil, "configuración archivo: " + object_name + " no encontrada"
}

func (f FileHandler) BuildFilePath(data map[string]string) (file_path, area string) {
	return f.UploadFolderPath(data) + "/" + data[f.table.Id_file] + data[f.table.Extension], data[f.table.File_area]
}

func (f FileHandler) UploadFolderPath(data map[string]string) string {
	return f.root_folder + "/" + data[f.table.Module_name] + "/" + data[f.table.Field_name] + "/" + data[f.table.Object_id]
}
