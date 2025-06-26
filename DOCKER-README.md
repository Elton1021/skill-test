# Docker Setup for School Management System

## Quick Start

1. **Start all services:**

   ```bash
   docker-compose up -d
   ```

2. **Access the application:**

   - Frontend: http://localhost:5173
   - Backend API: http://localhost:5007
   - Go Microservice API: http://localhost:8080
   - PostgreSQL: localhost:5432

3. **Stop services:**
   ```bash
   docker-compose down
   ```

## Services

- **PostgreSQL** (port 5432) - Database
- **Backend** (port 5007) - Node.js API
- **Go-Service** (port 8080) - Go Microservice
- **Frontend** (port 5173) - React App

## Development

- View logs: `docker-compose logs -f [service_name]`
- Restart service: `docker-compose restart [service_name]`
- Restore backup: `docker exec -i school_mgmt_postgres psql -U postgres school_mgmt < seed_db/backup_20250626_222719.sql`

## Database

- Database: `school_mgmt`
- User: `postgres`
- Password: `postgres123`
- Automatically initialized with seed data
