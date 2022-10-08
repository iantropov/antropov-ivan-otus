import { Body, Controller, HttpCode, Post, Req, Res, UseGuards } from '@nestjs/common';
import { AuthService } from './auth.service';
import { RegisterDto } from './dto/register.dto';
import { LocalAuthGuard } from './local-auth.guard';
import { LoggedInGuard } from './logged-in.guard';
import { Request, Response } from 'express';

@Controller('auth')
export class AuthController {
    constructor(private readonly authService: AuthService) {}

    @Post('register')
    async register(@Body() registerDto: RegisterDto) {
        return this.authService.register(registerDto);
    }

    @HttpCode(200)
    @UseGuards(LocalAuthGuard)
    @Post('login')
    async logIn() {}

    @UseGuards(LoggedInGuard)
    @Post('logout')
    async logOut(@Req() request: Request, @Res() response: Response) {
        request.logout(function (err) {
            if (!err) {
                response.status(200).send();
            } else {
                console.log(err);
                response.status(403);
            }
        });
    }
}
