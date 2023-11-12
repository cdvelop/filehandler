package filehandler

import "github.com/cdvelop/model"

type FileHandler struct {
	model.Logger
	model.DataBaseAdapter

	root_folder string // carpeta raíz de archivos
	// nombre del api archivo y su configuración
	file_settings map[string]FileSetting

	*fileTable
}

type fileTable struct {
	Object *model.Object

	Id_file     string `Legend:"File Id" Input:"InputPK"`
	Module_name string `Legend:"Modulo" Input:"TextNumCode"`
	Field_name  string `Legend:"Carpeta Campo" Input:"TextNum" PrincipalField:"ok"`
	Object_id   string `Legend:"Carpeta Id Archivos" Input:"InputPK"`
	File_area   string `Legend:"Area archivo" Input:"Text" PrincipalField:"ok" SkipCompletionAllowed:"true"`
	Extension   string `Legend:"Tipo Archivo" Input:"Text" SkipCompletionAllowed:"true"`
	Description string `Legend:"Descripción" Input:"Text" PrincipalField:"ok" SkipCompletionAllowed:"true"`
}

type FileSetting struct {
	MaximumFilesAllowed int64  // numero máximo de archivos permitidos ej: 4,100 etc
	InputNameWithFiles  string // nombre del campo con los archivos multipart del formulario ej: files
	MaximumFileSize     int64  // tamaño máximo de todos los archivos
	MaximumKbSize       int64  // tamaño máximo individual kb ej: 100
	AllowedExtensions   string // exenciones permitidas ej: ".jpg, .png, .jpeg"

	ImagenWidth  string // ej: 800
	ImagenHeight string // ej: 600

	RootFolder string //ej: static_files default "app_files"
	FileType   string // ej: imagen,video,document,pdf

	//field
	FieldNameWithObjectID string //ej: id_medicalhistory
	Name                  string //ej: endoscopia, pictures
	Legend                string //ej: Imágenes,Boletas etc
	DefaultEnableInput    bool   // si se necesita habilitado resetear el campo por defecto falso

}
