import { Args, ID, Mutation, Query, Resolver } from '@nestjs/graphql';
import { CreateUserInput } from './input/create-user.input';

import { UpdateUserInput } from './input/update-user.input';
import { UsersService } from './users.service';
import { GraphQLUser } from './entities/user-graphql.entity';
import { CurrentUser } from '../auth/current-user.decorator';
import { UseGuards } from '@nestjs/common';
import { LoggedInGraphQLGuard } from '../auth/logged-in.graphql.guard';
import { AdminRequiredGraphQLGuard } from '../auth/admin-required.graphql.guard';
import { ParseObjectIdPipe } from '../common/pipes/object-id.pipe';

@Resolver()
export class UsersResolver {
    constructor(private readonly usersService: UsersService) {}

    @UseGuards(LoggedInGraphQLGuard)
    @UseGuards(AdminRequiredGraphQLGuard)
    @Query(() => [GraphQLUser], { name: 'users' })
    async findAll() {
        return this.usersService.findAll();
    }

    @UseGuards(LoggedInGraphQLGuard)
    @UseGuards(AdminRequiredGraphQLGuard)
    @Query(() => GraphQLUser, { name: 'user' })
    async findOne(@Args('id', { type: () => ID }, ParseObjectIdPipe) id: string) {
        return this.usersService.findOne(id);
    }

    @UseGuards(LoggedInGraphQLGuard)
    @UseGuards(AdminRequiredGraphQLGuard)
    @Mutation(() => GraphQLUser, { name: 'createUser' })
    async create(@Args('createUserInput') createUserInput: CreateUserInput) {
        return this.usersService.create(createUserInput);
    }

    @UseGuards(LoggedInGraphQLGuard)
    @UseGuards(AdminRequiredGraphQLGuard)
    @Mutation(() => GraphQLUser, { name: 'updateUser' })
    async update(
        @Args('id', { type: () => ID }, ParseObjectIdPipe) userId: string,
        @Args('updateUserInput') updateUserInput: UpdateUserInput
    ) {
        return this.usersService.update(userId, updateUserInput);
    }

    @UseGuards(LoggedInGraphQLGuard)
    @UseGuards(AdminRequiredGraphQLGuard)
    @Mutation(() => GraphQLUser, { name: 'deleteUser' })
    async delete(@Args('id', { type: () => ID }, ParseObjectIdPipe) userId: string) {
        return this.usersService.remove(userId);
    }

    @UseGuards(LoggedInGraphQLGuard)
    @Mutation(() => GraphQLUser, { name: 'likeProblem' })
    async like(
        @CurrentUser() user,
        @Args('id', { type: () => ID }, ParseObjectIdPipe) problemId: string
    ) {
        return this.usersService.likeProblem(user, problemId);
    }

    @UseGuards(LoggedInGraphQLGuard)
    @Mutation(() => GraphQLUser, { name: 'unlikeProblem' })
    async unlike(
        @CurrentUser() user,
        @Args('id', { type: () => ID }, ParseObjectIdPipe) problemId: string
    ) {
        return this.usersService.unlikeProblem(user, problemId);
    }
}
