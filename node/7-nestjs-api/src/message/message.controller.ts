import { Body, Controller, Delete, Get, HttpCode, Param, Patch, Post } from '@nestjs/common';
import { CreateMessageDto } from './dto/create-message.dto';
import { UpdateMessageDto } from './dto/update-message.dto';
import { MessageService } from './message.service';

@Controller('users/:userId/messages')
export class MessageController {
    constructor(private readonly messageService: MessageService) {}

    @Get()
    findAll(@Param('userId') userId: string) {
        return this.messageService.findAll(userId);
    }

    @Get(':id')
    findOne(@Param('userId') userId: string, @Param('id') id: string) {
        return this.messageService.findOne(userId, id);
    }

    @Post()
    create(@Param('userId') userId: string, @Body() createMessageDto: CreateMessageDto) {
        return this.messageService.create(userId, createMessageDto);
    }

    @Patch(':id')
    update(
        @Param('userId') userId: string,
        @Param('id') id: string,
        @Body() updateMessageDto: UpdateMessageDto
    ) {
        return this.messageService.update(userId, id, updateMessageDto);
    }

    @Delete(':id')
    @HttpCode(204)
    remove(@Param('userId') userId: string, @Param('id') id: string) {
        return this.messageService.remove(userId, id);
    }
}
