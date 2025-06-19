# Backend - RESI

## Requerimientos
1. **Visual Studio Code** (you need to install Golang plugins inside VS Code)  
2. **Postgres SQL 17.4.1**  
3. **DBeaver 25.0.1** for GUI 
4. **Golang** tested with ```go version go1.24.1 windows/amd64```

## Jerarquia de directorios (falta actualizar)

```
Resy-Backen.
|   .gitignore
|   README.md
|   
+---Datos Extras
|       Correo oficial Salon Verde.txt
|       
+---soa
|   \---services
|       +---ServicioMenu
|       |       build.bat
|       |       
|       \---ServicioUsuario
|           |   build.bat
|           |   go.mod
|           |   go.sum
|           |   
|           +---cmd
|           |       main.go
|           |       
|           \---pkg
|               +---api
|               |   \---handlers
|               |           handlers.go
|               |           
|               +---core
|               |   +---domain
|               |   |       rol.go
|               |   |       usuario.go
|               |   |       
|               |   +---internal
|               |   |       internal.go
|               |   |       
|               |   +---response
|               |   |       user_service_response.go
|               |   |       
|               |   \---usecase
|               |       \---backBD
|               |               Service_Conection_with_BD.go
|               |               
|               \---repository
|                   |   MetodosGenericos.go
|                   |   
|                   +---crypton
|                   |       crypto.go
|                   |       
|                   +---database
|                   |       BDoperators.go
|                   |       database.go
|                   |       
|                   \---interfaces
|                           interface_repository.go
|                           
\---sql
        Creaci¾nDeLaBaseDeDatos.sql
        Insert para verificar usuario y roles.sql
        Notaimportante.txt
```

### Build para testeo local 
- Build: Execute ```build.bat``` within ```soa``` folder using ```cmd.exe``` or ```VS Code terminal```
- Run:  Either ```go run main.go``` or debug within VS Code.

### Developer considerations:
- If additional features were to be added, implement considering:
	```go 
	import(
		"soa/example_dir/example_package"   //example_package would be the new directory within "example_dir" example directory.
	)
	```
	Add as needed inside each ```example_dir``` directory, **DO NOT CHANGE THE TOP-LEVEL DIRECTORY HIERARCHY (soa)**.  
		

### Recommendations
- *Windows*: Make sure the enviroment variable PATH contains the path with the golang binaries, the path is usually ```C:\Go\bin``` 
- *Linux*: T.B.D


### Referencias adicionales
1. https://www.velotio.com/engineering-blog/build-a-containerized-microservice-in-golang
2. https://gokit.io/examples/stringsvc.html
3. https://github.com/athun-me/GO-microservice-clean-architecture/tree/master

