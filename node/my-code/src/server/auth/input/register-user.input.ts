import { InputType } from '@nestjs/graphql';

import { CreateUserInput } from '../../users/input/create-user.input';

@InputType()
export class RegisterUserInput extends CreateUserInput {}
