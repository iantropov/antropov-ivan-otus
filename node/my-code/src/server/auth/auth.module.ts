import { Module } from '@nestjs/common';
import { ConfigModule, ConfigService } from '@nestjs/config';
import { JwtModule } from '@nestjs/jwt';
import { PassportModule } from '@nestjs/passport';

import { UsersModule } from '../users/users.module';
import { AuthSerializer } from './auth-serializer.provider';
import { AuthResolver } from './auth.resolver';
import { AuthService } from './auth.service';
import { JwtStrategy } from './jwt.strategy';
import { LocalStrategy } from './local.strategy';

@Module({
    imports: [
        ConfigModule,
        UsersModule,
        PassportModule.register({
            session: true
        }),
        JwtModule.registerAsync({
            imports: [ConfigModule],
            useFactory: async (configService: ConfigService) => ({
                secret: configService.get<string>('JWT_SECRET'),
                signOptions: { expiresIn: '60s' }
            }),
            inject: [ConfigService]
        })
    ],
    providers: [AuthService, AuthResolver, LocalStrategy, JwtStrategy, AuthSerializer]
})
export class AuthModule {}
