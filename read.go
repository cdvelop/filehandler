package filehandler

import (
	"github.com/cdvelop/model"
)

func (f FileHandler) Read(u *model.User, params ...map[string]string) (result []map[string]string, err string) {

	// fmt.Println("file: parámetros lectura recibidos:", params, "tamaño", len(params))

	if len(params) == 0 {
		return nil, "error filehandler Read: no hay parámetros de lectura recibidos"
	}

	// for _, v := range params {
	// 	v["SELECT"] = f.Id_file + ","+ f.Folder_id + "," + f.Description
	// }

	data, err := f.ReadSyncDataDB(model.ReadParams{FROM_TABLE: f.Table}, params...)
	if err != "" {
		return nil, err
	}
	// fmt.Println("data devuelta:", data)

	return data, ""
}

func (f FileHandler) ReadByID(id string) (result []map[string]string, err string) {

	return f.ReadSyncDataDB(model.ReadParams{FROM_TABLE: f.Table}, map[string]string{
		f.table.Id_file: id,
	})
}
