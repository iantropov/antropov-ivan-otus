import { ParseIntPipe } from '@nestjs/common';
import { Args, ID, Query, Resolver } from '@nestjs/graphql';
import { User } from './user.entity';
import { UserService } from './user.service';

@Resolver()
export class UserResolver {
    constructor(private readonly userService: UserService) {}

    @Query(() => [User], { name: 'users' })
    async findAll() {
        return this.userService.findAll();
    }

    @Query(() => User, { name: 'user' })
    async findOne(@Args('id', { type: () => ID }, ParseIntPipe) id: number) {
        return this.userService.findById(id);
    }
}
