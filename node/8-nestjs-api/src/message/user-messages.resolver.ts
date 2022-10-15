import { Parent, ResolveField, Resolver } from '@nestjs/graphql';
import { User } from 'src/user/user.entity';
import { Message } from './message.entity';
import { MessageService } from './message.service';

@Resolver(() => User)
export class UserMessagesResolver {
    constructor(private readonly messageService: MessageService) {}

    @ResolveField('messages', () => [Message])
    getUserMessages(@Parent() user: User) {
        return this.messageService.findAll(user.id);
    }
}
