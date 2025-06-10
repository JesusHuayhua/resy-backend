# Backend - RESI

## Requerimientos
1. **Visual Studio Code** (you need to install Golang plugins inside VS Code)  
2. **Postgres SQL 17.4.1**  
3. **DBeaver 25.0.1** for GUI 
4. **Golang** tested with ```go version go1.24.1 windows/amd64```

## Jerarquia de directorios

```
soa/
├── services/
│   ├── users/
│       ├── cmd/
│       │   └── main.go         			  
│       │   │   ├── internal/          				 
│       │   ├── api/              			  
│       │   │   └── user_handlers.go
│       │   ├── core/             			 
│       │   │   ├── domain/
│       │   │   │   └── user.go
│       │   │   └── user_service.go
│       │   └── repository/       			  
│       │       └── user_repository.go
│       └── pkg/                 		      
│           └── models/
│               └── user_shared_models.go
│
├── shared/
│   ├── config/                   			   
│   │   └── config.go
│   ├── database/                   	     
│   │   └── common_db.go
│   ├── utils/                    			 
│   │   └── helpers.go
│   └── proto/                    			  
│       └── common.proto
├── go.mod                        			   
├── go.sum
└── README.md
```

### Build para testeo local 
- Build: Execute ```compile.bat``` within ```back``` folder using ```cmd.exe``` or VS Code terminal
	- If additional features were to be added, implement considering:
		```go 
		import(
			"mi-proyecto-soa/dir1/example_package"   //example_package would be the new directory within "dir1" example directory.
		)
		```
		Add as needed inside each ```dir1``` directory, **DO NOT CHANGE THE DIRECTORY HIERARCHY**.  
		
- Run:  Either ```go run main.go``` or debug within VS Code.

### Recommendations
- *Windows*: Make sure the enviroment variable PATH contains the path with the golang binaries, the path is usually ```C:\Go\bin``` 
- *Linux*: T.B.D
