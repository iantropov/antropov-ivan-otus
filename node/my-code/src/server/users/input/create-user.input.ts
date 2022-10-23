import { InputType } from "@nestjs/graphql";

@InputType()
export class CreateUserInput {
    readonly email: string;
    readonly name: string;
}
