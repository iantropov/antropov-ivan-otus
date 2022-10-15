import { IsString, Length } from "class-validator";

export class CreateMessageDto {
    @IsString()
    @Length(5, 20)
    readonly text: string;
}
