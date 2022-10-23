import { Field, ObjectType } from '@nestjs/graphql';
import { Schema, Prop, SchemaFactory } from '@nestjs/mongoose';
import { Document, ObjectId } from 'mongoose';

@ObjectType()
@Schema()
export class User extends Document {
    @Field(() => String)
    _id: ObjectId;

    @Prop()
    @Field(() => String)
    email: string;

    @Prop()
    @Field(() => String)
    name: string;

    @Prop()
    password: string;
}

export const UserSchema = SchemaFactory.createForClass(User);
