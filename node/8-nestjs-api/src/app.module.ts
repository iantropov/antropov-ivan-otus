import { MiddlewareConsumer, Module } from '@nestjs/common';

import { UserModule } from './user/user.module';
import { MessageModule } from './message/message.module';
import { LoggerMiddleware } from './common/middlewares/logger.middleware';
import { DatabaseModule } from './typeorm.module';
import { AuthModule } from './auth/auth.module';
import * as passport from 'passport';
import { GraphQLModule } from '@nestjs/graphql';
import { ApolloDriver, ApolloDriverConfig } from '@nestjs/apollo';
import { join } from 'path';

@Module({
    imports: [
        UserModule,
        MessageModule,
        DatabaseModule,
        AuthModule,
        GraphQLModule.forRoot<ApolloDriverConfig>({
            driver: ApolloDriver,
            autoSchemaFile: join(process.cwd(), 'src/schema.gql')
        })
    ]
})
export class AppModule {
    configure(consumer: MiddlewareConsumer): void {
        consumer.apply(LoggerMiddleware, passport.initialize(), passport.session()).forRoutes('*');
    }
}
