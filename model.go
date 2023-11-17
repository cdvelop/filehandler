package filehandler

import "github.com/cdvelop/model"

type FileHandler struct {
	*model.Object
	table *File

	model.Logger
	model.DataBaseAdapter
	model.FileDiskRW

	root_folder string // carpeta raíz de archivos
	// nombre del api archivo y su configuración
	file_settings map[string]*FileSetting
}

type File struct {
	Id_file     string `Legend:"File Id" Input:"InputPK"`
	Module_name string `Legend:"Modulo" Input:"TextNumCode"`
	Field_name  string `Legend:"Carpeta Campo" Input:"TextNum" PrincipalField:"ok"`
	Object_id   string `Legend:"Carpeta Id Archivos" Input:"InputPK"`
	File_area   string `Legend:"Area archivo" Input:"Text" PrincipalField:"ok" SkipCompletionAllowed:"true"`
	Extension   string `Legend:"Tipo Archivo" Input:"Text" SkipCompletionAllowed:"true"`
	Description string `Legend:"Descripción" Input:"Text" PrincipalField:"ok" SkipCompletionAllowed:"true"`
}

type FileSetting struct {
	AllowedExtensions   []string // exenciones permitidas ej: ".jpg, .png, .jpeg"
	MaximumFilesAllowed int64    // numero máximo de archivos permitidos ej: 4,100 etc
	MaximumKbSize       int64    // tamaño máximo individual kb ej: 100

	maximumFileSize int64 // tamaño máximo de todos los archivos

	ImagenWidth  string // ej: 800
	ImagenHeight string // ej: 600

	//field
	FieldNameWithObjectID string //ej: id_medicalhistory
	ModuleName            string // ej: user,patients
	DescriptiveName       string //ej: endoscopia, pictures documents
	Legend                string //ej: Imágenes,Boletas etc
	DefaultEnableInput    bool   // si se necesita habilitado resetear el campo por defecto falso

	// Source *model.Object // objeto modulo origen
}
