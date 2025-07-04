SET search_path TO "ResyDB", public;

-- Modalidades de pedido
INSERT INTO "ModalidadesPedido" ("nombre") VALUES ('Delivery'), ('Recojo en Local');

-- Roles
INSERT INTO "Roles" ("nombrerol") VALUES ('Admin'), ('Cajero'), ('Cliente');

-- Métodos de pago
INSERT INTO "MetodosPago" ("nombre") VALUES ('Efectivo'), ('Tarjeta'), ('Yape'), ('Plin');

-- Usuario (no se debe insertar id_usuario, es SERIAL)
INSERT INTO "Usuario" (nombres, apellidos, correo, telefono, direccion, fechanacimiento, contrasenia, rol)
VALUES 
('Juan', 'Pérez', 'juanperez@mail.com', '999888777', 'Av. Siempre Viva 123', '1990-05-10', 'eyJrbXNDaXBoZXJ0ZXh0IjoiQVFJREFIaW4wYXVRQnR4dXppdldKY1ZHVkRMTThIQllFTTVhbFRhWEV3ZlpqZk1XTFFGMHpDYWlOQlA3S2s4NzhDeW1RN1R4QUFBQWZqQjhCZ2txaGtpRzl3MEJCd2FnYnpCdEFnRUFNR2dHQ1NxR1NJYjNEUUVIQVRBZUJnbGdoa2dCWlFNRUFTNHdFUVFNNitCNGMxVysyS3lQZWVMaEFnRVFnRHVwbTZhYXYvd3p1azF6NnZMTm5xQnBJc0NmN0JIVkpIMlZnMUxNZWM2Lzc5cmlhdmNXM2ZzeDFrQ1pPcnhTN2RITkJHR1AyVmNLNkhBZ1lnPT0iLCJ3cmFwTm9uY2UiOiIxd3cwSElrSllIT0dSallsIiwid3JhcHBlZEtleSI6IkVCV1RUL1ppT1Z0OGsyUGM0WElTWDBFODRMNER5SzU5VzV6K3FTSFFWeGRoMExOaUhkYWJZNnRiMVBmNW1OREciLCJwYXlsb2FkTm9uY2UiOiJkWU5oVUtoWUJCSWc3L2JDIiwiY2lwaGVydGV4dCI6Ii9tU0Vhb2o0cDJEeDJjbE40L2NnQklqc1lmcXB4VzFPcDFseSIsInNhbHQiOiJsTFlPSHhaazFrMFZ0bVBjQU5PQnd3PT0iLCJpdGVyIjoxNTAwMDB9', 3),
('Juan', 'Kambria', 'lolapalusa@mail.com', '993388777', 'Av. Siempre Viva 323', '1990-05-10', 'eyJrbXNDaXBoZXJ0ZXh0IjoiQVFJREFIaW4wYXVRQnR4dXppdldKY1ZHVkRMTThIQllFTTVhbFRhWEV3ZlpqZk1XTFFGMHpDYWlOQlA3S2s4NzhDeW1RN1R4QUFBQWZqQjhCZ2txaGtpRzl3MEJCd2FnYnpCdEFnRUFNR2dHQ1NxR1NJYjNEUUVIQVRBZUJnbGdoa2dCWlFNRUFTNHdFUVFNNitCNGMxVysyS3lQZWVMaEFnRVFnRHVwbTZhYXYvd3p1azF6NnZMTm5xQnBJc0NmN0JIVkpIMlZnMUxNZWM2Lzc5cmlhdmNXM2ZzeDFrQ1pPcnhTN2RITkJHR1AyVmNLNkhBZ1lnPT0iLCJ3cmFwTm9uY2UiOiIxd3cwSElrSllIT0dSallsIiwid3JhcHBlZEtleSI6IkVCV1RUL1ppT1Z0OGsyUGM0WElTWDBFODRMNER5SzU5VzV6K3FTSFFWeGRoMExOaUhkYWJZNnRiMVBmNW1OREciLCJwYXlsb2FkTm9uY2UiOiJkWU5oVUtoWUJCSWc3L2JDIiwiY2lwaGVydGV4dCI6Ii9tU0Vhb2o0cDJEeDJjbE40L2NnQklqc1lmcXB4VzFPcDFseSIsInNhbHQiOiJsTFlPSHhaazFrMFZ0bVBjQU5PQnd3PT0iLCJpdGVyIjoxNTAwMDB9', 2),
('Ana', 'García', 'anagarcia@mail.com', '988777666', 'Calle Falsa 456', '1985-08-20', 'eyJrbXNDaXBoZXJ0ZXh0IjoiQVFJREFIaW4wYXVRQnR4dXppdldKY1ZHVkRMTThIQllFTTVhbFRhWEV3ZlpqZk1XTFFGMHpDYWlOQlA3S2s4NzhDeW1RN1R4QUFBQWZqQjhCZ2txaGtpRzl3MEJCd2FnYnpCdEFnRUFNR2dHQ1NxR1NJYjNEUUVIQVRBZUJnbGdoa2dCWlFNRUFTNHdFUVFNNitCNGMxVysyS3lQZWVMaEFnRVFnRHVwbTZhYXYvd3p1azF6NnZMTm5xQnBJc0NmN0JIVkpIMlZnMUxNZWM2Lzc5cmlhdmNXM2ZzeDFrQ1pPcnhTN2RITkJHR1AyVmNLNkhBZ1lnPT0iLCJ3cmFwTm9uY2UiOiIxd3cwSElrSllIT0dSallsIiwid3JhcHBlZEtleSI6IkVCV1RUL1ppT1Z0OGsyUGM0WElTWDBFODRMNER5SzU5VzV6K3FTSFFWeGRoMExOaUhkYWJZNnRiMVBmNW1OREciLCJwYXlsb2FkTm9uY2UiOiJkWU5oVUtoWUJCSWc3L2JDIiwiY2lwaGVydGV4dCI6Ii9tU0Vhb2o0cDJEeDJjbE40L2NnQklqc1lmcXB4VzFPcDFseSIsInNhbHQiOiJsTFlPSHhaazFrMFZ0bVBjQU5PQnd3PT0iLCJpdGVyIjoxNTAwMDB9', 1);

