-- Corrige el nombre del esquema (elimina el existente si es necesario)
DROP SCHEMA IF EXISTS restaruante CASCADE;
CREATE SCHEMA IF NOT EXISTS restaurante;
SET search_path TO restaurante, public;
CREATE EXTENSION IF NOT EXISTS "pgcrypto";

-- Tabla base Usuario
CREATE TABLE if not exists usuario (
    usuario_id VARCHAR(10) PRIMARY KEY,
    nombres VARCHAR(150),
    apellidos VARCHAR(150),
    tipo_documento VARCHAR(20) CHECK (tipo_documento IN ('DNI', 'CE', 'PASAPORTE')),
    numero_documento VARCHAR(20) UNIQUE,
    telefono VARCHAR(20) NOT NULL,
     email VARCHAR(100) unique,
    fecha_registro DATE NOT NULL,
    activo BOOLEAN DEFAULT TRUE,
    contrasena VARCHAR(255) not null default '*'
);

CREATE table if not exists sede (
	sede_id VARCHAR(36) PRIMARY KEY,
    nombre VARCHAR(100) NOT NULL,
    direccion VARCHAR(200) unique NOT NULL,
    ciudad VARCHAR(50) NOT NULL,
    telefono VARCHAR(20) unique NOT NULL,
    AFORO_TOTAL INT NOT NULL,
    horario_apertura TIME NOT NULL,
    horario_cierre TIME NOT NULL,
    activa BOOLEAN DEFAULT TRUE
);
CREATE TABLE IF NOT EXISTS cliente (
    -- Atributos específicos de cliente
    fecha_nacimiento DATE,
    es_frecuente BOOLEAN DEFAULT FALSE,
    puntos_fidelizacion INT DEFAULT 0,
    restricciones_alimenticias VARCHAR(255),
    -- Clave primaria heredada de usuario
    PRIMARY KEY (usuario_id)
) INHERITS (usuario);

CREATE TABLE IF NOT EXISTS empleado (
    -- Atributos específicos de empleado
    sede_id VARCHAR(36) NOT NULL,
    cargo VARCHAR(50) NOT NULL CHECK (cargo IN ('Mesero', 'Chef', 'Gerente', 'Delivery')),
    fecha_contratacion DATE NOT NULL,
    salario DECIMAL(10, 2) NOT NULL,
    -- Clave foránea a sede
    FOREIGN KEY (sede_id) REFERENCES sede(sede_id),
    -- Clave primaria heredada de usuario
    PRIMARY KEY (usuario_id)
) INHERITS (usuario);

select * from cliente;