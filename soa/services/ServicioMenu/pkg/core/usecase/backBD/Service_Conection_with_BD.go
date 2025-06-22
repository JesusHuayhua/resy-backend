package backBD

import (
	MenuModels "ServicioMenu/pkg/core/domain"
	"ServicioMenu/pkg/repository"
	repoInterface "ServicioMenu/pkg/repository/interfaces"
	"database/sql"
	"fmt"
)

type ServicioMenu struct {
	crud repoInterface.PlatoRepository
}

func NuevoServicioMenu(db *sql.DB) *ServicioMenu {
	crud := repository.NewPlatoRepository(db)
	return &ServicioMenu{crud: crud}
}

// CRUD PLATO
func (s *ServicioMenu) InsertarPlato(nombre string, categoria int, descripcion string, precio float64, imagen string, estado bool) error {
	plato := MenuModels.PlatoVariable{
		NombrePlato: nombre,
		Categoria:   categoria,
		Descripcion: descripcion,
		Precio:      precio,
		Imagen:      imagen,
		Estado:      estado,
	}
	return s.crud.Insertar(`"Plato"`, plato)
}

func (s *ServicioMenu) ActualizarPlato(id int, nombre string, categoria int, descripcion string, precio float64, imagen string, estado bool) error {
	plato := MenuModels.PlatoVariable{
		NombrePlato: nombre,
		Categoria:   categoria,
		Descripcion: descripcion,
		Precio:      precio,
		Imagen:      imagen,
		Estado:      estado,
	}
	where := "id_plato = $1"
	return s.crud.Actualizar(`"Plato"`, plato, where, id)
}

func (s *ServicioMenu) EliminarPlato(id int) error {
	return s.crud.Eliminar(`"Plato"`, fmt.Sprintf("%d", id))
}

func (s *ServicioMenu) SeleccionarPlatos(condicion string, args []interface{}) ([]MenuModels.Plato, error) {
	var platos []MenuModels.Plato
	columnas := []string{"id_plato", "nombrePlato", "categoria", "descripcion", "precio", "imagen", "estado"}
	var rows *sql.Rows
	var err error
	if condicion == "" {
		rows, err = s.crud.Seleccionar(`"Plato"`, columnas, "", args...)
	} else {
		rows, err = s.crud.Seleccionar(`"Plato"`, columnas, condicion, args...)
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var p MenuModels.Plato
		err := rows.Scan(
			&p.IDPlato,
			&p.Info.NombrePlato,
			&p.Info.Categoria,
			&p.Info.Descripcion,
			&p.Info.Precio,
			&p.Info.Imagen,
			&p.Info.Estado,
		)
		if err != nil {
			return nil, err
		}
		platos = append(platos, p)
	}
	return platos, nil
}

// CRUD CATEGORIA
func (s *ServicioMenu) InsertarCategoria(nombre string) error {
	cat := MenuModels.CategoriaPlatosVariable{Nombre: nombre}
	return s.crud.Insertar(`"CategoriaPlatos"`, cat)
}

func (s *ServicioMenu) ActualizarCategoria(id int, nombre string) error {
	cat := MenuModels.CategoriaPlatosVariable{Nombre: nombre}
	where := "id_categoria = $1"
	return s.crud.Actualizar(`"CategoriaPlatos"`, cat, where, id)
}

func (s *ServicioMenu) EliminarCategoria(id int) error {
	return s.crud.Eliminar(`"CategoriaPlatos"`, fmt.Sprintf("%d", id))
}

func (s *ServicioMenu) SeleccionarCategorias() ([]MenuModels.CategoriaPlatos, error) {
	var cats []MenuModels.CategoriaPlatos
	columnas := []string{"id_categoria", "nombre"}
	rows, err := s.crud.Seleccionar(`"CategoriaPlatos"`, columnas, "")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var c MenuModels.CategoriaPlatos
		err := rows.Scan(&c.IDCategoria, &c.Info.Nombre)
		if err != nil {
			return nil, err
		}
		cats = append(cats, c)
	}
	return cats, nil
}