-- Categorías de platos
INSERT INTO "CategoriaPlatos" ("nombre") VALUES ('Entradas'), ('Fondos'), ('Postres');

-- Plato
INSERT INTO "Plato" ("nombre_plato", "categoria", "descripcion", "precio", "imagen")
VALUES 
('Ceviche', 1, 'Ceviche clásico de pescado', 25.00, 'https://resy-ingesoft.s3.us-east-1.amazonaws.com/imagen_2025-06-21_213926386.png'),
('Lomo Saltado', 2, 'Lomo saltado tradicional', 30.00, 'https://resy-ingesoft.s3.us-east-1.amazonaws.com/imagen_2025-06-21_213926386.png'),
('Suspiro Limeño', 3, 'Postre tradicional', 12.00, 'https://resy-ingesoft.s3.us-east-1.amazonaws.com/imagen_2025-06-21_213926386.png');

-- MenuSemanal (id_menu se autogenera por trigger)
INSERT INTO "MenuSemanal" ("fecha_inicio", "fecha_fin")
VALUES ('2025-06-24', '2025-06-30');

-- Menudia (id_dia es SERIAL, id_menu debe existir)
INSERT INTO "Menudia" ("id_menu", "dia_semana")
VALUES ('Men1', 'Lunes'), ('Men1', 'Martes');

-- PlatosEnMenudia (id_dia e id_plato deben existir)
INSERT INTO "PlatosEnMenudia" ("id_dia", "id_plato", "cantidad_plato", "disponible_venta")
VALUES (1, 1, 10, true), (1, 2, 8, true), (2, 3, 5, true);

