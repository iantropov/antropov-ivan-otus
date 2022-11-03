import { MinLength } from 'class-validator';
import { ArgsType } from '@nestjs/graphql';

@ArgsType()
export class SearchProblemsArgs {
  @MinLength(3)
  text?: string;

  @MinLength(1)
  categoryIds?: string[];

  favorite?: boolean;
}