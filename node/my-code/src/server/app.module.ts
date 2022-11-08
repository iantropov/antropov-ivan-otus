import { Logger, MiddlewareConsumer, Module } from '@nestjs/common';
import { MongooseModule } from '@nestjs/mongoose';
import { GraphQLModule } from '@nestjs/graphql';
import { ApolloDriver, ApolloDriverConfig } from '@nestjs/apollo';
import { join } from 'path';
import passport from 'passport';
import mongoose from 'mongoose';
import * as util from 'node:util';
import { ConfigModule, ConfigService } from '@nestjs/config';

import { AppController } from './app.controller';
import { AppService } from './app.service';
import { ViewModule } from './view/view.module';
import { UsersModule } from './users/users.module';
import { LoggerMiddleware } from './common/middlewares/logger.middleware';
import { ProblemsModule } from './problems/problems.module';
import { AuthModule } from './auth/auth.module';

@Module({
    imports: [
        ConfigModule.forRoot({
            envFilePath: '.development.env',
            ignoreEnvFile: process.env.NODE_ENV === 'production'
        }),
        ViewModule,
        MongooseModule.forRootAsync({
            imports: [ConfigModule],
            useFactory: async (configService: ConfigService) => ({
                uri: configService.get<string>('MONGODB_URI')
            }),
            inject: [ConfigService]
        }),
        GraphQLModule.forRoot<ApolloDriverConfig>({
            driver: ApolloDriver,
            autoSchemaFile: join(process.cwd(), 'src/server/schema.gql')
        }),
        UsersModule,
        ProblemsModule,
        AuthModule
    ],
    controllers: [AppController],
    providers: [AppService]
})
export class AppModule {
    logger = new Logger('Mongoose');

    configure(consumer: MiddlewareConsumer): void {
        // consumer.apply(LoggerMiddleware).forRoutes('*');
        consumer.apply(LoggerMiddleware, passport.initialize(), passport.session()).forRoutes('*');

        mongoose.set('debug', (collectionName, methodName, ...methodArgs) => {
            const msgMapper = m => {
                return util
                    .inspect(m, false, 10, true)
                    .replace(/\n/g, '')
                    .replace(/\s{2,}/g, ' ');
            };
            this.logger.log(
                `${collectionName}.${methodName}(${methodArgs.map(msgMapper).join(', ')})`
            );
        });
    }
}
