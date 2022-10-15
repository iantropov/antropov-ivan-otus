import { Module } from '@nestjs/common';
import { TypeOrmModule } from '@nestjs/typeorm';
import { UserModule } from 'src/user/user.module';
import { MessageController } from './message.controller';
import { Message } from './message.entity';
import { MessageService } from './message.service';
import { UserMessagesResolver } from './user-messages.resolver';

@Module({
    imports: [TypeOrmModule.forFeature([Message]), UserModule],
    controllers: [MessageController],
    providers: [MessageService, UserMessagesResolver]
})
export class MessageModule {}
