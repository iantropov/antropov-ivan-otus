import { ObjectType } from '@nestjs/graphql';

@ObjectType()
export class Coffee {
    // 👈 make sure Entity is removed
    id: number;
    name: string;
    brand: string;
    flavors: string[];
}
