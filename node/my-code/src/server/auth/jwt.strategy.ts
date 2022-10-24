import { ExtractJwt, Strategy } from 'passport-jwt';
import { PassportStrategy } from '@nestjs/passport';
import { Injectable, UnauthorizedException } from '@nestjs/common';

import { jwtConstants } from './constants';
import { AuthService } from './auth.service';
import { User } from '../users/entities/user.entity';
import { UsersService } from '../users/users.service';

@Injectable()
export class JwtStrategy extends PassportStrategy(Strategy) {
    constructor(private readonly authService: AuthService, private readonly usersService: UsersService) {
        super({
            usernameField: 'email',
            jwtFromRequest: ExtractJwt.fromAuthHeaderAsBearerToken(),
            ignoreExpiration: true,
            secretOrKey: jwtConstants.secret
        });
    }

    async validate(user: User): Promise<any> {
        // check for revoked token
        const refreshedUser = await this.usersService.findOne(user._id);
        return user._id === String(refreshedUser._id) && refreshedUser;
    }
}
