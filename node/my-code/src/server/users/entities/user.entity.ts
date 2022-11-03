import { Schema, Prop, SchemaFactory } from '@nestjs/mongoose';
import { Document, SchemaTypes } from 'mongoose';

@Schema()
export class User extends Document {
    @Prop()
    email: string;

    @Prop()
    name: string;

    @Prop()
    password: string;

    @Prop({ default: false })
    isAdmin: boolean;

    @Prop({ default: [] })
    favorites: string[]
}

export const UserSchema = SchemaFactory.createForClass(User);
