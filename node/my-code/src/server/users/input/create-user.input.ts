import { InputType } from "@nestjs/graphql";
import { IsEmail, IsNotEmpty, Length } from "class-validator";

@InputType()
export class CreateUserInput {
    @IsEmail()
    readonly email: string;

    @IsNotEmpty()
    readonly name: string;

    @IsNotEmpty()
    @Length(5)
    readonly password: string;
}

