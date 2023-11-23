package filehandler

import (
	"github.com/cdvelop/model"
)

func (f FileHandler) AddBootFiles(u *model.User, o *model.Object, from_data []map[string]string, out *[]model.Response) (err string) {

	// fmt.Println(" 1 EJECUTANDO AddBootFiles")
	pk_name := o.PrimaryKeyName()

	field_ids := []map[string]string{}
	for _, data := range from_data {
		field_ids = append(field_ids, map[string]string{
			f.table.Object_id: data[pk_name],
			// f.Object_id: data[f.FieldNameWithObjectID],
			"SELECT": f.table.Id_file + "," + f.table.Object_id + "," + f.table.Description,
		})
	}

	our_files, err := f.Read(u, field_ids...)
	if err != "" {
		return err
	}

	if len(our_files) >= 1 {
		*out = append(*out, f.Response(our_files))
	}
	// fmt.Println("ARCHIVOS LEÍDOS:", len(our_files))

	return ""
}
