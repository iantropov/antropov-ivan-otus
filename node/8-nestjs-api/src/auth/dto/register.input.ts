import { InputType } from "@nestjs/graphql";
import { IsEmail, IsNotEmpty, Length } from "class-validator";

@InputType()
export class RegisterDto {
    @IsEmail()
    readonly email: string;

    @IsNotEmpty()
    @Length(5)
    password: string;
}
