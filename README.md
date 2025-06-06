# Backend - RESI

Estructura general del proyecto:

```lua
mi-proyecto-soa/
├── services/
│   ├── users/
│       ├── cmd/
│       │   └── main.go         
│       │         │   ├── internal/           
│       │   ├── api/              
│       │   │   └── user_handlers.go
│       │   ├── core/             
│       │   │   ├── domain/
│       │   │   │   └── user.go
│       │   │   └── user_service.go
│       │   └── repository/        
│       │       └── user_repository.go
│       └── pkg/                  
│           └── models/
│               └── user_shared_models.go
│
├── shared/
│   ├── config/                 
│   │   └── config.go
│   ├── database/                 
│   │   └── common_db.go
│   ├── utils/                 
│   │   └── helpers.go
│   └── proto/              
│       └── common.proto
├── go.mod                      
├── go.sum
└── README.md
```

