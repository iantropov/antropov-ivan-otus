import { Module } from '@nestjs/common';
import { TypeOrmModule } from '@nestjs/typeorm';
import { UserModule } from 'src/user/user.module';
import { MessageController } from './message.controller';
import { Message } from './message.entity';
import { MessageService } from './message.service';

@Module({
    imports: [TypeOrmModule.forFeature([Message]), UserModule],
    controllers: [MessageController],
    providers: [MessageService]
})
export class MessageModule {}
