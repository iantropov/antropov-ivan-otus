import { MiddlewareConsumer, Module } from '@nestjs/common';
import { TypeOrmModule } from '@nestjs/typeorm';

import { UserModule } from './user/user.module';
import { MessageModule } from './message/message.module';
import { LoggerMiddleware } from './common/middlewares/logger.middleware';
import { DatabaseModule } from './typeorm.module';
import { AuthModule } from './auth/auth.module';
import * as passport from 'passport';

@Module({
    imports: [UserModule, MessageModule, DatabaseModule, AuthModule]
})
export class AppModule {
    configure(consumer: MiddlewareConsumer): void {
        consumer.apply(LoggerMiddleware, passport.initialize(), passport.session()).forRoutes('*');
    }
}
