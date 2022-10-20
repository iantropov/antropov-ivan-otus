import { Injectable, Scope } from '@nestjs/common';
import { InjectRepository } from '@nestjs/typeorm';
import DataLoader from 'dataloader';
import { User } from 'src/user/user.entity';
import { In, Repository } from 'typeorm';
import { Message } from './message.entity';

@Injectable({ scope: Scope.REQUEST })
export class MessagesByUserLoader extends DataLoader<number, Message[]> {
    constructor(@InjectRepository(User) private readonly usersRepository: Repository<User>) {
        super(keys => this.batchLoadFn(keys));
    }

    private async batchLoadFn(userIds: readonly number[]): Promise<Message[][]> {
        const usersWithMessages = await this.usersRepository.find({
            select: ['id'], // since we don't really need a coffee object here
            relations: ['messages'], // to fetch related flavors
            where: {
                id: In(userIds as number[]) // to make sure we only query requested coffees
            }
        });

        // to map an array of coffees two a 2-dimensional array of flavors where position in the array indicates to which coffee flavors belong
        return usersWithMessages.map(user => user.messages);
    }
}
