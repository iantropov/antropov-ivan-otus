import { UserService } from 'src/user/user.service';
import { RegisterDto } from './dto/register.input';
import * as bcrypt from 'bcrypt';
import { UNIQUE_CONSTRAINT_ERROR } from 'src/database/constants';
import { BadRequestException, HttpException, HttpStatus, Injectable } from '@nestjs/common';
import { QueryFailedError } from 'typeorm';
import { UNIQUE_EMAIL_CONTRAINT, User } from 'src/user/user.entity';

@Injectable()
export class AuthService {
    constructor(private readonly userService: UserService) {}

    async register(registerDto: RegisterDto) {
        const hashedPassword = await bcrypt.hash(registerDto.password, 10);
        try {
            const createdUser = await this.userService.create({
                ...registerDto,
                password: hashedPassword
            });
            return createdUser;
        } catch (error) {
            if (
                error instanceof QueryFailedError &&
                error.driverError.code === UNIQUE_CONSTRAINT_ERROR &&
                error.driverError.constraint === UNIQUE_EMAIL_CONTRAINT
            ) {
                throw new BadRequestException('Please, choose another email');
            } else {
                throw error;
            }
        }
    }

    async getAuthenticatedUser(email: string, hashedPassword: string) {
        try {
            const user = await this.userService.findByEmail(email);
            const isPasswordMatching = await bcrypt.compare(hashedPassword, user.password);
            if (!isPasswordMatching) {
                throw new HttpException('Wrong credentials provided', HttpStatus.UNAUTHORIZED);
            }
            user.password = undefined;
            return user;
        } catch (error) {
            throw new HttpException('Wrong credentials provided', HttpStatus.UNAUTHORIZED);
        }
    }
}
