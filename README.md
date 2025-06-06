# Backend - RESI

## Requirements
1. **Visual Studio Code** (you need to install Golang plugins inside VS Code)  
2. **Postgres SQL 17.4.1**  
3. **DBeaver 25.0.1** for GUI 
4. **Golang** tested with ```go version go1.24.1 windows/amd64```

## Directory hierarchy

```lua
mi-proyecto-soa/
├── services/
│   ├── users/
│       ├── cmd/
│       │   └── main.go         # Punto de entrada del servicio de usuarios, contiene su configuracion y como arranca el servicio
│       │         │   ├── internal/           # Todos lo que no se espera que sea importado por otros servicios
│       │   ├── api/              # Handlers HTTP, controladores, mas que todo como se conecta el fronted con el servicio
│       │   │   └── user_handlers.go
│       │   ├── core/             # Lógica de negocio principal (domain, use cases)
│       │   │   ├── domain/
│       │   │   │   └── user.go
│       │   │   └── user_service.go
│       │   └── repository/       # Capa de acceso a datos (interactúa con la DB)
│       │       └── user_repository.go
│       └── pkg/                  # Código reusable y compartido por otros servicios (opcional)
│           └── models/
│               └── user_shared_models.go
│
├── shared/
│   ├── config/                   # Archivos de configuración globales
│   │   └── config.go
│   ├── database/                   # Aqui el Singleton porque tenemos una DB compartida
│   │   └── common_db.go
│   ├── utils/                    # Funciones de utilidad comunes
│   │   └── helpers.go
│   └── proto/                    # Definiciones de gRPC o modelos compartidos (si aplica)
│       └── common.proto
├── go.mod                        # Módulo Go principal
├── go.sum
└── README.md
```

## Building and running the project

### Main project setup
- Build: Execute ```compile.bat``` within ```back``` folder using ```cmd.exe``` or VS Code terminal
	- If additional features were to be added, implement considering:
		```go 
		import(
			"back/src/example_package"   //example_package would be the new directory within "src" dir.
		)
		```
		Add as needed inside ```src``` directory, **DO NOT CHANGE THE DIRECTORY HIERARCHY**.  
		
- Run:  Either ```go run main.go``` or debug within VS Code.

### Recommendations
- *Windows*: Make sure the enviroment variable PATH contains the path with the golang binaries, the path is usually ```C:\Go\bin``` 
- *Linux*: T.B.D
