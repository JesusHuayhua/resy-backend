@echo off
if exist go.mod if exist go.sum (
    del go.mod go.sum  
)
echo "[BUILD] Building go project"
go mod init ServicioUsuario
go mod tidy

echo "[BUILD] go project built."
pause