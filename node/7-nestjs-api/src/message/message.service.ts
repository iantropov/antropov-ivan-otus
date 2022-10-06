import { Injectable, NotFoundException } from '@nestjs/common';
import { InjectRepository } from '@nestjs/typeorm';
import { User } from 'src/user/user.enity';
import { UserService } from 'src/user/user.service';
import { Repository } from 'typeorm';
import { CreateMessageDto } from './dto/create-message.dto';
import { UpdateMessageDto } from './dto/update-message.dto';
import { Message } from './message.entity';

@Injectable()
export class MessageService {
    constructor(
        @InjectRepository(Message)
        private readonly messageRepository: Repository<Message>,
        private readonly userService: UserService
    ) {}

    findAll(userId: string) {
        return this.messageRepository.find({
            where: {
                user: {
                    id: +userId
                }
            }
        });
    }

    async findOne(userId: string, id: string) {
        const message = await this.messageRepository.findOne({
            where: { id: +id, userId: +userId }
        });
        if (!message) {
            throw new NotFoundException(`Message #${id} not found!`);
        }
        return message;
    }

    async create(userId: string, createMessageDto: CreateMessageDto) {
        const user = await this.userService.findOne(userId);
        const message = this.messageRepository.create({
            userId: user.id,
            ...createMessageDto
        });
        return await this.messageRepository.save(message);
    }

    async update(userId: string, id: string, updateMessageDto: UpdateMessageDto) {
        const message = await this.messageRepository.preload({
            id: +id,
            ...updateMessageDto
        });
        if (!message) {
            throw new NotFoundException(`Message #${id} not found!`);
        }
        return this.messageRepository.save(message);
    }

    async remove(userId: string, id: string) {
        const message = await this.findOne(userId, id);
        return this.messageRepository.remove(message);
    }
}
