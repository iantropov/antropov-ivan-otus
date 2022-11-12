import { Test, TestingModule } from '@nestjs/testing';
import request from 'supertest';
import mongoose, { Types } from 'mongoose';
import { INestApplication } from '@nestjs/common';
import { MongoMemoryServer } from 'mongodb-memory-server';
import { ConfigModule } from '@nestjs/config';
import { GraphQLModule } from '@nestjs/graphql';
import { MongooseModule, MongooseModuleOptions } from '@nestjs/mongoose';
import { ApolloDriver, ApolloDriverConfig } from '@nestjs/apollo';

import { ProblemsModule } from '../src/server/problems/problems.module';
import { CategoriesService } from '../src/server/problems/categories.service';
import { UsersModule } from '../src/server/users/users.module';
import { AuthModule } from '../src/server/auth/auth.module';
import { LoggedInGraphQLGuard } from '../src/server/auth/logged-in.graphql.guard';

const GQL_URL = '/graphql';

const mockCategory = (name = 'category') => ({
    name,
    _id: String(new Types.ObjectId())
});

const category = mockCategory();
const category1 = mockCategory('category-1');
const category2 = mockCategory('category-2');

const categoriesArray = [category, category1, category2];

let mongod: MongoMemoryServer;

let testConfig = {
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
            testConfig.MONGODB_URI = mongoUri;
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

describe('GraphQL CategoriesResolver (e2e)', () => {
    let app: INestApplication;
    let categoriesService = { findAll: () => categoriesArray };
    let loggedInGuard = { canActivate: () => true };

    beforeAll(async () => {
        const moduleFixture: TestingModule = await Test.createTestingModule({
            imports: [
                await rootMongooseTestModule(),
                ConfigModule.forRoot({
                    load: [() => testConfig]
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
            .overrideGuard(LoggedInGraphQLGuard)
            .useValue(loggedInGuard)
            .compile();

        app = moduleFixture.createNestApplication();
        await app.init();
    });

    afterAll(async () => {
        await app.close();
        await closeInMongodConnection();
    });

    describe(GQL_URL, () => {
        describe('categories', () => {
            it('should get the categories array', () => {
                return request(app.getHttpServer())
                    .post(GQL_URL)
                    .send({ query: '{categories {_id name }}' })
                    .expect(200)
                    .expect(res => {
                        expect(res.body.data.categories).toEqual(categoriesArray);
                    });
            });
        });
    });
});
