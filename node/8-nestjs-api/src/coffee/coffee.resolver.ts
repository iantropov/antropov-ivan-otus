import { Query, Resolver } from '@nestjs/graphql';
import { Coffee } from './coffee.entity';

@Resolver()
export class CoffeeResolver {
    @Query(() => [Coffee], { name: 'coffees' })
    async findAll() {
        return [];
    }
}
