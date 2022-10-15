import { IsString, Length } from "class-validator";

export class UpdateMessageDto {
    @IsString()
    @Length(5, 20)
    readonly text?: string;
}
