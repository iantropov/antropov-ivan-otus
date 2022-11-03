import { InputType } from "@nestjs/graphql";

@InputType()
export class CreateProblemInput {
    readonly summary: string;
    readonly description: string;
    readonly solution: string;
    readonly categoryIds: string[];
}
