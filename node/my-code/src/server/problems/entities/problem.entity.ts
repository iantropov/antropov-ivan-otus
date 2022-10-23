import { Field, ObjectType } from '@nestjs/graphql';
import { Schema, Prop, SchemaFactory } from '@nestjs/mongoose';
import { Document, ObjectId } from 'mongoose';

@ObjectType()
@Schema()
export class Problem extends Document {
  @Field(() => String)
  _id: ObjectId;

  @Prop()
  summary: string

  @Prop()
  description: string

  @Prop()
  solution: string
}

export const ProblemSchema = SchemaFactory.createForClass(Problem);
