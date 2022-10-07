import { Body, Controller, HttpCode, Post, Req, UseGuards } from '@nestjs/common';
import { AuthService } from './auth.service';
import { RegisterDto } from './dto/register.dto';
import { LocalAuthGuard } from './local-auth.guard';
import RequestWithUser from './request-with-user.interface';

@Controller('auth')
export class AuthController {
    constructor(
        private readonly authService: AuthService
      ) {}

      @Post('register')
      async register(@Body() registerDto: RegisterDto) {
        return this.authService.register(registerDto);
      }

      @HttpCode(200)
      @UseGuards(LocalAuthGuard)
      @Post('login')
      async logIn(@Req() request: RequestWithUser) {
        // const user = request.user;
        // user.password = undefined;
        // return user;
        return request.session;
      }
}
