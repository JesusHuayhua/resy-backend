@echo off
setlocal enableDelayedExpansion
set SERVICES=ServicioMenu ServicioUsuario
for %%S in (%SERVICES%) do (
  echo [BUILD] %%S
  pushd %%S
    if not exist go.mod (
      echo   go mod init (%%S)
      go mod init %%S
    )
    go mod tidy
    go build -o bin\%%S.exe ./cmd
    if errorlevel 1 (
      echo build fail (%%S)
      exit /b 1
    )
  popd
)
echo [BUILD] all services built successfully.
pause
