import { MiddlewareConsumer, Module } from '@nestjs/common';
import { MongooseModule } from '@nestjs/mongoose';
import { GraphQLModule } from '@nestjs/graphql';
import { ApolloDriver, ApolloDriverConfig } from '@nestjs/apollo';
import { join } from 'path';
import passport from 'passport';

import { AppController } from './app.controller';
import { AppService } from './app.service';
import { ViewModule } from './view/view.module';
import { UsersModule } from './users/users.module';
import { LoggerMiddleware } from './common/middlewares/logger.middleware';
import { ProblemsModule } from './problems/problems.module';
import { AuthModule } from './auth/auth.module';

@Module({
    imports: [
        ViewModule,
        MongooseModule.forRoot('mongodb://localhost:27017/my-code'),
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
    configure(consumer: MiddlewareConsumer): void {
        // consumer.apply(LoggerMiddleware).forRoutes('*');
        consumer.apply(LoggerMiddleware, passport.initialize(), passport.session()).forRoutes('*');
    }
}
