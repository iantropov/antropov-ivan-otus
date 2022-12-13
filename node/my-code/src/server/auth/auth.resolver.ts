import { Args, Context, Field, Mutation, ObjectType, Query, Resolver } from '@nestjs/graphql';

import { User } from '../users/entities/user.entity';
import { RegisterUserInput } from './input/register-user.input';
import { AuthService } from './auth.service';
import { Logger, UseGuards } from '@nestjs/common';
import { Request } from 'express';
import { CurrentUser } from './current-user.decorator';
import { UsersService } from '../users/users.service';
import { GraphQLUser } from '../users/entities/user-graphql.entity';
import { LocalAuthGraphQLGuard } from './local-auth-graphql.guard';
import { LoggedInGraphQLGuard } from './logged-in.graphql.guard';

@ObjectType()
class AccessToken {
    @Field(() => String)
    accessToken: string;
}

@Resolver()
export class AuthResolver {
    private readonly logger = new Logger('Auth');

    constructor(
        private readonly authService: AuthService,
        private readonly usersService: UsersService
    ) {}

    @Query(() => GraphQLUser, { name: 'whoAmI', nullable: true })
    async getCurrentUser(@Context('req') request) {
        if (!request || !request['user']) {
            return null;
        }

        try {
            return this.usersService.findOne(request['user']._id);
        } catch (error) {
            this.logger.error(`Can't find current user: ${error}`);
            return null;
        }
    }

    @Mutation(() => GraphQLUser, { name: 'registerUser' })
    async registerUser(@Args('registerUserInput') registerUserInput: RegisterUserInput) {
        return this.authService.register(registerUserInput);
    }

    @UseGuards(LocalAuthGraphQLGuard)
    @Mutation(() => AccessToken, { name: 'loginUser' })
    async loginUser(
        @Args('email') email: string,
        @Args('password') password: string,
        @CurrentUser() user: User
    ) {
        return this.authService.login(user);
    }

    @UseGuards(LoggedInGraphQLGuard)
    @Mutation(() => String, { name: 'logoutUser' })
    async logoutUser(@Context('req') request: Request) {
        await request.logout(() => {});
        return 'OK';
    }
}
