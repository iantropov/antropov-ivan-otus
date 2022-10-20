import { Module } from '@nestjs/common';
import { AuthService } from './auth.service';
import { AuthController } from './auth.controller';
import { UserModule } from 'src/user/user.module';
import { PassportModule } from '@nestjs/passport';
import { LocalStrategy } from './local.strategy';
import { AuthSerializer } from './auth-serializer.provider';
import { AuthResolver } from './auth.resolver';

@Module({
    imports: [
        UserModule,
        PassportModule.register({
            session: true
        })
    ],
    providers: [AuthService, AuthSerializer, LocalStrategy, AuthResolver],
    controllers: [AuthController]
})
export class AuthModule {}
