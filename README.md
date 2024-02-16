# Dcard-2024-Backend-Intern-Assignment

## Setup

### Environment

set environment variables in `env.yaml`

```yaml
name:
version:
port:
db_host:
db_port:
db_user:
db_password:
db_name:
```

### Migrations

```bash
# create migration
make migrate-create name=<migration-name>

# migrate up
make migrate-up DB_HOST=<db_host> DB_PORT=<db_port> DB_USER=<db_user> DB_PASSWORD=<db_password> DB_NAME=<db_name>

# migrate down
make migrate-down DB_HOST=<db_host> DB_PORT=<db_port> DB_USER=<db_user> DB_PASSWORD=<db_password> DB_NAME=<db_name>
```
