import { User } from "src/user/user.enity"
import { Entity, PrimaryGeneratedColumn, Column, ManyToOne, JoinColumn } from "typeorm"

@Entity('messages')
export class Message {
    @PrimaryGeneratedColumn()
    id: number

    @Column()
    text: string

    @Column()
    userId: number

    @ManyToOne(() => User, (user) => user.messages, {
        onDelete: 'CASCADE'
    })
    @JoinColumn()
    user: User
}