@echo off
if exist go.mod if exist go.sum (
    del go.mod go.sum  
)
echo "[BUILD] Building go project"
go mod init ServicioUsuario
:: Instalar paquetes requeridos
go get github.com/Shauanth/Singleton_Encription_ServiceGolang
go get github.com/lib/pq
go get gopkg.in/alexcesaro/quotedprintable.v3
go get gopkg.in/gomail.v2
go mod tidy
echo "[BUILD] go project built."
pause