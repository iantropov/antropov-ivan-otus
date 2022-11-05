import { Field, ObjectType } from '@nestjs/graphql';
import { Problem } from './problem.entity';

@ObjectType()
class PageInfo {
    @Field(() => String, { nullable: true })
    cursor: string;

    @Field(() => Boolean)
    hasNextPage: boolean;
}

@ObjectType()
export class SearchProblemsResult {
    @Field(() => [Problem])
    edges: Problem[];

    @Field(() => PageInfo)
    pageInfo: PageInfo;
}


