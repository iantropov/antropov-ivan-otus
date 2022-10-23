import { ObjectType } from '@nestjs/graphql';
import { Schema, Prop, SchemaFactory } from '@nestjs/mongoose';
import { Document, ObjectId } from 'mongoose';

@ObjectType()
@Schema()
export class User extends Document {
  @Prop()
  email: string;

  @Prop()
  name: string;
}

export const UserSchema = SchemaFactory.createForClass(User);