-- Reserva (el id_reserva se genera automáticamente)
INSERT INTO "Reserva" ("id_cliente", "fecha_reservada", "num_personas", "estado_reserva", "especificaciones")
VALUES (1, NOW() + INTERVAL '1 day', 4, 'Pendiente', 'Mesa cerca de la ventana');

-- Reserva para usuario registrado
INSERT INTO "Reserva" ("id_cliente", "fecha_reservada", "num_personas", "estado_reserva", "especificaciones")
VALUES (1, NOW() + INTERVAL '1 day', 4, 'Pendiente', 'Mesa cerca de la ventana');

-- Reserva para NO registrado
INSERT INTO "Reserva" ("nombre_cliente", "telefono_cliente", "correo_cliente", "fecha_reservada", "num_personas", "estado_reserva", "especificaciones")
VALUES ('Carlos López', '912345678', 'carlos@mail.com', NOW() + INTERVAL '2 day', 2, 'Pendiente', 'Sin registro');

-- Pedido (el id_pedido se genera automáticamente)
INSERT INTO "Pedido" ("id_cliente", "fecha", "total", "estado_pedido", "id_modalidad")
VALUES (1, NOW(), 75.50, 'Registrado', 1);

-- Pedido para usuario registrado
INSERT INTO "Pedido" ("id_cliente", "fecha", "total", "estado_pedido", "id_modalidad")
VALUES (1, NOW(), 75.50, 'Registrado', 1);

-- Pedido para NO registrado
INSERT INTO "Pedido" ("nombre_cliente", "telefono_cliente", "correo_cliente", "fecha", "total", "estado_pedido", "id_modalidad")
VALUES ('Maria Torres', '987654321', 'maria@mail.com', NOW(), 50.00, 'Registrado', 2);

-- Linea_Pedido
INSERT INTO "Linea_Pedido" ("id_pedido", "id_plato", "cantidad_plato", "subtotal")
VALUES ('PED1', 1, 2, 50.00), ('PED1', 3, 1, 12.00);

-- PagoRegistrado
INSERT INTO "PagoRegistrado" ("nombre_pagante", "fecha_registro", "monto", "id_metodo")
VALUES ('Juan Pérez', NOW(), 62.00, 1);

-- Reserva_x_pago
INSERT INTO "Reserva_x_pago" ("id_pago", "id_reserva")
VALUES (1, 'RES1');

-- Pedido_x_pago
INSERT INTO "Pedido_x_pago" ("id_pago", "id_pedido")
VALUES (1, 'PED1');

-- PlatosReservados
INSERT INTO "PlatosReservados" ("id_reserva", "id_plato", "cantidad", "subtotal")
VALUES ('RES1', 1, 2, 50.00);

-- InformacionLocal
INSERT INTO "InformacionLocal" ("horarios", "direccion", "telefono", "correo", "facebook")
VALUES ('Lun-Dom 8am-10pm', 'Av. Principal 123', '987654321', 'info@resy.com', 'facebook.com/resy');

-- Más Platos
INSERT INTO "Plato" ("nombre_plato", "categoria", "descripcion", "precio", "imagen")
VALUES
('Aji de Gallina', 2, 'Clásico ají de gallina cremoso', 28.00, 'https://resy-ingesoft.s3.us-east-1.amazonaws.com/imagen_2025-06-21_213926386.png'),
('Causa Rellena', 1, 'Causa de papa con relleno de atún', 20.00, 'https://resy-ingesoft.s3.us-east-1.amazonaws.com/imagen_2025-06-21_213926386.png'),
('Arroz con Leche', 3, 'Postre tradicional de arroz con leche', 10.00, 'https://resy-ingesoft.s3.us-east-1.amazonaws.com/imagen_2025-06-21_213926386.png'),
('Sopa Criolla', 1, 'Sopa sustanciosa con fideos y huevo', 15.00, 'https://resy-ingesoft.s3.us-east-1.amazonaws.com/imagen_2025-06-21_213926386.png'),
('Arroz con Pato', 2, 'Arroz verde con pato tierno', 35.00, 'https://resy-ingesoft.s3.us-east-1.amazonaws.com/imagen_2025-06-21_213926386.png'),
('Mazamorra Morada', 3, 'Dulce tradicional de maíz morado', 9.00, 'https://resy-ingesoft.s3.us-east-1.amazonaws.com/imagen_2025-06-21_213926386.png');

