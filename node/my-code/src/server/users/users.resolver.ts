import { Args, ID, Mutation, Query, Resolver } from '@nestjs/graphql';
import { CreateUserInput } from './input/create-user.input';

import { UpdateUserInput } from './input/update-user.input';
import { User } from './entities/user.entity';
import { UsersService } from './users.service';

@Resolver()
export class UsersResolver {
    constructor(private readonly usersService: UsersService) {}

    @Query(() => [User], { name: 'users' })
    async findAll() {
        return this.usersService.findAll();
    }

    @Query(() => User, { name: 'user' })
    async findOne(@Args('id', { type: () => ID }) id: string) {
        return this.usersService.findOne(id);
    }

    @Mutation(() => User, { name: 'createUser' })
    async create(
        @Args('createUserInput') createUserInput: CreateUserInput
    ) {
        return this.usersService.create(createUserInput);
    }


    @Mutation(() => User, { name: 'updateUser' })
    async update(
        @Args('userId', { type: () => ID }) userId: string,
        @Args('updateUserInput') updateUserInput: UpdateUserInput
    ) {
        return this.usersService.update(userId, updateUserInput);
    }

    @Mutation(() => User, { name: 'deleteUser' })
    async delete(@Args('userId', { type: () => ID }) userId: string) {
        return this.usersService.remove(userId);
    }
}
