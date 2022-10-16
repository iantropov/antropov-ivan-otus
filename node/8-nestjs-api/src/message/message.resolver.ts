import { ParseIntPipe } from '@nestjs/common';
import { Args, ID, Mutation, Resolver } from '@nestjs/graphql';
import { CreateMessageDto } from './dto/create-message.input';
import { UpdateMessageDto } from './dto/update-message.input';
import { Message } from './message.entity';
import { MessageService } from './message.service';

@Resolver(() => Message)
export class MessageResolver {
    constructor(private readonly messageService: MessageService) {}

    @Mutation(() => Message, { name: 'createMessage' })
    async create(
        @Args('userId', { type: () => ID }, ParseIntPipe) userId: number,
        @Args('createMessageInput') createMessageInput: CreateMessageDto
    ) {
        return this.messageService.create(userId, createMessageInput);
    }

    @Mutation(() => Message, { name: 'updateMessage' })
    async update(
        @Args('userId', { type: () => ID }, ParseIntPipe) userId: number,
        @Args('messageId', { type: () => ID }, ParseIntPipe) messageId: number,
        @Args('updateMessageInput') updateMessageInput: UpdateMessageDto
    ) {
        return this.messageService.update(userId, messageId, updateMessageInput);
    }

    @Mutation(() => Message, { name: 'deleteMessage' })
    async delete(
        @Args('userId', { type: () => ID }, ParseIntPipe) userId: number,
        @Args('messageId', { type: () => ID }, ParseIntPipe) messageId: number
    ) {
        return this.messageService.remove(userId, messageId);
    }
}
