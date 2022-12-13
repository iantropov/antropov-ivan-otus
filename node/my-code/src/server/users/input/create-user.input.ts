import { InputType } from "@nestjs/graphql";
import { IsEmail, Length } from "class-validator";

@InputType()
export class CreateUserInput {
    @IsEmail()
    readonly email: string;

    @Length(3, 100)
    readonly name: string;

    @Length(5)
    readonly password: string;
}

