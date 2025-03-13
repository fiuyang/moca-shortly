# Short URL API (Golang)

A simple and efficient URL shortener API built with Go and Fiber framework.

## 🚀 Features
- 🔗 Shorten long URLs into short, easy-to-share links.
- 📂 Retrieve original URLs from shortened links.
- 🔄 Automatic redirection from short URLs to the original URLs.
- 🗄️ PostgresSQL database with GORM ORM.
- 📖 API documentation with Swagger (OpenAPI).
- ⏳ Rate Limiting to prevent abuse and ensure fair usage.

# Tech Used
 ![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white) ![Postgres](https://img.shields.io/badge/postgres-%23316192.svg?style=for-the-badge&logo=postgresql&logoColor=white)![Swagger](https://img.shields.io/badge/-Swagger-%23Clojure?style=for-the-badge&logo=swagger&logoColor=white) ![Fiber Badge](https://img.shields.io/badge/Fiber-008ECF?logo=fiber&logoColor=fff&style=for-the-badge) ![RateLimiter Badge](https://img.shields.io/badge/RateLimiter-008ECF?logo=fiber&logoColor=fff&style=for-the-badge)
      
# Getting Start:
Before you running the program, make sure you've run this command:
```bash
 make install
```

### Run the program
```bash
 make dev
```

### Re-Init Docs Swagger
```bash
 make doc
```

### Create Table Sql
```bash
 make migration command={create_table_name or alter_table_name_add_email}
```

### Migrate Up
```bash
  make migrateUp
```

### Migrate Down
```bash
  make migrateDown
```

### Migrate Fix
```bash
  make migrateForce command={version}
```

### Migrate Drop
```bash
  make migrateDrop
```

### Check Docs Swagger
```bash
 http://localhost:3000/docs/index.html#/
```

The program will run on http://localhost:3000 
<!-- </> with 💛 by readMD (https://readmd.itsvg.in) -->