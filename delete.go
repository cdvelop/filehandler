package filehandler

import (
	"github.com/cdvelop/model"
)

func (f FileHandler) Delete(u *model.User, params ...map[string]string) (err string) {

	// fmt.Println("parámetros Delete recibidos:", params)
	recover_data, err := f.ReadObjectsInDB(f.Table, params...)
	if err != "" {
		return "error Delete al recuperar data anterior " + err
	}
	// fmt.Println("DATA RECOBRADA:", recover_data)

	err = f.DeleteObjectsInDB(f.Table, params...)
	if err != "" {
		return err
	}

	if len(recover_data) == 0 {
		return "el archivo ya no existe en la base de datos"
	}

	for i, data := range recover_data {

		file_path, _ := f.BuildFilePath(data)

		// Borrar archivos desde hdd
		err := f.FileDelete(file_path)
		if err != "" {
			return "archivo " + data[f.table.Field_name] + " fue eliminado de la db pero no del hdd " + err
		}

		// agregar campos a la data de entrada para retornarla
		for k, v := range data {
			params[i][k] = v
		}

	}

	// fmt.Println("DATA RECOBRADA DESPUÉS DE BORRAR: ", recover_data)

	return ""
}
