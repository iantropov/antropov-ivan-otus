import { ArrayMinSize, Min, MinLength } from 'class-validator';
import { ArgsType } from '@nestjs/graphql';

@ArgsType()
export class SearchProblemsArgs {
  @MinLength(3)
  text?: string;

  @ArrayMinSize(1)
  categoryIds?: string[];

  favorites?: boolean;

  cursor?: string;

  @Min(1)
  limit?: number;
}