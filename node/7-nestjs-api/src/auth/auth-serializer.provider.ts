import { Injectable } from '@nestjs/common';
import { PassportSerializer } from '@nestjs/passport';
import { User } from '../user/user.enity';
import { UserService } from '../user/user.service';

@Injectable()
export class AuthSerializer extends PassportSerializer {
    constructor(private readonly userService: UserService) {
        super();
    }
    serializeUser(user: User, done: (err: Error, user: { id: number }) => void) {
        done(null, { id: user.id });
    }

    async deserializeUser(
        payload: { id: number; role: string },
        done: (err: Error, user: Omit<User, 'password'>) => void
    ) {
        const user = await this.userService.findById(payload.id);
        done(null, user);
    }
}
