import { Args, ID, Mutation, Query, Resolver } from '@nestjs/graphql';
import { Types } from 'mongoose';
import { UseGuards } from '@nestjs/common';

import { CreateUserInput } from './input/create-user.input';
import { UpdateUserInput } from './input/update-user.input';
import { UsersService } from './users.service';
import { GraphQLUser } from './entities/user-graphql.entity';
import { CurrentUser } from '../auth/current-user.decorator';
import { LoggedInGraphQLGuard } from '../auth/logged-in.graphql.guard';
import { AdminRequiredGraphQLGuard } from '../auth/admin-required.graphql.guard';
import { ParseObjectIdPipe } from '../common/pipes/object-id.pipe';

@Resolver()
export class UsersResolver {
    constructor(private readonly usersService: UsersService) {}

    @UseGuards(LoggedInGraphQLGuard)
    @UseGuards(AdminRequiredGraphQLGuard)
    @Query(() => [GraphQLUser], { name: 'users' })
    findAll() {
        return this.usersService.findAll();
    }

    @UseGuards(LoggedInGraphQLGuard)
    @UseGuards(AdminRequiredGraphQLGuard)
    @Query(() => GraphQLUser, { name: 'user' })
    findOne(@Args('id', { type: () => ID }, ParseObjectIdPipe) id: Types.ObjectId) {
        return this.usersService.findOne(id);
    }

    @UseGuards(LoggedInGraphQLGuard)
    @UseGuards(AdminRequiredGraphQLGuard)
    @Mutation(() => GraphQLUser, { name: 'createUser' })
    create(@Args('createUserInput') createUserInput: CreateUserInput) {
        return this.usersService.create(createUserInput);
    }

    @UseGuards(LoggedInGraphQLGuard)
    @UseGuards(AdminRequiredGraphQLGuard)
    @Mutation(() => GraphQLUser, { name: 'updateUser' })
    update(
        @Args('id', { type: () => ID }, ParseObjectIdPipe) userId: Types.ObjectId,
        @Args('updateUserInput') updateUserInput: UpdateUserInput
    ) {
        return this.usersService.update(userId, updateUserInput);
    }

    @UseGuards(LoggedInGraphQLGuard)
    @UseGuards(AdminRequiredGraphQLGuard)
    @Mutation(() => GraphQLUser, { name: 'deleteUser' })
    delete(@Args('id', { type: () => ID }, ParseObjectIdPipe) userId: Types.ObjectId) {
        return this.usersService.remove(userId);
    }

    @UseGuards(LoggedInGraphQLGuard)
    @Mutation(() => GraphQLUser, { name: 'likeProblem' })
    like(
        @CurrentUser() user,
        @Args('id', { type: () => ID }, ParseObjectIdPipe) problemId: Types.ObjectId
    ) {
        return this.usersService.likeProblem(user, problemId);
    }

    @UseGuards(LoggedInGraphQLGuard)
    @Mutation(() => GraphQLUser, { name: 'unlikeProblem' })
    unlike(
        @CurrentUser() user,
        @Args('id', { type: () => ID }, ParseObjectIdPipe) problemId: Types.ObjectId
    ) {
        return this.usersService.unlikeProblem(user, problemId);
    }
}
