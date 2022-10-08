import { Body, Controller, Delete, Get, HttpCode, Param, Patch, Post, UseGuards } from '@nestjs/common';
import { LoggedInGuard } from 'src/auth/logged-in.guard';
import { CreateMessageDto } from './dto/create-message.dto';
import { UpdateMessageDto } from './dto/update-message.dto';
import { MessageService } from './message.service';

@Controller('users/:userId/messages')
export class MessageController {
    constructor(private readonly messageService: MessageService) {}

    @Get()
    @UseGuards(LoggedInGuard)
    findAll(@Param('userId') userId: number) {
        return this.messageService.findAll(userId);
    }

    @Get(':id')
    @UseGuards(LoggedInGuard)
    findOne(@Param('userId') userId: number, @Param('id') id: number) {
        return this.messageService.findOne(userId, id);
    }

    @Post()
    @UseGuards(LoggedInGuard)
    create(@Param('userId') userId: number, @Body() createMessageDto: CreateMessageDto) {
        return this.messageService.create(userId, createMessageDto);
    }

    @Patch(':id')
    @UseGuards(LoggedInGuard)
    update(
        @Param('userId') userId: number,
        @Param('id') id: number,
        @Body() updateMessageDto: UpdateMessageDto
    ) {
        return this.messageService.update(userId, id, updateMessageDto);
    }

    @Delete(':id')
    @HttpCode(204)
    @UseGuards(LoggedInGuard)
    remove(@Param('userId') userId: number, @Param('id') id: number) {
        return this.messageService.remove(userId, id);
    }
}
