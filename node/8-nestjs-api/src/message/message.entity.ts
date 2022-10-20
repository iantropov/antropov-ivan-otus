import { Field, ID, ObjectType } from "@nestjs/graphql"
import { User } from "src/user/user.entity"
import { Entity, PrimaryGeneratedColumn, Column, ManyToOne, JoinColumn } from "typeorm"

@ObjectType()
@Entity('messages')
export class Message {
    @Field(() => ID, { description: 'A unique identifier' })
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