package filehandler

import (
	"github.com/cdvelop/input"
	"github.com/cdvelop/model"
	"github.com/cdvelop/object"
	"github.com/cdvelop/unixid"
)

var f *FileHandler

// optional root_folder default: "app_files"
func Add(l model.Logger, db model.DataBaseAdapter, hdd model.FileDiskRW, root_folder ...string) (*FileHandler, error) {
	var out_err error
	if f == nil {

		if hdd == nil {
			return nil, model.Error("error filehandler FileDiskRW nil")
		}

		inputs := []*model.Input{
			unixid.InputPK(),
			input.TextNumCode(),
			input.TextNum(),
			input.Text(),
		}

		module := &model.Module{
			ModuleName: "filehandler",
			Title:      "Manejador de Archivos",
			Objects:    []*model.Object{},
			Inputs:     inputs,
		}

		table := &File{}

		object, err := object.BuildObjectFromStruct(true, table, module)
		if err != nil {
			return nil, err
		}

		f = &FileHandler{
			Object:          object,
			table:           table,
			Logger:          l,
			DataBaseAdapter: db,
			FileDiskRW:      hdd,
			root_folder:     "app_files",
			file_settings:   map[string]*FileSetting{},
		}
		for _, v := range root_folder {
			if v != "" {
				f.root_folder = v
			}
		}

		if !f.RunOnClientDB() { // verificamos la base de datos solo si estamos en el servidor
			f.CreateTablesInDB([]*model.Object{object}, func(err error) {
				if err != nil {
					out_err = err
					// fmt.Println("ESTAMOS EN SERVIDOR CREAMOS TABLA ", f.Object.Table, " EN DB CON ERROR", err)
				}
			})
		}
	}

	return f, out_err
}
