import { UseGuards } from '@nestjs/common';
import { Args, ID, Mutation, Query, Resolver } from '@nestjs/graphql';
import { AdminRequiredGraphQLGuard } from '../auth/admin-required.graphql.guard';
import { LoggedInGraphQLGuard } from '../auth/logged-in.graphql.guard';
import { ParseObjectIdPipe } from '../common/pipes/object-id.pipe';

import { CategoriesService } from './categories.service';
import { CreateCategoryInput } from './dto/create-category.input';
import { Category } from './entities/category.entity';

@Resolver()
export class CategoriesResolver {
    constructor(private readonly categoriesService: CategoriesService) {}

    @UseGuards(LoggedInGraphQLGuard)
    @Query(() => [Category], { name: 'categories' })
    async findAll() {
        return this.categoriesService.findAll();
    }

    @UseGuards(LoggedInGraphQLGuard)
    @Query(() => Category, { name: 'category' })
    async findOne(@Args('id', { type: () => ID }, ParseObjectIdPipe) id: string) {
        return this.categoriesService.findOne(id);
    }

    @UseGuards(LoggedInGraphQLGuard)
    @UseGuards(AdminRequiredGraphQLGuard)
    @Mutation(() => Category, { name: 'createCategory' })
    async create(@Args('createCategoryInput') createCategoryInput: CreateCategoryInput) {
        return this.categoriesService.create(createCategoryInput.name);
    }
}
