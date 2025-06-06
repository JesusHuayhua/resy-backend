DROP SCHEMA IF EXISTS "restaurante" CASCADE;
CREATE SCHEMA IF NOT EXISTS "restaurante";
SET search_path TO "restaurante", public;
CREATE EXTENSION IF NOT EXISTS "pgcrypto";

-- Tipos ENUM corregidos
CREATE TYPE "EstadoReserva" AS ENUM (
  'Pendiente',
  'Confirmada',
  'Cancelada'
);

CREATE TYPE "EstadosPedido" AS ENUM (
  'Registrado',
  'Pendiente',
  'Entregado',
  'Cancelado',
  'Rechazado'
);

CREATE TYPE "ModalidadPedido" AS ENUM (
  'Delivery',
  'Recojo en Local'
);

CREATE TYPE "DiaSemana" AS ENUM (
  'Lunes',
  'Martes',
  'Miercoles',
  'Jueves',
  'Viernes',
  'Sabado',
  'Domingo'
);

CREATE TYPE "MetodosPago" AS ENUM (
  'Efectivo',
  'Tarjeta',
  'Yape',
  'Plin'
);

-- Tablas
CREATE TABLE "Roles" (
  "id_rol" SERIAL PRIMARY KEY,
  "NombreRol" VARCHAR(10) NOT NULL
);

CREATE TABLE "Usuario" (
  "Id_usuario" SERIAL PRIMARY KEY,
  "Nombres" VARCHAR(50) NOT NULL,
  "Apellidos" VARCHAR(50) NOT NULL,
  "correo" VARCHAR(50) NOT NULL,  -- Longitud aumentada
  "fechaNacimiento" DATE,
  "contrasenia" TEXT NOT NULL,  -- Cambiado a TEXT para hashes
  "rol" INT NOT NULL REFERENCES "Roles"("id_rol"),
  "EstadoAcceso" BOOLEAN NOT NULL
);

CREATE TABLE "Mensaje" (
  "idMensaje" SERIAL PRIMARY KEY,
  "idDestinatario" INT NOT NULL REFERENCES "Usuario"("Id_usuario"),
  "fechaHoraMensaje" TIMESTAMP NOT NULL,  -- Corregido a TIMESTAMP
  "ContenidoMensaje" VARCHAR(100) NOT NULL
);

CREATE TABLE "Reserva" (
  "Id_reserva" VARCHAR(8) PRIMARY KEY,
  "id_clienteSolicitante" INT REFERENCES "Usuario"("Id_usuario"),
  "fechaHoraReservada" TIMESTAMP NOT NULL,  -- Corregido a TIMESTAMP
  "numPersonas" INT NOT NULL,
  "estadoReserva" "EstadoReserva" NOT NULL,  -- Referencia directa al tipo
  "especificacionesDeLaReserva" VARCHAR(100)
);

CREATE TABLE "Pedido" (
  "Id_Pedido" VARCHAR(8) PRIMARY KEY,
  "id_clienteSolicitante" INT REFERENCES "Usuario"("Id_usuario"),
  "fecha" TIMESTAMP NOT NULL,  -- Cambiado a TIMESTAMP
  "total" DECIMAL(10,2) NOT NULL,  -- Cambiado a DECIMAL
  "EstadoPedido" "EstadosPedido" NOT NULL  -- Referencia directa al tipo
);

CREATE TABLE "CategoriaPlatos" (
  "idCategoria" SERIAL PRIMARY KEY,
  "nombre" VARCHAR(20) NOT NULL
);

CREATE TABLE "Plato" (
  "Id_Plato" SERIAL PRIMARY KEY,
  "NombrePlato" VARCHAR(20) NOT NULL,
  "Categoria" INT NOT NULL REFERENCES "CategoriaPlatos"("idCategoria"), -- FK corregida
  "descripcion" VARCHAR(200) NOT NULL,
  "precio" DECIMAL(10,2) NOT NULL,  -- Cambiado a DECIMAL
  "imagen" TEXT NOT NULL,
  "estado" BOOLEAN DEFAULT true
);

CREATE TABLE "MenuSemanal" (
  "idMenu" VARCHAR(8) PRIMARY KEY,
  "fechaDeInicio" DATE NOT NULL,  -- Cambiado a DATE
  "fechaFin" DATE NOT NULL        -- Cambiado a DATE
);

CREATE TABLE "Menudia" (
  "idDia" SERIAL PRIMARY KEY,
  "idMenu" VARCHAR(8) REFERENCES "MenuSemanal"("idMenu"),
  "Dia_semana" "DiaSemana" NOT NULL
);

CREATE TABLE "PlatosEnMenudia" (
  "id_dia" INT NOT NULL REFERENCES "Menudia"("idDia"),
  "id_plato" INT NOT NULL REFERENCES "Plato"("Id_Plato"),
  "cantidadDelPlato" INT,
  "DisponibleParaVender" BOOLEAN,
  PRIMARY KEY ("id_dia", "id_plato")  -- PK compuesta
);

CREATE TABLE "Linea_Pedido" (
  "id_linea" SERIAL PRIMARY KEY,
  "id_pedido" VARCHAR(8) NOT NULL REFERENCES "Pedido"("Id_Pedido"),
  "id_plato" INT NOT NULL REFERENCES "Plato"("Id_Plato"),
  "cantidad" INT NOT NULL,
  "subtotal" DECIMAL(10,2) NOT NULL  -- Cambiado a DECIMAL
);

CREATE TABLE "PagoRegistrado" (
  "id_pago" SERIAL PRIMARY KEY,
  "NombrePagante" VARCHAR(50) NOT NULL,
  "fechaRegistro" TIMESTAMP NOT NULL,
  "monto" DECIMAL(10,2) NOT NULL,  -- Cambiado a DECIMAL
  "metodosDePago" "MetodosPago" NOT NULL
);

CREATE TABLE "Reserva_x_pago" (
  "id_pago" INT REFERENCES "PagoRegistrado"("id_pago"),
  "id_reserva" VARCHAR(8) REFERENCES "Reserva"("Id_reserva"),
  PRIMARY KEY ("id_pago", "id_reserva")
);

CREATE TABLE "Pedido_x_pago" (
  "id_pago" INT REFERENCES "PagoRegistrado"("id_pago"),
  "id_Pedido" VARCHAR(8) REFERENCES "Pedido"("Id_Pedido"),
  PRIMARY KEY ("id_pago", "id_Pedido")
);

CREATE TABLE "PlatosReservados" (
  "id_linea" SERIAL PRIMARY KEY,
  "id_reserva" VARCHAR(8) NOT NULL REFERENCES "Reserva"("Id_reserva"),
  "id_plato" INT NOT NULL REFERENCES "Plato"("Id_Plato"),
  "cantidad" INT NOT NULL,
  "subtotal" DECIMAL(10,2) NOT NULL  -- Cambiado a DECIMAL
);

CREATE TABLE "InformacionLocal" (
  "id_info" SERIAL PRIMARY KEY,  -- PK añadida
  "horarios" VARCHAR(100),
  "direccion" VARCHAR(100),
  "telefono" VARCHAR(20),
  "correo" VARCHAR(50),
  "facebook" VARCHAR(100)  -- Nombre corregido
);

