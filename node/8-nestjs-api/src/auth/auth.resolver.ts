import { Args, Context, Mutation, Resolver } from '@nestjs/graphql';
import { User } from 'src/user/user.entity';
import { RegisterDto } from './dto/register.input';
import { AuthService} from './auth.service'
import { UseGuards } from '@nestjs/common';
import { LocalAuthGraphQLGuard } from './local-auth.graphql.guard';
import { LoggedInGraphQLGuard } from './logged-in.graphql.guard';
import { Request } from 'express';

@Resolver()
export class AuthResolver {
    constructor(private readonly authService: AuthService){}

    @Mutation(() => User, { name: 'registerUser'})
    async registerUser(@Args('registerUserInput') registerUserInput: RegisterDto) {
        return this.authService.register(registerUserInput);
    }

    @UseGuards(LocalAuthGraphQLGuard)
    @Mutation(() => String, { name: 'loginUser'})
    async loginUser(@Args('email') email: string, @Args('password') password: string) {
        return 'OK';
    }

    @UseGuards(LoggedInGraphQLGuard)
    @Mutation(() => String, { name: 'logoutUser'})
    async logoutUser(@Context('req') request: Request) {
        await request.logout(() => {});
        return 'OK';
    }

}
