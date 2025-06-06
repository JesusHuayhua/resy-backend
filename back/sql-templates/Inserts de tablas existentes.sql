
SET search_path TO restaurante, public;

-- Insert 1: Cliente regular sin restricciones alimenticias
INSERT INTO cliente (id, nombres, apellidos, tipo_documento, numero_documento, telefono, email, fecha_registro, fecha_nacimiento, es_frecuente, puntos_fidelizacion)
VALUES ('CLI001', 'Juan', 'Pérez García', 'DNI', '12345678', '987654321', 'juan.perez@email.com', '2023-01-15', '1990-05-20', FALSE, 0);

-- Insert 2: Cliente frecuente vegetariano
INSERT INTO cliente (id, nombres, apellidos, tipo_documento, numero_documento, telefono, email, fecha_registro, fecha_nacimiento, es_frecuente, puntos_fidelizacion, restricciones_alimenticias)
VALUES ('CLI002', 'María', 'López Fernández', 'DNI', '87654321', '987123456', 'maria.lopez@email.com', '2022-11-10', '1985-08-12', TRUE, 150, 'Vegetariano');

-- Insert 3: Cliente extranjero con pasaporte
INSERT INTO cliente (id, nombres, apellidos, tipo_documento, numero_documento, telefono, email, fecha_registro, fecha_nacimiento, restricciones_alimenticias)
VALUES ('CLI003', 'John', 'Smith', 'PASAPORTE', 'PA1234567', '987456123', 'john.smith@email.com', '2023-03-05', '1982-11-30', 'Vegano, Sin gluten');

-- Insert 4: Cliente con documento de extranjería
INSERT INTO cliente (id, nombres, apellidos, tipo_documento, numero_documento, telefono, email, fecha_registro, fecha_nacimiento, es_frecuente, puntos_fidelizacion)
VALUES ('CLI004', 'Carlos', 'Gómez Ruiz', 'CE', 'X12345678', '987789456', 'carlos.gomez@email.com', '2023-02-20', '1995-04-25', TRUE, 75);

-- Insert 5: Cliente joven recién registrado
INSERT INTO cliente (id, nombres, apellidos, tipo_documento, numero_documento, telefono, email, fecha_registro, fecha_nacimiento)
VALUES ('CLI005', 'Ana', 'Martínez Sánchez', 'DNI', '56781234', '987321654', 'ana.martinez@email.com', '2023-04-01', '2000-07-15');

-- Inserts para la tabla sede
INSERT INTO sede (id, nombre, direccion, ciudad, telefono, AFORO_TOTAL, horario_apertura, horario_cierre)
VALUES ('550e8400-e29b-41d4-a716-446655440000', 'Sede Principal', 'Av. Principal 123', 'Lima', '012345678', 100, '09:00:00', '22:00:00');

INSERT INTO sede (id, nombre, direccion, ciudad, telefono, AFORO_TOTAL, horario_apertura, horario_cierre)
VALUES ('550e8400-e29b-41d4-a716-446655440001', 'Sede Miraflores', 'Av. Larco 456', 'Lima', '012345679', 80, '10:00:00', '23:00:00');

INSERT INTO sede (id, nombre, direccion, ciudad, telefono, AFORO_TOTAL, horario_apertura, horario_cierre)
VALUES ('550e8400-e29b-41d4-a716-446655440002', 'Sede San Isidro', 'Calle Los Pinos 789', 'Lima', '012345680', 60, '08:00:00', '21:00:00');

-- Insert 1: Gerente en la sede principal
INSERT INTO empleado (id, nombres, apellidos, tipo_documento, numero_documento, telefono, email, fecha_registro, sede_id, cargo, fecha_contratacion, salario)
VALUES ('EMP001', 'Roberto', 'Jiménez Vargas', 'DNI', '11222333', '987111222', 'roberto.jimenez@restaurante.com', '2020-05-10', '550e8400-e29b-41d4-a716-446655440000', 'Gerente', '2020-05-10', 5000.00);

-- Insert 2: Chef en la sede principal
INSERT INTO empleado (id, nombres, apellidos, tipo_documento, numero_documento, telefono, email, fecha_registro, sede_id, cargo, fecha_contratacion, salario)
VALUES ('EMP002', 'Lucía', 'Fernández Castro', 'DNI', '44555666', '987333444', 'lucia.fernandez@restaurante.com', '2021-02-15', '550e8400-e29b-41d4-a716-446655440000', 'Chef', '2021-02-15', 3500.00);

-- Insert 3: Mesero en la sede Miraflores
INSERT INTO empleado (id, nombres, apellidos, tipo_documento, numero_documento, telefono, email, fecha_registro, sede_id, cargo, fecha_contratacion, salario)
VALUES ('EMP003', 'Pedro', 'García López', 'DNI', '77888999', '987555666', 'pedro.garcia@restaurante.com', '2022-01-20', '550e8400-e29b-41d4-a716-446655440001', 'Mesero', '2022-01-20', 1800.00);

-- Insert 4: Delivery en la sede San Isidro
INSERT INTO empleado (id, nombres, apellidos, tipo_documento, numero_documento, telefono, email, fecha_registro, sede_id, cargo, fecha_contratacion, salario)
VALUES ('EMP004', 'Sofía', 'Ramírez Díaz', 'DNI', '99000111', '987777888', 'sofia.ramirez@restaurante.com', '2022-06-05', '550e8400-e29b-41d4-a716-446655440002', 'Delivery', '2022-06-05', 1500.00);

-- Insert 5: Chef en la sede Miraflores
INSERT INTO empleado (id, nombres, apellidos, tipo_documento, numero_documento, telefono, email, fecha_registro, sede_id, cargo, fecha_contratacion, salario)
VALUES ('EMP005', 'Jorge', 'Torres Medina', 'DNI', '22333444', '987999000', 'jorge.torres@restaurante.com', '2021-11-15', '550e8400-e29b-41d4-a716-446655440001', 'Chef', '2021-11-15', 3800.00);

