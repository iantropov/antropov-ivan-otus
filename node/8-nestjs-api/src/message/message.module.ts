import { Module } from '@nestjs/common';
import { TypeOrmModule } from '@nestjs/typeorm';
import { User } from 'src/user/user.entity';
import { UserModule } from 'src/user/user.module';
import { MessageController } from './message.controller';
import { Message } from './message.entity';
import { MessageResolver } from './message.resolver';
import { MessageService } from './message.service';
import { MessagesByUserLoader } from './messages-by-user.loader.ts';
import { UserMessagesResolver } from './user-messages.resolver';

@Module({
    imports: [TypeOrmModule.forFeature([Message, User]), UserModule],
    controllers: [MessageController],
    providers: [MessageService, UserMessagesResolver, MessagesByUserLoader, MessageResolver]
})
export class MessageModule {}
