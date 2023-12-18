package filehandler

import (
	"fmt"

	"github.com/cdvelop/input"
	"github.com/cdvelop/model"
	"github.com/cdvelop/object"
	"github.com/cdvelop/unixid"
)

var f *FileHandler

// optional root_folder default: "app_files"
func Add(h *model.MainHandler) (out *FileHandler, err string) {
	var out_err string
	if f == nil {

		if h.FileDiskRW == nil {
			return nil, "error filehandler FileDiskRW nil"
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
		if err != "" {
			return nil, err
		}

		// agregamos el objeto al manejador central
		h.MainHandlerAddModules(object.Module)

		f = &FileHandler{
			Object:          object,
			table:           table,
			Logger:          h.Logger,
			DataBaseAdapter: h.DataBaseAdapter,
			FileDiskRW:      h.FileDiskRW,
			root_folder:     "app_files",
			file_settings:   map[string]*FileSetting{},
		}

		if h.FileRootFolder != "" {
			f.root_folder = h.FileRootFolder
		} else {
			h.FileRootFolder = f.root_folder
		}

		if !f.RunOnClientDB() { // verificamos la base de datos solo si estamos en el servidor
			f.CreateTablesInDB([]*model.Object{object}, func(err string) {
				if err != "" {
					out_err = err
					fmt.Println("ESTAMOS CREANDO TABLA ", f.Object.Table, " EN DB ERROR:", err)
				}
			})
		}
	}

	if f.Logger == nil {
		return nil, "filehandler error Logger == nil"
	}
	if f.DataBaseAdapter == nil {
		return nil, "filehandler error DataBaseAdapter == nil"
	}
	if f.FileDiskRW == nil {
		return nil, "filehandler error FileDiskRW == nil"
	}

	return f, out_err
}
