import { IsEmail, IsNotEmpty, Length } from "class-validator";

export class CreateUserDto {
    @IsEmail()
    readonly email: string;

    @IsNotEmpty()
    @Length(5)
    password: string;
}
