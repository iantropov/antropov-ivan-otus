import { InputType } from '@nestjs/graphql';
import { Length, Validate } from 'class-validator';
import { ObjectId } from 'src/server/common/validators/object-id.validator';

@InputType()
export class CreateProblemInput {
    @Length(3, 100)
    readonly summary: string;

    @Length(3, 200)
    readonly description: string;

    @Length(3, 200)
    readonly solution?: string;

    @Validate(ObjectId, {
        each: true
    })
    readonly categoryIds: string[];
}
