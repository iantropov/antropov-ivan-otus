import { Module } from '@nestjs/common';
import { JwtModule } from '@nestjs/jwt';
import { PassportModule } from '@nestjs/passport';
import { UsersModule } from '../users/users.module';
import { AuthSerializer } from './auth-serializer.provider';

import { AuthResolver } from './auth.resolver';
import { AuthService } from './auth.service';
import { jwtConstants } from './constants';
import { JwtStrategy } from './jwt.strategy';
import { LocalStrategy } from './local.strategy';

@Module({
    imports: [
        UsersModule,
        PassportModule.register({
            session: true
        }),
        JwtModule.register({
            secret: jwtConstants.secret,
            signOptions: { expiresIn: '60s' }
        })
    ],
    providers: [AuthService, AuthResolver, LocalStrategy, JwtStrategy, AuthSerializer]
})
export class AuthModule {}
