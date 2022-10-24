import { Schema, Prop, SchemaFactory } from '@nestjs/mongoose';
import { Document } from 'mongoose';

@Schema()
export class User extends Document {
    @Prop()
    email: string;

    @Prop()
    name: string;

    @Prop()
    password: string;
}

export const UserSchema = SchemaFactory.createForClass(User);
