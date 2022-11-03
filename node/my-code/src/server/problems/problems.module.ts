import { Module } from '@nestjs/common';
import { MongooseModule } from '@nestjs/mongoose';
import { UsersModule } from '../users/users.module';
import { CategoriesResolver } from './categories.resolver';
import { CategoriesService } from './categories.service';
import { Category, CategorySchema } from './entities/category.entity';
import { Problem, ProblemSchema } from './entities/problem.entity';
import { ProblemsResolver } from './problems.resolver';
import { ProblemsService } from './problems.service';

@Module({
    imports: [
        MongooseModule.forFeature([
            {
                name: Problem.name,
                schema: ProblemSchema
            },
            {
                name: Category.name,
                schema: CategorySchema
            }
        ]),
        UsersModule
    ],
    providers: [ProblemsService, ProblemsResolver, CategoriesService, CategoriesResolver]
})
export class ProblemsModule {}
