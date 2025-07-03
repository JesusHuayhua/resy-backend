package backBD

import (
	"database/sql"
	"fmt"
	MenuModels "soa/pkg/services/ServicioMenu/core/domain"
	"soa/pkg/services/ServicioMenu/repository"
	repoInterface "soa/pkg/services/ServicioMenu/repository/interfaces"
	"strconv"
	"time"
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
	return s.crud.Insertar(`"ResyDB"."Plato"`, plato)
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
	return s.crud.Actualizar(`"ResyDB"."Plato"`, plato, where, id)
}

func (s *ServicioMenu) EliminarPlato(id int) error {
	return s.crud.Eliminar(`"ResyDB"."Plato"`, "id_plato")
}

func (s *ServicioMenu) SeleccionarPlatos(condicion string, args []interface{}) ([]MenuModels.Plato, error) {
	var platos []MenuModels.Plato
	columnas := []string{"id_plato", "nombre_plato", "categoria", "descripcion", "precio", "imagen", "estado"}
	var rows *sql.Rows
	var err error
	if condicion == "" {
		rows, err = s.crud.Seleccionar(`"ResyDB"."Plato"`, columnas, "", args...)
	} else {
		rows, err = s.crud.Seleccionar(`"ResyDB"."Plato"`, columnas, condicion, args...)
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
	return s.crud.Insertar(`"ResyDB"."CategoriaPlatos"`, cat)
}

func (s *ServicioMenu) ActualizarCategoria(id int, nombre string) error {
	cat := MenuModels.CategoriaPlatosVariable{Nombre: nombre}
	where := "id_categoria = $1"
	return s.crud.Actualizar(`"ResyDB"."CategoriaPlatos"`, cat, where, id)
}

func (s *ServicioMenu) EliminarCategoria(id int) error {
	return s.crud.Eliminar(`"ResyDB"."CategoriaPlatos"`, strconv.Itoa(id))
}

func (s *ServicioMenu) SeleccionarCategorias() ([]MenuModels.CategoriaPlatos, error) {
	var cats []MenuModels.CategoriaPlatos
	columnas := []string{"id_categoria", "nombre"}
	rows, err := s.crud.Seleccionar(`"ResyDB"."CategoriaPlatos"`, columnas, "")
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

func (s *ServicioMenu) InsertarMenuSemanal(idMenu string, fechaInicio, fechaFin string) error {
	// Parse the string dates to time.Time
	const layout = "2006-01-02" // adjust layout as needed
	fechaInicioTime, err := time.Parse(layout, fechaInicio)
	if err != nil {
		return fmt.Errorf("error parsing fechaInicio: %w", err)
	}
	fechaFinTime, err := time.Parse(layout, fechaFin)
	if err != nil {
		return fmt.Errorf("error parsing fechaFin: %w", err)
	}
	menu := MenuModels.MenuSemanalVariable{
		FechaDeInicio: fechaInicioTime,
		FechaFin:      fechaFinTime,
	}
	return s.crud.Insertar(`"ResyDB"."MenuSemanal"`, menu)
}

// Agregar día al menú semanal
func (s *ServicioMenu) InsertarMenudia(idMenu string, diaSemana string) error {
	menudia := MenuModels.MenudiaVariable{
		IDMenu:    idMenu,
		DiaSemana: diaSemana,
	}
	return s.crud.Insertar(`"ResyDB"."Menudia"`, menudia)
}

// Asignar plato a un día del menú
func (s *ServicioMenu) InsertarPlatoEnMenudia(idDia int, idPlato int, cantidad int, disponible bool) error {
	platoDia := MenuModels.PlatosEnMenudiaInsert{
		IDDia:            idDia,
		IDPlato:          idPlato,
		CantidadDelPlato: cantidad,
		DisponibleVenta:  disponible,
	}
	return s.crud.Insertar(`"ResyDB"."PlatosEnMenudia"`, platoDia)
}

// Listar menús semanales
func (s *ServicioMenu) ListarMenusSemanales() ([]map[string]interface{}, error) {
	columnas := []string{"id_menu", "fecha_inicio", "fecha_fin"}
	rows, err := s.crud.Seleccionar(`"ResyDB"."MenuSemanal"`, columnas, "")
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
			"id_menu":      idMenu,
			"fecha_inicio": fechaInicio,
			"fechafin":     fechaFin,
		})
	}
	return menus, nil
}

// Listar días de un menú semanal
func (s *ServicioMenu) ListarDiasDeMenu(idMenu string) ([]map[string]interface{}, error) {
	columnas := []string{"id_dia", "id_menu", "dia_semana"}
	rows, err := s.crud.Seleccionar(`"ResyDB"."Menudia"`, columnas, "id_menu = $1", idMenu)
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
	columnas := []string{"id_dia", "id_plato", "cantidad_plato", "disponible_venta"}
	rows, err := s.crud.Seleccionar(`"ResyDB"."PlatosEnMenudia"`, columnas, "id_dia = $1", idDia)
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
			"id_dia":         idDia,
			"id_plato":       idPlato,
			"cantidad_plato": cantidad,
			"disponible":     disponible,
		})
	}
	return platos, nil
}

// Obtener información completa de un menú semanal
func (s *ServicioMenu) ObtenerMenuSemanalCompleto(idMenu string) (map[string]interface{}, error) {
	result := make(map[string]interface{})

	// 1. Obtener info básica del menú semanal
	columnasMenu := []string{"id_menu", "fecha_inicio", "fecha_fin"}
	rowsMenu, err := s.crud.Seleccionar(`"ResyDB"."MenuSemanal"`, columnasMenu, "id_menu = $1", idMenu)
	if err != nil {
		return nil, err
	}
	defer rowsMenu.Close()
	if !rowsMenu.Next() {
		return nil, fmt.Errorf("no existe el menú semanal con id %s", idMenu)
	}
	var idMenuDB, fechaInicio, fechaFin string
	if err := rowsMenu.Scan(&idMenuDB, &fechaInicio, &fechaFin); err != nil {
		return nil, err
	}
	result["id_menu"] = idMenuDB
	result["fecha_inicio"] = fechaInicio
	result["fecha_fin"] = fechaFin

	// 2. Obtener días del menú semanal
	dias, err := s.ListarDiasDeMenu(idMenu)
	if err != nil {
		return nil, err
	}

	// 3. Para cada día, obtener los platos
	for i, dia := range dias {
		idDia, ok := dia["id_dia"].(int)
		if !ok {
			// Si por alguna razón el id_dia no es int, intentar convertir
			switch v := dia["id_dia"].(type) {
			case int64:
				idDia = int(v)
			case float64:
				idDia = int(v)
			default:
				continue
			}
		}
		platos, err := s.ListarPlatosEnMenudia(idDia)
		if err != nil {
			return nil, err
		}
		dias[i]["platos"] = platos
	}
	result["dias"] = dias

	return result, nil
}
