DROP SCHEMA IF EXISTS "ResyDB" CASCADE;
CREATE SCHEMA IF NOT EXISTS "ResyDB";
SET search_path TO "ResyDB", public;
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
  "nombrerol" VARCHAR(10) unique NOT NULL
);

CREATE TABLE "Usuario" (
  "id_usuario" SERIAL PRIMARY KEY,
  "nombres" VARCHAR(50) NOT NULL,
  "apellidos" VARCHAR(50) NOT NULL,
  "correo" VARCHAR(50) unique NOT NULL,  -- Longitud aumentada
  "telefono" varchar(15) unique not null,
  "direccion" text not null,
  "fechanacimiento" DATE,
  "contrasenia" TEXT NOT NULL,  -- Cambiado a TEXT para hashes
  "rol" INT NOT NULL REFERENCES "Roles"("id_rol"),
  "estadoacceso" BOOLEAN default true
);

CREATE TABLE "RecuperacionPassword" (
  correo VARCHAR(50) PRIMARY key references "Usuario"("correo"),
  token VARCHAR(32) NOT NULL,
  expira_en TIMESTAMP NOT NULL
);

CREATE TABLE "Mensaje" (
  "idMensaje" SERIAL PRIMARY KEY,
  "idDestinatario" INT NOT NULL REFERENCES "Usuario"("id_usuario"),
  "fechaHoraMensaje" TIMESTAMP NOT NULL,  -- Corregido a TIMESTAMP
  "contenidoMensaje" VARCHAR(100) NOT NULL
);

CREATE TABLE "Reserva" (
  "id_reserva" VARCHAR(8) PRIMARY KEY,
  "id_clienteSolicitante" INT REFERENCES "Usuario"("id_usuario"),
  "fechaHoraReservada" TIMESTAMP NOT NULL,  -- Corregido a TIMESTAMP
  "numPersonas" INT NOT NULL,
  "estadoReserva" "EstadoReserva" NOT NULL,  -- Referencia directa al tipo
  "especificacionesDeLaReserva" VARCHAR(100)
);

CREATE TABLE "Pedido" (
  "id_pedido" VARCHAR(8) PRIMARY KEY,
  "id_clienteSolicitante" INT REFERENCES "Usuario"("id_usuario"),
  "fecha" TIMESTAMP NOT NULL,  -- Cambiado a TIMESTAMP
  "total" DECIMAL(10,2) NOT NULL,  -- Cambiado a DECIMAL
  "estadopedido" "EstadosPedido" NOT NULL  -- Referencia directa al tipo
);

CREATE TABLE "CategoriaPlatos" (
  "id_categoria" SERIAL PRIMARY KEY,
  "nombre" VARCHAR(20) unique NOT NULL
);

CREATE TABLE "Plato" (
  "id_plato" SERIAL PRIMARY KEY,
  "nombrePlato" VARCHAR(20) NOT NULL,
  "categoria" INT NOT NULL REFERENCES "CategoriaPlatos"("id_categoria"), -- FK corregida
  "descripcion" VARCHAR(200) NOT NULL,
  "precio" DECIMAL(10,2) NOT NULL,  -- Cambiado a DECIMAL
  "imagen" TEXT NOT NULL,
  "estado" BOOLEAN DEFAULT TRUE
);

CREATE TABLE "MenuSemanal" (
  "id_menu" VARCHAR(8) PRIMARY KEY,
  "fechadeinicio" DATE NOT NULL,  -- Cambiado a DATE
  "fechaFin" DATE NOT NULL        -- Cambiado a DATE
);

CREATE TABLE "Menudia" (
  "id_dia" SERIAL PRIMARY KEY,
  "id_menu" VARCHAR(8) REFERENCES "MenuSemanal"("id_menu"),
  "dia_semana" "DiaSemana" NOT NULL
);

CREATE TABLE "PlatosEnMenudia" (
  "id_dia" INT NOT NULL REFERENCES "Menudia"("id_dia"),
  "id_plato" INT NOT NULL REFERENCES "Plato"("id_plato"),
  "cantidadDelPlato" INT not null,
  "disponibleParaVender" BOOLEAN default true,
  PRIMARY KEY ("id_dia", "id_plato")  -- PK compuesta
);

CREATE TABLE "Linea_Pedido" (
  "id_linea" SERIAL PRIMARY KEY,
  "id_pedido" VARCHAR(8) NOT NULL REFERENCES "Pedido"("id_pedido"),
  "id_plato" INT NOT NULL REFERENCES "Plato"("id_plato"),
  "cantidad" INT NOT NULL,
  "subtotal" DECIMAL(10,2) NOT NULL  -- Cambiado a DECIMAL
);

CREATE TABLE "PagoRegistrado" (
  "id_pago" SERIAL PRIMARY KEY,
  "Nombrepagante" VARCHAR(50) NOT NULL,
  "fecharegistro" TIMESTAMP NOT NULL,
  "monto" DECIMAL(10,2) NOT NULL,  -- Cambiado a DECIMAL
  "metodosDePago" "MetodosPago" NOT NULL
);

CREATE TABLE "Reserva_x_pago" (
  "id_pago" INT REFERENCES "PagoRegistrado"("id_pago"),
  "id_reserva" VARCHAR(8) REFERENCES "Reserva"("id_reserva"),
  PRIMARY KEY ("id_pago", "id_reserva")
);

CREATE TABLE "Pedido_x_pago" (
  "id_pago" INT REFERENCES "PagoRegistrado"("id_pago"),
  "id_pedido" VARCHAR(8) REFERENCES "Pedido"("id_pedido"),
  PRIMARY KEY ("id_pago", "id_pedido")
);

CREATE TABLE "PlatosReservados" (
  "id_linea" SERIAL PRIMARY KEY,
  "id_reserva" VARCHAR(8) NOT NULL REFERENCES "Reserva"("id_reserva"),
  "id_plato" INT NOT NULL REFERENCES "Plato"("id_plato"),
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

