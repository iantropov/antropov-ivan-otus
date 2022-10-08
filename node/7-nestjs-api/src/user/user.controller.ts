import { Body, Controller, Delete, Get, HttpCode, Param, Patch, UseGuards } from '@nestjs/common';
import { LocalAuthGuard } from 'src/auth/local-auth.guard';
import { LoggedInGuard } from 'src/auth/logged-in.guard';
import { UpdateUserDto } from './dto/update-user.dto';
import { UserService } from './user.service';

@Controller('users')
export class UserController {
    constructor(private readonly userService: UserService) {}

    @Get()
    @UseGuards(LoggedInGuard)
    findAll() {
        return this.userService.findAll();
    }

    @Get(':id')
    @UseGuards(LoggedInGuard)
    findOne(@Param('id') id: number) {
        return this.userService.findOne(id);
    }

    @Patch(':id')
    @UseGuards(LoggedInGuard)
    update(@Param('id') id: number, @Body() updateUserDto: UpdateUserDto) {
        return this.userService.update(id, updateUserDto);
    }

    @Delete(':id')
    @HttpCode(204)
    @UseGuards(LoggedInGuard)
    remove(@Param('id') id: number) {
        this.userService.remove(id);
    }
}
