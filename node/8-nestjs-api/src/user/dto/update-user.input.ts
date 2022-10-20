import { InputType } from "@nestjs/graphql";
import { IsEmail } from "class-validator";

@InputType()
export class UpdateUserDto {
    @IsEmail()
    readonly email?: string;
}
