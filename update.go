package filehandler

import (
	"github.com/cdvelop/model"
)

func (f FileHandler) Update(u *model.User, data ...map[string]string) (err string) {

	// fmt.Println("UPDATE DATA:", data)

	return f.UpdateObjectsInDB(f.Table, data...)
}
