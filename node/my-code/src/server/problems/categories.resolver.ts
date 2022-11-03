import { Args, ID, Mutation, Query, Resolver } from '@nestjs/graphql';

import { CategoriesService } from './categories.service';
import { Category } from './entities/category.entity';

@Resolver()
export class CategoriesResolver {
    constructor(private readonly categoriesService: CategoriesService) {}

    @Query(() => [Category], { name: 'categories' })
    async findAll() {
        return this.categoriesService.findAll();
    }

    @Query(() => Category, { name: 'category' })
    async findOne(@Args('id', { type: () => ID }) id: string) {
        return this.categoriesService.findOne(id);
    }

    @Mutation(() => Category, { name: 'createCategory' })
    async create(@Args('name', { type: () => String! }) name: string) {
        return this.categoriesService.create(name);
    }
}
