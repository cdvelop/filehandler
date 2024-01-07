package filehandler

import (
	"github.com/cdvelop/model"
)

func (f FileHandler) Delete(u *model.User, params ...map[string]string) (err string) {
	// fmt.Println("parámetros Delete recibidos:", params)
	recover_data, err := f.ReadSyncDataDB(model.ReadParams{FROM_TABLE: f.Table}, params...)
	if err != "" {
		return "error Delete al recuperar data anterior " + err
	}
	// fmt.Println("DATA RECOBRADA:", recover_data)

	if len(recover_data) == 0 {
		return "el archivo ya no existe en la base de datos"
	}

	err = f.DeleteFileFromHDDandDB(recover_data...)
	if err != "" {
		return err
	}

	// agregar campos recover_data a params para retornarla
	for i, data := range recover_data {
		for k, v := range data {
			params[i][k] = v
		}
	}
	// fmt.Println("DATA params DESPUÉS DE BORRAR: ", params)

	return ""
}

func (f FileHandler) DeleteFileFromHDDandDB(recover_data ...map[string]string) (err string) {

	err = f.DeleteObjectsInDB(f.Table, false, recover_data...)
	if err != "" {
		return err
	}

	for _, data := range recover_data {

		file_path, _ := f.BuildFilePath(data)

		// fmt.Println("file_path ", file_path)
		// Borrar archivos desde hdd
		err := f.FileDelete(file_path)
		if err != "" {
			return "archivo " + data[f.table.Field_name] + " fue eliminado de la db pero no del hdd " + err
		}
		// fmt.Println("FileDelete err", err)
	}

	return
}
