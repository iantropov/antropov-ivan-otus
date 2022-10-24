import { Args, ID, Mutation, Query, Resolver } from '@nestjs/graphql';
import { CreateUserInput } from './input/create-user.input';

import { UpdateUserInput } from './input/update-user.input';
import { UsersService } from './users.service';
import { GraphQLUser } from './entities/user-graphql.entity';

@Resolver()
export class UsersResolver {
    constructor(private readonly usersService: UsersService) {}

    @Query(() => [GraphQLUser], { name: 'users' })
    async findAll() {
        return this.usersService.findAll();
    }

    @Query(() => GraphQLUser, { name: 'user' })
    async findOne(@Args('id', { type: () => ID }) id: string) {
        return this.usersService.findOne(id);
    }

    @Mutation(() => GraphQLUser, { name: 'createUser' })
    async create(
        @Args('createUserInput') createUserInput: CreateUserInput
    ) {
        return this.usersService.create(createUserInput);
    }

    @Mutation(() => GraphQLUser, { name: 'updateUser' })
    async update(
        @Args('id', { type: () => ID }) userId: string,
        @Args('updateUserInput') updateUserInput: UpdateUserInput
    ) {
        return this.usersService.update(userId, updateUserInput);
    }

    @Mutation(() => GraphQLUser, { name: 'deleteUser' })
    async delete(@Args('id', { type: () => ID }) userId: string) {
        return this.usersService.remove(userId);
    }
}
