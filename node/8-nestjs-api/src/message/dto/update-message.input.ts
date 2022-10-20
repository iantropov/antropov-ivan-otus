import { InputType } from "@nestjs/graphql";
import { IsString, Length } from "class-validator";

@InputType()
export class UpdateMessageDto {
    @IsString()
    @Length(5, 20)
    readonly text?: string;
}
