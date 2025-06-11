# Backend - RESI

Estructura general del proyecto:
```lua
ResiBackend #Nombre de la carpeta despues de hacer gitclon
│   .gitignore
│   README.md
│
├───database-scripts #Almacena los scripts usados en BD
│       CreaciónDeLaBaseDeDatos.sql
│       Insert para verificar usuario y roles.sql
│       Notaimportante.txt
│
└───services #Almacena los servicios que usará el frontend
    ├───ServicioMenu #Almacena los scripts usados en BD
    └───ServicioUsuario #Servicio abocado a las necesidades del modulo de usuarios
        │   go.mod
        │   go.sum
        │   main.go
        │
        ├───dominio #Maneja todas las clases necesarias para manejar el trabajo
        │       rol.go
        │       usuario.go
        │
        ├───persistencia #Conexion a base de datos
        │       rolBD.go
        │       usuarioBD.go
        │
        └───servicio #interfaces para el flujo de información con el frontend
```
