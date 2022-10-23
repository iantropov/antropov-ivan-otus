import { InputType } from "@nestjs/graphql";

@InputType()
export class UpdateProblemInput {
    readonly summary?: string;
    readonly description?: string;
    readonly solution?: string;
}
