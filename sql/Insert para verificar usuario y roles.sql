SET search_path TO 'ResyDB', public;

-- Insertar las modalidades de pedido iniciales
INSERT INTO "ModalidadesPedido" ("nombre") values ('Delivery'), ('Recojo en Local');
INSERT INTO "Roles" (nombrerol) values ('Admin'),('Cajero'),('Cliente');
INSERT INTO "MetodosPago" ("nombre") values ('Efectivo'), ('Tarjeta'), ('Yape'), ('Plin');

-- Inserts de prueba para todas las tablas

-- Usuario
INSERT INTO "Usuario" (nombres, apellidos, correo, telefono, direccion, fechanacimiento, contrasenia, rol)
VALUES 
('Juan', 'Pérez', 'juanperez@mail.com', '999888777', 'Av. Siempre Viva 123', '1990-05-10', '123456', 3),
('Ana', 'García', 'anagarcia@mail.com', '988777666', 'Calle Falsa 456', '1985-08-20', 'abcdef', 1);

-- CategoriaPlatos
INSERT INTO "CategoriaPlatos" ("nombre") VALUES ('Entradas'), ('Fondos'), ('Postres');

-- Plato
INSERT INTO "Plato" ("nombrePlato", "categoria", "descripcion", "precio", "imagen")
VALUES 
('Ceviche', 1, 'Ceviche clásico de pescado', 25.00, 'ceviche.jpg'),
('Lomo Saltado', 2, 'Lomo saltado tradicional', 30.00, 'lomo.jpg'),
('Suspiro Limeño', 3, 'Postre tradicional', 12.00, 'suspiro.jpg');

-- Reserva (el id_reserva se genera automáticamente)
INSERT INTO "Reserva" ("id_clienteSolicitante", "fechaHoraReservada", "numPersonas", "estadoReserva", "especificacionesDeLaReserva")
VALUES (1, NOW() + INTERVAL '1 day', 4, 'Pendiente', 'Mesa cerca de la ventana');

-- Pedido (el id_pedido se genera automáticamente)
INSERT INTO "Pedido" ("id_clienteSolicitante", "fecha", "total", "estadopedido", "id_modalidad")
VALUES (1, NOW(), 75.50, 'Registrado', 1);

-- MenuSemanal (el id_menu se genera automáticamente)
INSERT INTO "MenuSemanal" ("fecha_inicio", "fechaFin")
VALUES ('2025-06-24', '2025-06-30');

-- Menudia
INSERT INTO "Menudia" ("id_menu", "dia_semana")
VALUES ('Men1', 'Lunes'), ('Men1', 'Martes');

-- PlatosEnMenudia
INSERT INTO "PlatosEnMenudia" ("id_dia", "id_plato", "cantidadDelPlato")
VALUES (1, 1, 10), (1, 2, 8), (2, 3, 5);

-- Linea_Pedido
INSERT INTO "Linea_Pedido" ("id_pedido", "id_plato", "cantidad", "subtotal")
VALUES ('PED1', 1, 2, 50.00), ('PED1', 3, 1, 12.00);

-- PagoRegistrado
INSERT INTO "PagoRegistrado" ("Nombrepagante", "fecharegistro", "monto", "id_metodo")
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
select * from "Roles";
select * from "Usuario";