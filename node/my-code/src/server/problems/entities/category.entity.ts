import { Field, ObjectType } from '@nestjs/graphql';
import { Schema, Prop, SchemaFactory } from '@nestjs/mongoose';
import { Document, Types } from 'mongoose';

@ObjectType()
@Schema()
export class Category extends Document {
  @Field(() => String)
  _id: Types.ObjectId;

  @Prop()
  name: string
}

export const CategorySchema = SchemaFactory.createForClass(Category);
