package filehandler

import "github.com/cdvelop/maps"

type FileRegisterAdapter interface {
	FileRegisterInDB(t *File) (map[string]string, error)
}

func (f FileHandler) FileRegisterInDB(t *File) (map[string]string, error) {

	form_data, err := maps.BuildFormString(t)
	if err != nil {
		return nil, err
	}

	err = f.CreateObjectsInDB("file", false, form_data)
	if err != nil {
		return nil, err
	}

	return map[string]string{
		f.table.Id_file:     t.Id_file,
		f.table.Description: form_data[f.table.Description],
	}, nil
}
