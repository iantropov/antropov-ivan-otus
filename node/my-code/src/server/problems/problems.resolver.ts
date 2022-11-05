import { ValidationPipe } from '@nestjs/common';
import { Args, ID, Mutation, Query, Resolver } from '@nestjs/graphql';
import { CurrentUser } from '../auth/current-user.decorator';
import { GraphQLUser } from '../users/entities/user-graphql.entity';
import { CreateProblemInput } from './dto/create-problem.input';
import { SearchProblemsArgs } from './dto/search-problems.args';

import { UpdateProblemInput } from './dto/update-problem.input';
import { Problem } from './entities/problem.entity';
import { SearchProblemsResult } from './entities/search-problem-result.entity';
import { ProblemsService } from './problems.service';

@Resolver()
export class ProblemsResolver {
    constructor(private readonly problemsService: ProblemsService) {}

    @Query(() => [Problem], { name: 'problems' })
    async findAll() {
        return this.problemsService.findAll();
    }

    @Query(() => Problem, { name: 'problem' })
    async findOne(@Args('id', { type: () => ID }) id: string) {
        return this.problemsService.findOne(id);
    }

    @Query(() => SearchProblemsResult, { name: 'searchProblems' })
    async search(
        @CurrentUser() user: GraphQLUser,
        @Args(new ValidationPipe({ skipMissingProperties: true })) args: SearchProblemsArgs
    ) {
        return this.problemsService.search(user, args);
    }

    @Mutation(() => Problem, { name: 'createProblem' })
    async create(@Args('createProblemInput') createProblemInput: CreateProblemInput) {
        return this.problemsService.create(createProblemInput);
    }

    @Mutation(() => Problem, { name: 'updateProblem' })
    async update(
        @Args('id', { type: () => ID }) problemId: string,
        @Args('updateProblemInput') updateProblemInput: UpdateProblemInput
    ) {
        return this.problemsService.update(problemId, updateProblemInput);
    }

    @Mutation(() => Problem, { name: 'deleteProblem' })
    async delete(@Args('id', { type: () => ID }) problemId: string) {
        return this.problemsService.remove(problemId);
    }
}
