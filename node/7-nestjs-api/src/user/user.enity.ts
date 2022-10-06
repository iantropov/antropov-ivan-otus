import { Message } from '../message/message.entity';
import { Column, Entity, OneToMany, PrimaryGeneratedColumn, Unique } from 'typeorm';

export const UNIQUE_EMAIL_CONTRAINT = 'unique_email_constraint';

@Entity('users')
export class User {
    @PrimaryGeneratedColumn()
    id: number;

    @Column()
    @Unique(UNIQUE_EMAIL_CONTRAINT, ['email'])
    email: string;

    @OneToMany(() => Message, message => message.user)
    messages: Message[];
}
