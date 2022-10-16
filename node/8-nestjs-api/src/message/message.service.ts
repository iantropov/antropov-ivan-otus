import { Injectable, NotFoundException } from '@nestjs/common';
import { InjectRepository } from '@nestjs/typeorm';
import { UserService } from 'src/user/user.service';
import { Repository } from 'typeorm';
import { CreateMessageDto } from './dto/create-message.input';
import { UpdateMessageDto } from './dto/update-message.input';
import { Message } from './message.entity';

@Injectable()
export class MessageService {
    constructor(
        @InjectRepository(Message)
        private readonly messageRepository: Repository<Message>,
        private readonly userService: UserService
    ) {}

    findAll(userId: number) {
        return this.messageRepository.find({
            where: {
                userId
            }
        });
    }

    async findOne(userId: number, id: number) {
        const message = await this.messageRepository.findOne({
            where: { id, userId }
        });
        if (!message) {
            throw new NotFoundException(`Message #${id} not found!`);
        }
        return message;
    }

    async create(userId: number, createMessageDto: CreateMessageDto) {
        const user = await this.userService.findOne(userId);
        const message = this.messageRepository.create({
            userId: user.id,
            ...createMessageDto
        });
        return await this.messageRepository.save(message);
    }

    async update(userId: number, id: number, updateMessageDto: UpdateMessageDto) {
        const message = await this.messageRepository.preload({
            id,
            userId,
            ...updateMessageDto
        });
        if (!message) {
            throw new NotFoundException(`Message #${id} not found!`);
        }
        return this.messageRepository.save(message);
    }

    async remove(userId: number, id: number) {
        const message = await this.findOne(userId, id);
        await this.messageRepository.remove(message);
        return {
            ...message,
            id
        };
    }
}
