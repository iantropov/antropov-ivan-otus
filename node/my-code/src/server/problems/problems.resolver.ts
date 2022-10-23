import { Args, ID, Mutation, Query, Resolver } from '@nestjs/graphql';
import { CreateProblemInput } from './input/create-problem.input';

import { UpdateProblemInput } from './input/update-problem.input';
import { Problem } from './entities/problem.entity';
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

    @Mutation(() => Problem, { name: 'createProblem' })
    async create(
        @Args('createProblemInput') createProblemInput: CreateProblemInput
    ) {
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
