# GO Backend Ingesoft 2025-1
GO backend for software engineering class

## Requirements
1. **Visual Studio Code** (you need to install Golang plugins inside VS Code)  
2. **Postgres SQL 17.4.1**  
3. **DBeaver 25.0.1** for GUI 
4. **Golang** tested with ```go version go1.24.1 windows/amd64```

## Directory hierarchy
T.B.D

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

### DB scripts setup
- T.B.D

### Recommendations
- *Windows*: Make sure the enviroment variable PATH contains the path with the golang binaries, the path is usually ```C:\Go\bin``` 
- *Linux*: T.B.D
