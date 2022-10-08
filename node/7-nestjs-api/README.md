## Description

NestJS API app

Plan:
- User module
  + Module
  + Controller
  + Entity
  + Service
  + DTO
  + TypeORM
  + CRUD
- Message module
  + CRUD
  + Relation to user
  + Validations
- Authz / Authn
  + Login
  + Logout
  + Access to messages

## Installation

```bash
$ npm install
```

## Running the app

```bash
# db
$ docker-compose up -d

# development
$ npm run start

# watch mode
$ npm run start:dev

# production mode
$ npm run start:prod
```

## License

Nest is [MIT licensed](LICENSE).