-- Nuevo MenuSemanal (id_menu se autogenera por trigger)
INSERT INTO "MenuSemanal" ("fecha_inicio", "fecha_fin")
VALUES ('2025-07-01', '2025-07-07');

-- Más Menudia para 'Men1' (Semana 1: 2025-06-24 a 2025-06-30)
INSERT INTO "Menudia" ("id_menu", "dia_semana")
VALUES
('Men1', 'Miercoles'),
('Men1', 'Jueves'),
('Men1', 'Viernes'),
('Men1', 'Sabado'),
('Men1', 'Domingo');

-- Menudia para 'Men2' (Semana 2: 2025-07-01 a 2025-07-07)
INSERT INTO "Menudia" ("id_menu", "dia_semana")
VALUES
('Men2', 'Lunes'),
('Men2', 'Martes'),
('Men2', 'Miercoles'),
('Men2', 'Jueves'),
('Men2', 'Viernes'),
('Men2', 'Sabado'),
('Men2', 'Domingo');

-- PlatosEnMenudia para los nuevos días de 'Men1'
-- Asumiendo que el id_dia para Miercoles es 3, Jueves 4, etc. (ajusta según tus IDs reales si usas una secuencia que ya ha corrido)
-- Para este ejemplo, estoy asumiendo que los IDs de Menudia se incrementan secuencialmente.
-- Puedes verificar los IDs con SELECT * FROM "Menudia"; si no estás seguro.
INSERT INTO "PlatosEnMenudia" ("id_dia", "id_plato", "cantidad_plato", "disponible_venta")
VALUES
-- Miercoles (id_dia 3, asumiendo)
(3, 1, 12, true), -- Ceviche
(3, 4, 7, true),  -- Aji de Gallina
(3, 6, 6, true),  -- Arroz con Leche

-- Jueves (id_dia 4, asumiendo)
(4, 5, 9, true),  -- Causa Rellena
(4, 2, 10, true), -- Lomo Saltado
(4, 7, 5, true),  -- Sopa Criolla

-- Viernes (id_dia 5, asumiendo)
(5, 1, 15, true), -- Ceviche
(5, 8, 8, true),  -- Arroz con Pato
(5, 3, 10, true), -- Suspiro Limeño

-- Sabado (id_dia 6, asumiendo)
(6, 4, 11, true), -- Aji de Gallina
(6, 2, 9, true),  -- Lomo Saltado
(6, 9, 7, true),  -- Mazamorra Morada

-- Domingo (id_dia 7, asumiendo)
(7, 7, 10, true), -- Sopa Criolla
(7, 8, 10, true), -- Arroz con Pato
(7, 3, 8, true);  -- Suspiro Limeño

-- PlatosEnMenudia para los días de 'Men2'
-- Asumiendo que los IDs de Menudia para 'Men2' comienzan desde 8.
INSERT INTO "PlatosEnMenudia" ("id_dia", "id_plato", "cantidad_plato", "disponible_venta")
VALUES
-- Lunes (id_dia 8, asumiendo)
(8, 2, 10, true), -- Lomo Saltado
(8, 4, 8, true),  -- Aji de Gallina
(8, 6, 6, true),  -- Arroz con Leche

-- Martes (id_dia 9, asumiendo)
(9, 1, 12, true), -- Ceviche
(9, 5, 7, true),  -- Causa Rellena
(9, 9, 5, true),  -- Mazamorra Morada

-- Miercoles (id_dia 10, asumiendo)
(10, 8, 9, true), -- Arroz con Pato
(10, 7, 10, true), -- Sopa Criolla
(10, 3, 12, true); -- Suspiro Limeño

select * from "MenuSemanal";
select * from "Menudia";
select * from "Roles";
select * from "ResyDB"."Plato";