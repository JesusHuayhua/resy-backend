@echo off
setlocal

rem build.bat <command>
rem   build-local   Inicializa go.mod en soa/ (si hace falta) y compila ambos servicios en soa\cmd\bin
rem   run-local     Compila + arranca los servicios (puertos 8081/8082)

if "%~1"=="" (
  set "CMD=build-local"
) else (
  set "CMD=%~1"
)

set "SCRIPT_DIR=%~dp0"
set "SOA_DIR=%SCRIPT_DIR%.."
set "CMD_DIR=%SOA_DIR%\cmd"
set "ENTRYPOINTS_DIR=%CMD_DIR%\entrypoints"
set "BIN_DIR=%CMD_DIR%\bin"

if /i "%CMD%"=="build-local" (
  pushd "%SOA_DIR%"
    if not exist go.mod (
      echo [GO MOD] no encontrado en %SOA_DIR%, inicializando módulo soa...
      go mod init soa || exit /b 1
      go mod tidy
    ) else (
      echo [GO MOD] ya existe, omitiendo init.
    )
  popd
  if not exist "%BIN_DIR%" mkdir "%BIN_DIR%" 
  for %%S in (ServicioUsuario) do (
    echo [BUILD] %%S
    pushd "%ENTRYPOINTS_DIR%\%%S"
      go build -ldflags "-s -w" -o "%BIN_DIR%\%%S.exe" . || (
        echo Build failed for %%S
        exit /b 1
      )
    popd
  )

  echo Compilacion completada.
  exit /b
)

if /i "%CMD%"=="run-local" (
  call "%~dp0build.bat" build-local
  echo [RUN] ServicioMenu en 8081
  start "" cmd /k "%BIN_DIR%\ServicioMenu.exe"
  echo [RUN] ServicioUsuario en 8082
  start "" cmd /k "%BIN_DIR%\ServicioUsuario.exe"
  exit /b
)

echo Comando desconocido: %CMD%
echo Uso: %~nx0 [build-local^|run-local]
exit /b 1