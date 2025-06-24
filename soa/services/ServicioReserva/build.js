// build.js
const { execSync } = require('child_process');
const fs = require('fs');
const readline = require('readline');

try {
  // Eliminar archivos
  ['go.mod', 'go.sum'].forEach(file => {
    if (fs.existsSync(file)) fs.unlinkSync(file);
  });

  console.log('[BUILD] Building go project');
  execSync('go mod init ServicioReserva', { stdio: 'inherit' });
  execSync('go mod tidy', { stdio: 'inherit' });
  // Compilar el servicio
  execSync('go build -o servicio_reserva.exe ./cmd/main.go', { stdio: 'inherit' });
  console.log('[BUILD] go project built.');
} catch (error) {
  console.error('Error:', error.message);
  process.exit(1);
}

// Pausa
const rl = readline.createInterface({
  input: process.stdin,
  output: process.stdout
});
rl.question('Press Enter to continue...', () => {
  rl.close();
});