import { ArgsType } from '@nestjs/graphql';

@ArgsType()
export class SearchProblemsArgs {
  text?: string;
  categoryIds?: string[];
  favorites?: boolean;
  cursor?: string;
  limit?: number;
}