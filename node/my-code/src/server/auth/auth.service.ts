import * as bcrypt from 'bcrypt';
import { Injectable, UnauthorizedException } from '@nestjs/common';
import { JwtService } from '@nestjs/jwt';

import { UsersService } from '../users/users.service';
import { RegisterUserInput } from './input/register-user.input';

@Injectable()
export class AuthService {
    constructor(
        private readonly usersService: UsersService,
        private readonly jwtService: JwtService
    ) {}

    async register(registerUserInput: RegisterUserInput) {
        const hashedPassword = await bcrypt.hash(registerUserInput.password, 10);
        const createdUser = await this.usersService.create({
            ...registerUserInput,
            password: hashedPassword
        });
        return createdUser;
    }

    // async getAuthenticatedUser(email: string, hashedPassword: string) {
    //     try {
    //         const user = await this.usersService.findByEmail(email);
    //         const isPasswordMatching = await bcrypt.compare(hashedPassword, user.password);
    //         if (!isPasswordMatching) {
    //             throw new UnauthorizedException();
    //         }
    //         user.password = undefined;
    //         return user;
    //     } catch (error) {
    //         throw new UnauthorizedException();
    //     }
    // }

    async login(user: any) {
        const payload = { username: user.username, sub: user.userId };
        return {
            access_token: this.jwtService.sign(payload)
        };
    }
}