// CRUD MENU SEMANAL (solo ejemplo de inserción)
func (s *ServicioMenu) InsertarMenuSemanal(idMenu string, fechaInicio, fechaFin string) error {
	menu := struct {
		IDMenu        string `db:"id_menu"`
		FechaDeInicio string `db:"fechadeinicio"`
		FechaFin      string `db:"fechafin"`
	}{
		IDMenu:        idMenu,
		FechaDeInicio: fechaInicio,
		FechaFin:      fechaFin,
	}
	return s.crud.Insertar(`"MenuSemanal"`, menu)
}

// Agregar día al menú semanal
func (s *ServicioMenu) InsertarMenudia(idMenu string, diaSemana string) error {
	menudia := struct {
		IDMenu    string `db:"id_menu"`
		DiaSemana string `db:"dia_semana"`
	}{
		IDMenu:    idMenu,
		DiaSemana: diaSemana,
	}
	return s.crud.Insertar(`"Menudia"`, menudia)
}

// Asignar plato a un día del menú
func (s *ServicioMenu) InsertarPlatoEnMenudia(idDia int, idPlato int, cantidad int, disponible bool) error {
	platoDia := struct {
		IDDia                int  `db:"id_dia"`
		IDPlato              int  `db:"id_plato"`
		CantidadDelPlato     int  `db:"cantidaddelplato"`
		DisponibleParaVender bool `db:"disponibleparavender"`
	}{
		IDDia:                idDia,
		IDPlato:              idPlato,
		CantidadDelPlato:     cantidad,
		DisponibleParaVender: disponible,
	}
	return s.crud.Insertar(`"PlatosEnMenudia"`, platoDia)
}

// Listar menús semanales
func (s *ServicioMenu) ListarMenusSemanales() ([]map[string]interface{}, error) {
	columnas := []string{"id_menu", "fechadeinicio", "fechafin"}
	rows, err := s.crud.Seleccionar(`"MenuSemanal"`, columnas, "")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var menus []map[string]interface{}
	for rows.Next() {
		var idMenu, fechaInicio, fechaFin string
		if err := rows.Scan(&idMenu, &fechaInicio, &fechaFin); err != nil {
			return nil, err
		}
		menus = append(menus, map[string]interface{}{
			"id_menu":       idMenu,
			"fechadeinicio": fechaInicio,
			"fechafin":      fechaFin,
		})
	}
	return menus, nil
}

// Listar días de un menú semanal
func (s *ServicioMenu) ListarDiasDeMenu(idMenu string) ([]map[string]interface{}, error) {
	columnas := []string{"id_dia", "id_menu", "dia_semana"}
	rows, err := s.crud.Seleccionar(`"Menudia"`, columnas, "id_menu = $1", idMenu)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var dias []map[string]interface{}
	for rows.Next() {
		var idDia int
		var idMenu, diaSemana string
		if err := rows.Scan(&idDia, &idMenu, &diaSemana); err != nil {
			return nil, err
		}
		dias = append(dias, map[string]interface{}{
			"id_dia":     idDia,
			"id_menu":    idMenu,
			"dia_semana": diaSemana,
		})
	}
	return dias, nil
}

// Listar platos de un día del menú
func (s *ServicioMenu) ListarPlatosEnMenudia(idDia int) ([]map[string]interface{}, error) {
	columnas := []string{"id_dia", "id_plato", "cantidaddelplato", "disponibleparavender"}
	rows, err := s.crud.Seleccionar(`"PlatosEnMenudia"`, columnas, "id_dia = $1", idDia)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var platos []map[string]interface{}
	for rows.Next() {
		var idDia, idPlato, cantidad int
		var disponible bool
		if err := rows.Scan(&idDia, &idPlato, &cantidad, &disponible); err != nil {
			return nil, err
		}
		platos = append(platos, map[string]interface{}{
			"id_dia":     idDia,
			"id_plato":   idPlato,
			"cantidad":   cantidad,
			"disponible": disponible,
		})
	}
	return platos, nil
}
