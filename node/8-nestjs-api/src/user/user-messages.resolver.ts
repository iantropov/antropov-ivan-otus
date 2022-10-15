import { Parent, Query, ResolveField, Resolver } from '@nestjs/graphql';
import { InjectRepository } from '@nestjs/typeorm';
import { Message } from 'src/message/message.entity';
import { Repository } from 'typeorm';
import { User } from './user.entity';

@Resolver(() => User)
export class UserMessagesResolver {
    constructor(
        // ⚙️ Inject the Flavor Repository
        // @InjectRepository(Message)
        // private readonly messagesRepository: Repository<Message>
    ) {}

    @ResolveField('messages', () => [Message])
    async getMessagesOfUser(@Parent() user: User) {
        return [];
        // return this.messagesRepository
        //     .createQueryBuilder('message')
        //     .innerJoin('message.users', 'users', 'users.id = :userId', {
        //         coffeeId: user.id
        //     })
        //     .getMany();
    }
}
