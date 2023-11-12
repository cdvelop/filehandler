package filehandler

import "github.com/cdvelop/model"

func (f *FileHandler) AddNewFileSetting(o *model.Object, fs *FileSetting) {
	f.file_settings[o.Name] = fs
}

func (f *FileHandler) GetFileSettings(object_name string) (FileSetting, error) {
	if config, exist := f.file_settings[object_name]; exist {
		return *config, nil
	}

	return FileSetting{}, model.Error("configuración archivo:", object_name, "no encontrada")
}

func (f *FileHandler) BuildFilePath(data map[string]string) (file_path string) {
	return f.UploadFolderPath(data) + "/" + data[f.Id_file] + data[f.Extension]
}

func (f *FileHandler) UploadFolderPath(data map[string]string) string {
	return f.root_folder + "/" + data[f.Module_name] + "/" + data[f.Field_name] + "/" + data[f.Object_id]
}
