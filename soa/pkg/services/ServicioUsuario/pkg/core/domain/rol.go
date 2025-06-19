package UserModels

type RolDB struct {
	ID      int `db:"id_rol"`
	DataRol Rol
}

type Rol struct {
	NombreRol string `db:"nombrerol"`
}
