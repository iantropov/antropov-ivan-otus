import { ParseIntPipe } from '@nestjs/common';
import { Args, ID, Mutation, Query, Resolver } from '@nestjs/graphql';
import { CreateUserDto } from './dto/create-user.input';
import { UpdateUserDto } from './dto/update-user.input';
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

    @Mutation(() => User, { name: 'createUser' })
    async create(@Args('createUserInput') createUserInput: CreateUserDto) {
        console.log("CREATE_USER_INPUT", createUserInput)
        return this.userService.create(createUserInput);
    }

    @Mutation(() => User, { name: 'updateUser' })
    async update(
        @Args('userId', { type: () => ID }, ParseIntPipe) userId: number,
        @Args('updateUserInput') updateUserInput: UpdateUserDto
    ) {
        return this.userService.update(userId, updateUserInput);
    }

    @Mutation(() => User, { name: 'deleteUser' })
    async delete(@Args('userId', { type: () => ID }, ParseIntPipe) userId: number) {
        return this.userService.remove(userId);
    }
}
