import { Module } from '@nestjs/common';
import { Coffee } from './coffee.entity';
import { CoffeeResolver } from './coffee.resolver';

@Module({
    providers: [CoffeeResolver, Coffee]
})
export class CoffeeModule {}
