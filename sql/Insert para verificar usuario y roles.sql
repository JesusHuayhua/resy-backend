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
('Juan', 'Pérez', 'juanperez@mail.com', '999888777', 'Av. Siempre Viva 123', '1990-05-10', 'Vc/yfJJ/JzCMwtrpMJ2uHNF0hYEn1vDOaz8Px+pkKp6WTTXRGiX6ALx/', 3),
('Ana', 'García', 'anagarcia@mail.com', '988777666', 'Calle Falsa 456', '1985-08-20', 'Vc/yfJJ/JzCMwtrpMJ2uHNF0hYEn1vDOaz8Px+pkKp6WTTXRGiX6ALx/', 1);

-- Categorías de platos
INSERT INTO "CategoriaPlatos" ("nombre") VALUES ('Entradas'), ('Fondos'), ('Postres');

-- Plato
INSERT INTO "Plato" ("nombre_plato", "categoria", "descripcion", "precio", "imagen")
VALUES 
('Ceviche', 1, 'Ceviche clásico de pescado', 25.00, 'ceviche.jpg'),
('Lomo Saltado', 2, 'Lomo saltado tradicional', 30.00, 'lomo.jpg'),
('Suspiro Limeño', 3, 'Postre tradicional', 12.00, 'suspiro.jpg');

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
INSERT INTO "Reserva" ("id_cliente", "fecha_reservada", "numPersonas", "estado_reserva", "especificaciones")
VALUES (1, NOW() + INTERVAL '1 day', 4, 'Pendiente', 'Mesa cerca de la ventana');

-- Reserva para usuario registrado
INSERT INTO "Reserva" ("id_cliente", "fecha_reservada", "numPersonas", "estado_reserva", "especificaciones")
VALUES (1, NOW() + INTERVAL '1 day', 4, 'Pendiente', 'Mesa cerca de la ventana');

-- Reserva para NO registrado
INSERT INTO "Reserva" ("nombre_cliente", "telefono_cliente", "correo_cliente", "fecha_reservada", "numPersonas", "estado_reserva", "especificaciones")
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

select * from "MenuSemanal";
select * from "Menudia";
select * from "Roles";
select * from "ResyDB"."Usuario";