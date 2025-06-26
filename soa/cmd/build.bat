@echo off
setlocal enabledelayedexpansion

if "%~1"=="" goto :usage
set "CMD=%~1"
set "SCRIPT_DIR=%~dp0"
set "SOA_DIR=%SCRIPT_DIR%.."
set "ENTRYPOINTS_DIR=%SOA_DIR%\cmd\entrypoints"
set "BIN_DIR=%SOA_DIR%\cmd\bin"

if /i "%CMD%"=="build-local" (
  pushd "%SOA_DIR%"
    if not exist go.mod (
      echo [GO MOD] no encontrado, inicializando…
      go mod init soa || exit /b 1
      go mod tidy
    ) else (
      echo [GO MOD] ya existe, omitiendo.
    )
  popd
  if not exist "%BIN_DIR%" mkdir "%BIN_DIR%"
  for /D %%D in ("%ENTRYPOINTS_DIR%\*") do (
    set "SVC=%%~nxD"
    echo [BUILD] !SVC!
    pushd "%ENTRYPOINTS_DIR%\!SVC!"
      go build -ldflags "-s -w" -o "%BIN_DIR%\!SVC!.exe" . || (
        echo ERROR building !SVC!
        exit /b 1
      )
    popd
  )
  echo Compilación completada.
  exit /b
)

if /i "%CMD%"=="run-local" (
  call "%~dp0build.bat" build-local
  pushd "%BIN_DIR%"
  for /D %%D in ("%ENTRYPOINTS_DIR%\*") do (
    set "SVC=%%~nxD"
    echo [RUN] !SVC!...
    set "EXE=%BIN_DIR%\!SVC!.exe"
    if exist "!EXE!" (
      start "" cmd /k "!EXE!"
    ) else (
      echo "!EXE!" not found, skip.
    )
  )
  popd
  exit /b
)

:usage
echo.
echo Uso: build.bat ^<comando^>
echo.
echo   build-local   Inicializa y compila todos los entrypoints.
echo   run-local     Compila y arranca cada ejecutable en su propia ventana.
echo.
exit /b 1