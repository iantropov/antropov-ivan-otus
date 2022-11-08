import { ValidationPipe } from '@nestjs/common';
import { ConfigService } from '@nestjs/config';
import { NestFactory } from '@nestjs/core';
import cookieParser from 'cookie-parser';
import session from 'express-session';

import { AppModule } from './app.module';
import { HttpExceptionFilter } from './common/filters/http-exception.filter';

async function bootstrap() {
    const app = await NestFactory.create(AppModule);
    const configService = app.get(ConfigService);
    app.useGlobalFilters(new HttpExceptionFilter());
    app.useGlobalPipes(new ValidationPipe({ skipMissingProperties: true }));
    app.use(cookieParser());
    app.use(
        session({
            secret: '9CwrzJ3QJsihZmXHqO85dAo',
            resave: false,
            saveUninitialized: false
        })
    );
    await app.listen(configService.get<string>('PORT'));
}
bootstrap();
