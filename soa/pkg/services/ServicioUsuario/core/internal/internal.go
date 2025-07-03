// paquete de paquete internal contiene tipos y constantes internas utilizadas en el servicio de usuario
// Este paquete no debe ser importado por otros servicios o paquetes externos.
// Este paquete define tipos y constantes internas utilizadas en el servicio de usuario
// y no debe ser importado por otros servicios o paquetes externos.
// Este paquete define tipos y constantes internas utilizadas en el servicio de usuario
// y no debe ser importado por otros servicios o paquetes externos.
package internal

type Filter struct {
	Key   string `json:"key"`
	Value string `json:"value,omitempty"`
}
type StatusCode int

const (
	InProgress StatusCode = iota + 1
	Busy
	Halted
	Error
)
