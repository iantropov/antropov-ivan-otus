import { InputType } from "@nestjs/graphql";

@InputType()
export class UpdateUserInput {
    readonly email?: string;
    readonly name?: string;
}
