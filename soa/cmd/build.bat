@echo off
setlocal

if "%~1"=="" goto :usage
 
set "CMD=%~1"
:: Al agregar los servicios 
set "SERVICES=ServicioUsuario"
set "SCRIPT_DIR=%~dp0"
set "SOA_DIR=%SCRIPT_DIR%.."
set "CMD_DIR=%SOA_DIR%\cmd"
set "ENTRYPOINTS_DIR=%CMD_DIR%\entrypoints"
set "BIN_DIR=%CMD_DIR%\bin"

if /i "%CMD%"=="build-local" (
  pushd "%SOA_DIR%"
    if not exist go.mod (
      echo [GO MOD] no encontrado en %SOA_DIR%, inicializando.
      go mod init soa || exit /b 1
      go mod tidy
    ) else (
      echo [GO MOD] ya existe, omitiendo init.
    )
  popd
  if not exist "%BIN_DIR%" mkdir "%BIN_DIR%" 
  for %%S in (%SERVICES%) do (
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

:: buildeamos y ejecutamos.
if /i "%CMD%"=="run-local" (
  rem asegúrate de compilar antes
  call "%~dp0build.bat" build-local

  rem cambiar al directorio de binaria (opcional)
  pushd "%BIN_DIR%"

  rem -- el bloque FOR con paréntesis --
  for %%S in (%SERVICES%) do (
    echo [RUN] %%S en ejecucion...
    start "" cmd /k "%%S.exe"
  )

  popd
  exit /b
)

:usage
echo build.bat ^<command^>
echo  build-local   Inicializa y compila.
echo  run-local     Compila + arranca los servicios.

exit /b 1