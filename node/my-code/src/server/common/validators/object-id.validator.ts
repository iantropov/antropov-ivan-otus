import {
    ValidatorConstraint,
    ValidatorConstraintInterface,
    ValidationArguments
} from 'class-validator';
import { Types } from 'mongoose';

@ValidatorConstraint({ name: 'objectId', async: false })
export class ObjectId implements ValidatorConstraintInterface {
    validate(value: string) {
        return Types.ObjectId.isValid(value)
    }

    defaultMessage(args: ValidationArguments) {
        return `Value (${args.value}) is not ObjectID`;
    }
}
