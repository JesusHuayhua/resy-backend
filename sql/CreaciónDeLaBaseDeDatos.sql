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

CREATE TYPE "DiaSemana" AS ENUM (
  'Lunes',
  'Martes',
  'Miercoles',
  'Jueves',
  'Viernes',
  'Sabado',
  'Domingo'
);

-- Tablas
CREATE TABLE "Roles" (
  "id_rol" SERIAL PRIMARY KEY,
  "nombrerol" VARCHAR(10) unique NOT NULL
);

-- Tabla para métodos de pago
CREATE TABLE "MetodosPago" (
  "id_metodo" SERIAL PRIMARY KEY,
  "nombre" VARCHAR(10) UNIQUE NOT NULL
);

-- Nueva tabla para modalidades de pedido
CREATE TABLE "ModalidadesPedido" (
  "id_modalidad" SERIAL PRIMARY KEY,
  "nombre" VARCHAR(20) UNIQUE NOT NULL
);

CREATE TABLE "Usuario" (
  "id_usuario" SERIAL PRIMARY KEY,
  "nombres" VARCHAR(50) NOT NULL,
  "apellidos" VARCHAR(50) NOT NULL,
  "correo" VARCHAR(50) unique NOT NULL,
  "telefono" varchar(15) unique not null,
  "direccion" text not null,
  "fechanacimiento" DATE,
  "contrasenia" TEXT NOT NULL,
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
  "fechaHoraMensaje" TIMESTAMP NOT NULL,
  "contenidoMensaje" VARCHAR(100) NOT NULL
);

CREATE TABLE "Reserva" (
  "id_reserva" VARCHAR(8) PRIMARY KEY,
  "id_clienteSolicitante" INT REFERENCES "Usuario"("id_usuario"),
  "fechaHoraReservada" TIMESTAMP NOT NULL,
  "numPersonas" INT NOT NULL,
  "estadoReserva" "EstadoReserva" NOT NULL,
  "especificacionesDeLaReserva" VARCHAR(100)
);

-- Tabla Pedido modificada con relación a ModalidadesPedido
CREATE TABLE "Pedido" (
  "id_pedido" VARCHAR(8) PRIMARY KEY,
  "id_clienteSolicitante" INT REFERENCES "Usuario"("id_usuario"),
  "fecha" TIMESTAMP NOT NULL,
  "total" DECIMAL(10,2) NOT NULL,
  "estadopedido" "EstadosPedido" NOT NULL,
  "id_modalidad" INT NOT NULL REFERENCES "ModalidadesPedido"("id_modalidad")  -- Nueva relación
);

CREATE TABLE "CategoriaPlatos" (
  "id_categoria" SERIAL PRIMARY KEY,
  "nombre" VARCHAR(20) unique NOT NULL
);

CREATE TABLE "Plato" (
  "id_plato" SERIAL PRIMARY KEY,
  "nombrePlato" VARCHAR(20) NOT NULL,
  "categoria" INT NOT NULL REFERENCES "CategoriaPlatos"("id_categoria"),
  "descripcion" VARCHAR(200) NOT NULL,
  "precio" DECIMAL(10,2) NOT NULL,
  "imagen" TEXT NOT NULL,
  "estado" BOOLEAN DEFAULT TRUE
);

CREATE TABLE "MenuSemanal" (
  "id_menu" VARCHAR(8) PRIMARY KEY,
  "fechadeinicio" DATE NOT NULL,
  "fechaFin" DATE NOT NULL
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
  PRIMARY KEY ("id_dia", "id_plato")
);

CREATE TABLE "Linea_Pedido" (
  "id_linea" SERIAL PRIMARY KEY,
  "id_pedido" VARCHAR(8) NOT NULL REFERENCES "Pedido"("id_pedido"),
  "id_plato" INT NOT NULL REFERENCES "Plato"("id_plato"),
  "cantidad" INT NOT NULL,
  "subtotal" DECIMAL(10,2) NOT NULL
);

-- Tabla PagoRegistrado con relación a MetodosPago
CREATE TABLE "PagoRegistrado" (
  "id_pago" SERIAL PRIMARY KEY,
  "Nombrepagante" VARCHAR(50) NOT NULL,
  "fecharegistro" TIMESTAMP NOT NULL,
  "monto" DECIMAL(10,2) NOT NULL,
  "id_metodo" INT NOT NULL REFERENCES "MetodosPago"("id_metodo")
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
  "subtotal" DECIMAL(10,2) NOT NULL
);

CREATE TABLE "InformacionLocal" (
  "id_info" SERIAL PRIMARY KEY,
  "horarios" VARCHAR(100),
  "direccion" VARCHAR(100),
  "telefono" VARCHAR(20),
  "correo" VARCHAR(50),
  "facebook" VARCHAR(100)
);

-- Insertar las modalidades de pedido iniciales
INSERT INTO "ModalidadesPedido" ("nombre") VALUES
  ('Delivery'),
  ('Recojo en Local');

INSERT INTO "Roles" (nombrerol) values ('Admin'),('Cajero'),('Cliente');

INSERT INTO "MetodosPago" ("nombre") values ('Efectivo'), ('Tarjeta'), ('Yape'), ('Plin');

