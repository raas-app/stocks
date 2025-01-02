## Database Configuration

### Database Setup

1. Create database `stocks` in mysql
2. Install `golang-migrate`
```bash
go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
``` 
3. Run Migrations with `make up`
4. To down migrationn, run `make down`

### Migrations

- To create migration, run the following command.
```bash
make create NAME=<migration_name>
```
- To Up/Down migration upto specific version, run:
```bash
make up-to ID=<timestamp>
```