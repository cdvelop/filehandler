package filehandler

import "github.com/cdvelop/maps"

// type FileRegisterAdapter interface {
// 	FileRegisterInDB(t *File) (map[string]string, error)
// }

func (f FileHandler) FileRegisterInDB(t *File) (out map[string]string, err string) {

	out, err = maps.BuildFormString(t)
	if err != "" {
		return nil, err
	}

	err = f.CreateObjectsInDB("file", false, out)
	if err != "" {
		return nil, err
	}

	return
}
