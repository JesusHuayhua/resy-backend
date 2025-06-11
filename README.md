# Backend - RESI

Estructura general del proyecto:
```lua
ResiBackend #Nombre de la carpeta despues de hacer gitclone
├───database-scripts #Almacena los scripts usados en BD
└───services #Almacena los servicios que usará el frontend
    ├───ServicioMenu #Servicio abocado a las necesidades del menu
    └───ServicioUsuario #Servicio abocado a las necesidades del modulo de usuarios
        ├───dominio #Maneja todas las clases necesarias para manejar el trabajo
        ├───persistencia #Conexion a base de datos
        └───servicio #interfaces para el flujo de información con el frontend
```
