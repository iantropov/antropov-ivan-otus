import { ParseIntPipe, UseGuards } from '@nestjs/common';
import { Args, ID, Mutation, Query, Resolver } from '@nestjs/graphql';
import { LoggedInGraphQLGuard } from 'src/auth/logged-in.graphql.guard';
import { UpdateUserDto } from './dto/update-user.input';
import { User } from './user.entity';
import { UserService } from './user.service';

@Resolver()
export class UserResolver {
    constructor(private readonly userService: UserService) {}

    @UseGuards(LoggedInGraphQLGuard)
    @Query(() => [User], { name: 'users' })
    async findAll() {
        return this.userService.findAll();
    }

    @UseGuards(LoggedInGraphQLGuard)
    @Query(() => User, { name: 'user' })
    async findOne(@Args('id', { type: () => ID }, ParseIntPipe) id: number) {
        return this.userService.findById(id);
    }

    @UseGuards(LoggedInGraphQLGuard)
    @Mutation(() => User, { name: 'updateUser' })
    async update(
        @Args('userId', { type: () => ID }, ParseIntPipe) userId: number,
        @Args('updateUserInput') updateUserInput: UpdateUserDto
    ) {
        return this.userService.update(userId, updateUserInput);
    }

    @UseGuards(LoggedInGraphQLGuard)
    @Mutation(() => User, { name: 'deleteUser' })
    async delete(@Args('userId', { type: () => ID }, ParseIntPipe) userId: number) {
        return this.userService.remove(userId);
    }
}
