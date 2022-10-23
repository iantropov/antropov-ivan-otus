import { Args, Context, Mutation, Query, Resolver } from '@nestjs/graphql';

import { User } from '../users/entities/user.entity';
import { RegisterUserInput } from './input/register-user.input';
import { AuthService } from './auth.service';
import { UseGuards } from '@nestjs/common';
import { Request } from 'express';
import { JwtAuthGraphqlGuard } from './jwt-auth-graphql.guard';
import { CurrentUser } from './current-user.decorator';
import { UsersService } from '../users/users.service';

@Resolver()
export class AuthResolver {
    constructor(private readonly authService: AuthService, private readonly usersService: UsersService) {}

    @Query(() => User, { name: 'whoAmI'})
    @UseGuards(JwtAuthGraphqlGuard)
    async getCurrentUser(@CurrentUser() user: User) {
        return this.usersService.findOne(user.id);
    }

    @Mutation(() => User, { name: 'registerUser' })
    async registerUser(@Args('registerUserInput') registerUserInput: RegisterUserInput) {
        return this.authService.register(registerUserInput);
    }

    // @UseGuards(LocalAuthGraphQLGuard)
    @Mutation(() => String, { name: 'loginUser' })
    async loginUser(@Args('email') email: string, @Args('password') password: string) {
        return 'OK';
    }

    // @UseGuards(LoggedInGraphQLGuard)
    @Mutation(() => String, { name: 'logoutUser' })
    async logoutUser(@Context('req') request: Request) {
        await request.logout(() => {});
        return 'OK';
    }
}
