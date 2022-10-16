import { Parent, ResolveField, Resolver } from '@nestjs/graphql';
import { User } from 'src/user/user.entity';
import { Message } from './message.entity';
import { MessagesByUserLoader } from './messages-by-user.loader.ts';

@Resolver(() => User)
export class UserMessagesResolver {
    constructor(private readonly messagesByUserLoader: MessagesByUserLoader) {}

    @ResolveField('messages', () => [Message])
    getUserMessages(@Parent() user: User) {
        return this.messagesByUserLoader.load(user.id);
    }
}
