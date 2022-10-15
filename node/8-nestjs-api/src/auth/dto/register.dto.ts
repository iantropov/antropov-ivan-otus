import { IsEmail, IsNotEmpty, Length } from "class-validator";

export class RegisterDto {
    @IsEmail()
    readonly email: string;

    @IsNotEmpty()
    @Length(5)
    password: string;
}
