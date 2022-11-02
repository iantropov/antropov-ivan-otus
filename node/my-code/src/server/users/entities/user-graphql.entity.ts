import { Field, ObjectType } from '@nestjs/graphql';

@ObjectType()
export class GraphQLUser {
    @Field(() => String)
    _id: string;

    @Field(() => String)
    email: string;

    @Field(() => String)
    name: string;

    @Field(() => Boolean)
    isAdmin: boolean;
}
