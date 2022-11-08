import { ArgsType } from '@nestjs/graphql';
import { Validate } from 'class-validator';

import { ObjectId } from '../../common/validators/object-id.validator';

@ArgsType()
export class SearchProblemsArgs {
    text?: string;

    @Validate(ObjectId, {
        each: true
    })
    categoryIds?: string[];

    favorites?: boolean;

    cursor?: string;

    limit?: number;
}
