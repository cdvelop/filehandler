package filehandler

import (
	"github.com/cdvelop/model"
	"github.com/cdvelop/object"
)

var f *FileHandler

func Add(root_folder string, l model.Logger, dba model.DataBaseAdapter) (*FileHandler, error) {
	if f == nil {

		table := &fileTable{}

		err := object.New(table)
		if err != nil {
			return nil, err
		}

		f = &FileHandler{
			Logger:          l,
			DataBaseAdapter: dba,
			root_folder:     root_folder,
			file_settings:   map[string]FileSetting{},
			fileTable:       table,
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
