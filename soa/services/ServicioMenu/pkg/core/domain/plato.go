package MenuModels

// CategoriaPlatos representa la categoría de un plato.
type CategoriaPlatos struct {
	IDCategoria int `db:"id_categoria"`
	Info        CategoriaPlatosVariable
}

type CategoriaPlatosVariable struct {
	Nombre string `db:"nombre"`
}

// Plato representa un plato del menú.
type Plato struct {
	IDPlato int `db:"id_plato"`
	Info    PlatoVariable
}

type PlatoVariable struct {
	NombrePlato string  `db:"nombre_plato"`
	Categoria   int     `db:"categoria"` // FK a CategoriaPlatos
	Descripcion string  `db:"descripcion"`
	Precio      float64 `db:"precio"`
	Imagen      string  `db:"imagen"`
	Estado      bool    `db:"estado"`
}
