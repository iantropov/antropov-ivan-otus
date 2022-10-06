import { User } from "src/user/user.enity"
import { Entity, PrimaryGeneratedColumn, Column, ManyToOne } from "typeorm"

@Entity('messages')
export class Message {
    @PrimaryGeneratedColumn()
    id: number

    @Column()
    text: string

    @ManyToOne(() => User, (user) => user.messages)
    user: User
}