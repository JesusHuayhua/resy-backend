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
