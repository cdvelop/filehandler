package filehandler

import (
	"github.com/cdvelop/model"
	"github.com/cdvelop/object"
)

var f *FileHandler

// root_folder default:app_files
func Add(root_folder string, l model.Logger, dba model.DataBaseAdapter) (*FileHandler, error) {
	if f == nil {
		const e = "error en filehandler handler nil:"

		if l == nil {
			return nil, model.Error(e, "Logger")
		}
		if dba == nil {
			return nil, model.Error(e, "DataBaseAdapter")
		}

		table := &fileTable{}

		err := object.New(table)
		if err != nil {
			return nil, err
		}

		f = &FileHandler{
			Logger:          l,
			DataBaseAdapter: dba,
			root_folder:     "app_files",
			file_settings:   map[string]*FileSetting{},
			fileTable:       table,
		}

		if root_folder != "" {
			f.root_folder = root_folder
		}

		if !f.RunOnClientDB() { // verificamos la base de datos solo si estamos en el servidor
			f.CreateTablesInDB([]*model.Object{table.Object}, func(err error) {
				if err != nil {
					l.Log(err)
					return
				}
			})
			// fmt.Println("ESTAMOS EN SERVIDOR CREAMOS TABLA ", f.Object.Table, " EN DB CON ERROR", err)
		}

	}
	return f, nil
}
