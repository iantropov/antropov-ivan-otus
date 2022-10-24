import { Injectable } from '@nestjs/common';
import { PassportSerializer } from '@nestjs/passport';
import { User } from '../users/entities/user.entity';
import { UsersService } from '../users/users.service';

@Injectable()
export class AuthSerializer extends PassportSerializer {
    constructor(private readonly userService: UsersService) {
        super();
    }
    serializeUser(user: User, done: (err: Error, user: { id: string }) => void) {
        done(null, { id: user.id });
    }

    async deserializeUser(payload: { id: string }, done: (err: Error, user: User) => void) {
        const user = await this.userService.findOne(payload.id);
        done(null, user);
    }
}
