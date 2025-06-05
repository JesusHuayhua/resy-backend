# Backend - RESI

Estructura general del proyecto:

```lua
mi-proyecto-soa/
├── services/
│   ├── users/
│       ├── cmd/
│       │   └── main.go         # Punto de entrada del servicio de usuarios, contiene su configuracion y como arranca el servicio
│       │         │   ├── internal/           # Todos lo que no se espera que sea importado por otros servicios
│       │   ├── api/              # Handlers HTTP, controladores, mas que todo como se conecta el fronted con el servicio
│       │   │   └── user_handlers.go
│       │   ├── core/             # Lógica de negocio principal (domain, use cases)
│       │   │   ├── domain/
│       │   │   │   └── user.go
│       │   │   └── user_service.go
│       │   └── repository/       # Capa de acceso a datos (interactúa con la DB)
│       │       └── user_repository.go
│       └── pkg/                  # Código reusable y compartido por otros servicios (opcional)
│           └── models/
│               └── user_shared_models.go
│
├── shared/
│   ├── config/                   # Archivos de configuración globales
│   │   └── config.go
│   ├── database/                   # Aqui el Singleton porque tenemos una DB compartida
│   │   └── common_db.go
│   ├── utils/                    # Funciones de utilidad comunes
│   │   └── helpers.go
│   └── proto/                    # Definiciones de gRPC o modelos compartidos (si aplica)
│       └── common.proto
├── go.mod                        # Módulo Go principal
├── go.sum
└── README.md
```