import { Test, TestingModule } from '@nestjs/testing';
import request from 'supertest';
import mongoose, { Types } from 'mongoose';
import { INestApplication } from '@nestjs/common';
import { MongoMemoryServer } from 'mongodb-memory-server';

import { ProblemsModule } from '../src/server/problems/problems.module';
import { CategoriesService } from '../src/server/problems/categories.service';
import { Category } from '../src/server/problems/entities/category.entity';
import { MongooseModule, MongooseModuleOptions } from '@nestjs/mongoose';
import { GraphQLModule } from '@nestjs/graphql';
import { ApolloDriver, ApolloDriverConfig } from '@nestjs/apollo';
import { join } from 'path';
import { UsersModule } from '../src/server/users/users.module';
import { AuthModule } from '../src/server/auth/auth.module';
import { ConfigModule } from '@nestjs/config';

const gql = '/graphql';

const mockCategory = (name = 'category'): Partial<Category> => ({
    name,
    _id: new Types.ObjectId()
});

const category = mockCategory();
const category1 = mockCategory('category-1');
const category2 = mockCategory('category-2');

const categoriesArray = [category, category1, category2];

let mongod: MongoMemoryServer;

let config = {
    PORT: '80',
    MONGODB_URI: 'mongodb',
    JWT_SECRET: 'secret',
    SESSION_SECRET: '9CwrzJ3QJsihZmXHqO85dAo'
};

export const rootMongooseTestModule = (options: MongooseModuleOptions = {}) =>
    MongooseModule.forRootAsync({
        useFactory: async () => {
            mongod = await MongoMemoryServer.create();
            const mongoUri = mongod.getUri();
            console.log(mongoUri);
            console.log(join(process.cwd(), '../src/server/*.gql'));
            config.MONGODB_URI = mongoUri; //==========> logs mongodb://127.0.0.1:37345/
            return {
                uri: mongoUri,
                ...options
            };
        }
    });

export const closeInMongodConnection = async () => {
    await mongoose.disconnect();
    if (mongod) await mongod.stop();
};

describe('GraphQL AppResolver (e2e) {Supertest}', () => {
    let app: INestApplication;
    let categoriesService = { findAll: () => categoriesArray };

    beforeAll(async () => {
        const moduleFixture: TestingModule = await Test.createTestingModule({
            imports: [
                await rootMongooseTestModule(),
                ConfigModule.forRoot({
                    load: [() => config]
                }),
                GraphQLModule.forRoot<ApolloDriverConfig>({
                    driver: ApolloDriver,
                    typePaths: ['./src/server/schema.gql']
                }),
                UsersModule,
                ProblemsModule,
                AuthModule
            ]
        })
            .overrideProvider(CategoriesService)
            .useValue(categoriesService)
            .compile();

        app = moduleFixture.createNestApplication();
        // await app.listen(80);
        await app.init();
    });

    afterAll(async () => {
        await app.close();
        await closeInMongodConnection();
    });

    describe(gql, () => {
        describe('categories', () => {
            it('should get the cats array', () => {
                return request(app.getHttpServer())
                    .post(gql)
                    .send({ query: '{categories {_id name }}' })
                    .expect(200)
                    .expect(res => {
                        expect(res.body.data.getCats).toEqual(categoriesArray);
                    });
            });
            // describe('one cat', () => {
            //   it('should get a single cat', () => {
            //     return request(app.getHttpServer())
            //       .post(gql)
            //       .send({ query: '{getCat(catId:{id:"2"}){id name age breed}}' })
            //       .expect(200)
            //       .expect((res) => {
            //         expect(res.body.data.getCat).toEqual({
            //           name: 'Terra',
            //           age: 5,
            //           breed: 'Siberian',
            //           id: '2',
            //         });
            //       });
            //   });
            //   it('should get an error for bad id', () => {
            //     return request(app.getHttpServer())
            //       .post(gql)
            //       .send({ query: '{getCat(catId: {id:"500"}){id name age breed}}' })
            //       .expect(200)
            //       .expect((res) => {
            //         expect(res.body.data).toBe(null);
            //         expect(res.body.errors[0].message).toBe(
            //           'No cat with id 500 found',
            //         );
            //       });
            //   });
            // });
            // it('should create a new cat and have it added to the array', () => {
            //   return (
            //     request(app.getHttpServer())
            //       .post(gql)
            //       .send({
            //         query:
            //           'mutation {insertCat(newCat: { name: "Vanitas", breed: "Calico", age: 100 }) {breed name id age}}',
            //       })
            //       .expect(200)
            //       .expect((res) => {
            //         expect(res.body.data.insertCat).toEqual({
            //           name: 'Vanitas',
            //           breed: 'Calico',
            //           age: 100,
            //           id: '4',
            //         });
            //       })
            //       // chain another request to see our original one works as expected
            //       .then(() =>
            //         request(app.getHttpServer())
            //           .post(gql)
            //           .send({ query: '{getCats {id name breed age}}' })
            //           .expect(200)
            //           .expect((res) => {
            //             expect(res.body.data.getCats).toEqual(
            //               cats.concat([
            //                 {
            //                   name: 'Vanitas',
            //                   breed: 'Calico',
            //                   age: 100,
            //                   id: '4',
            //                 },
            //               ]),
            //             );
            //           }),
            //       )
            //   );
            // });
        });
    });
});
