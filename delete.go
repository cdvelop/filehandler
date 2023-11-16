package filehandler

import (
	"github.com/cdvelop/model"
)

func (f FileHandler) Delete(u *model.User, params ...map[string]string) ([]map[string]string, error) {

	// fmt.Println("parámetros Delete recibidos:", params)
	recover_data, err := f.ReadObjectsInDB(f.object.Table, params...)
	if err != nil {
		return nil, err
	}

	err = f.DeleteObjectsInDB(f.object.Table, params...)
	if err != nil {
		return nil, err
	}

	if len(recover_data) == 0 {
		return nil, model.Error("el archivo ya no existe en la base de datos")
	}

	for _, data := range recover_data {

		file_path, _ := f.BuildFilePath(data)

		// Borrar archivos desde hdd
		err := f.FileDelete(file_path)
		if err != nil {
			return nil, model.Error("archivo", data[f.table.Field_name], "fue eliminado de la db pero no del hdd", err)
		}
	}

	// fmt.Println("DATA RECOBRADA DESPUÉS DE BORRAR: ", recover_data)

	return recover_data, nil
}
