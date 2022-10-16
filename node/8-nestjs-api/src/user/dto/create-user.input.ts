import { InputType } from "@nestjs/graphql";

@InputType()
export class CreateUserDto {
    readonly email: string;
    readonly password: string;
}
