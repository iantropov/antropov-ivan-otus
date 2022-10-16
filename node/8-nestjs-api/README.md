## Description

Plan:
    - Day 0:
      + Install GQL
      + Config GQL
      + Add a temporal entity for GraphQL
      + Add a response with GQL
      + Move current entities to GQL (without Auth)
    - Day 1:
      <- Field resolver for user.messages
      + DataLoader
      + Queries for Message
      + Queries for User
      + Mutations for Message
      + Mutations for User
      - Add Auth

NestJS API app (on GraphQL)

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
