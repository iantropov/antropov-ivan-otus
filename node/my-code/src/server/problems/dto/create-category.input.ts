import { InputType } from '@nestjs/graphql';
import { Length } from 'class-validator';

@InputType()
export class CreateCategoryInput {
    @Length(3, 20)
    name: string;
}
