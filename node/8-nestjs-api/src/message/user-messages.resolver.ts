import { ParseIntPipe } from '@nestjs/common';
import { Args, ID, Parent, ResolveField, Resolver } from '@nestjs/graphql';
import { User } from 'src/user/user.entity';
import { Message } from './message.entity';
import { MessageService } from './message.service';
import { MessagesByUserLoader } from './messages-by-user.loader.ts';

@Resolver(() => User)
export class UserMessagesResolver {
    constructor(private readonly messagesByUserLoader: MessagesByUserLoader, private readonly messageService: MessageService) {}

    @ResolveField('messages', () => [Message])
    getUserMessages(@Parent() user: User) {
        return this.messagesByUserLoader.load(user.id);
    }

    @ResolveField('message', () => Message)
    getUserMessage(@Parent() user: User, @Args('id', { type: () => ID }, ParseIntPipe) messageId: number) {
        return this.messageService.findOne(user.id, messageId);
    }
}
