my-echo-app/
│
├── cmd/                        # Entry point aplikasi
│   └── main.go                 # File utama untuk menjalankan aplikasi
│
├── config/                     # Konfigurasi aplikasi (misalnya, environment, config.json)
│   └── config.go
│
├── internal/                   # Logika aplikasi (paket yang hanya digunakan di dalam proyek ini)
│   ├── handlers/               # Handler untuk route HTTP
│   │   ├── user_handler.go     # Contoh handler untuk user
│   │   └── auth_handler.go     # Contoh handler untuk autentikasi
│   │
│   ├── models/                 # Struct untuk database atau representasi data
│   │   └── user.go             # Model User
│   │
│   ├── repositories/           # Abstraksi akses data (CRUD)
│   │   └── user_repository.go  # Repository untuk User
│   │
│   ├── services/               # Logika bisnis aplikasi
│   │   └── user_service.go     # Service untuk User
│   │
│   └── middlewares/            # Custom middleware untuk Echo
│       └── auth_middleware.go  # Middleware untuk autentikasi
│
├── migrations/                 # Skrip migrasi database
│   └── 0001_create_users.up.sql
│
├── pkg/                        # Paket umum yang bisa digunakan kembali di berbagai proyek
│   └── utils/                  # Utility functions
│       └── validator.go
│
├── routes/                     # Routing aplikasi
│   └── routes.go
│
├── .env                        # Variabel lingkungan untuk konfigurasi
├── go.mod                      # File Go modules
└── go.sum                      # Checksum untuk dependensi