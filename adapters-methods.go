package filehandler

import "github.com/cdvelop/model"

type FileSettingAdapter interface {
	GetFileSettings(object_name string) (*FileSetting, error)
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

func (h FileHandler) BuildFilePath(data map[string]string) (file_path, area string) {

	return h.UploadFolderPath(&File{
		Module_name: data[f.table.Module_name],
		Field_name:  data[f.table.Field_name],
		Object_id:   data[f.table.Object_id],
	}) + "/" + data[f.table.Id_file] + data[f.table.Extension], data[f.table.File_area]
}

func (h FileHandler) UploadFolderPath(f *File) string {
	return h.root_folder + "/" + f.Module_name + "/" + f.Field_name + "/" + f.Object_id
}
