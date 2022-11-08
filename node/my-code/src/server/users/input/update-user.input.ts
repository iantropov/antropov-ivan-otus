import { InputType } from "@nestjs/graphql";
import { IsEmail, IsNotEmpty, Length } from "class-validator";

@InputType()
export class UpdateUserInput {
    @IsEmail()
    readonly email?: string;

    @Length(3, 100)
    readonly name?: string;
}
