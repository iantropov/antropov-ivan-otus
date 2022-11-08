import { Field, ObjectType } from '@nestjs/graphql';
import { Schema, Prop, SchemaFactory } from '@nestjs/mongoose';
import { Document, ObjectId } from 'mongoose';

import { Category, CategorySchema } from './category.entity';

@ObjectType()
@Schema({
    toJSON: {
        virtuals: true
    }
})
export class Problem extends Document {
    @Field(() => String)
    _id: ObjectId;

    @Prop()
    summary: string;

    @Prop()
    description: string;

    @Field(() => String, { nullable: true })
    @Prop()
    solution?: string;

    @Prop({ type: [CategorySchema], default: [] })
    categories: Category[];

    categoryIds: string[];
}

export const ProblemSchema = SchemaFactory.createForClass(Problem);

ProblemSchema.virtual('categoryIds').get(function (this: Problem) {
    return this.categories.map(({ _id }) => _id);
});
