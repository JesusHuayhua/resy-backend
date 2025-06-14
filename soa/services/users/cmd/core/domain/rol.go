package UserModels

type RolDB struct {
	ID        int    `json:"id_rol"`
	NombreRol string `json:"nombrerol"`
}

type Rol struct {
	NombreRol string `db:"nombrerol"`
}
