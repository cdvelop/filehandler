package filehandler

import (
	"github.com/cdvelop/model"
)

func (f FileHandler) Read(u *model.User, params ...map[string]string) ([]map[string]string, error) {

	// fmt.Println("file: parámetros lectura recibidos:", params, "tamaño", len(params))

	if len(params) == 0 {
		return nil, model.Error("error filehandler Read: no hay parámetros de lectura recibidos")
	}

	// for _, v := range params {
	// 	v["SELECT"] = f.Id_file + ","+ f.Folder_id + "," + f.Description
	// }

	data, err := f.ReadObjectsInDB(f.Table, params...)
	if err != nil {
		return nil, err
	}
	// fmt.Println("data devuelta:", data)

	return data, nil
}

func (f FileHandler) ReadByID(id string) ([]map[string]string, error) {

	data, err := f.ReadObjectsInDB(f.Table, map[string]string{
		f.table.Id_file: id,
	})
	if err != nil {
		return nil, err
	}

	return data, nil
}
