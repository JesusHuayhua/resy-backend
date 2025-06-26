@echo off
setlocal

if "%~1"=="" goto :usage
 
set "CMD=%~1"
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
  call "%~dp0build.bat" build-local
  for %%S in (%SERVICES%) do (
    echo [RUN] %%S ejecutandose
    start "" cmd /k "%BIN_DIR%\%%S.exe"
  )
  exit /b
)

:usage
echo build.bat ^<command^>
echo  build-local   Inicializa y compila.
echo  run-local     Compila + arranca los servicios.

exit /b 1