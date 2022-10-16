import { ValidationPipe } from '@nestjs/common';
import { NestFactory } from '@nestjs/core';
import { AppModule } from './app.module';
import { HttpExceptionFilter } from './common/filters/http-exception.filter';
import cookieParser from 'cookie-parser';
import session from 'express-session';

async function bootstrap() {
    const app = await NestFactory.create(AppModule);
    app.useGlobalPipes(new ValidationPipe({
        transform: true,
        whitelist: true
    }))
    app.useGlobalFilters(new HttpExceptionFilter());
    app.use(cookieParser());
    app.use(
        session({
          secret: '9CwrzJ3QJsihZmXHqO85dAo',
          resave: false,
          saveUninitialized: false,
        }),
      );
    await app.listen(3000);
}
bootstrap();
