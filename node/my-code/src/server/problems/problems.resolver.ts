import { UseGuards } from '@nestjs/common';
import { Args, ID, Mutation, Query, Resolver } from '@nestjs/graphql';
import { Types } from 'mongoose';

import { AdminRequiredGraphQLGuard } from '../auth/admin-required.graphql.guard';
import { CurrentUser } from '../auth/current-user.decorator';
import { LoggedInGraphQLGuard } from '../auth/logged-in.graphql.guard';
import { ParseObjectIdPipe } from '../common/pipes/object-id.pipe';
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

    @UseGuards(LoggedInGraphQLGuard)
    @Query(() => [Problem], { name: 'problems' })
    async findAll() {
        return this.problemsService.findAll();
    }

    @UseGuards(LoggedInGraphQLGuard)
    @Query(() => Problem, { name: 'problem' })
    async findOne(@Args('id', { type: () => ID }, ParseObjectIdPipe) id: Types.ObjectId) {
        return this.problemsService.findOne(id);
    }

    @UseGuards(LoggedInGraphQLGuard)
    @Query(() => SearchProblemsResult, { name: 'searchProblems' })
    async search(@CurrentUser() user: GraphQLUser, @Args() args: SearchProblemsArgs) {
        return this.problemsService.search(user, args);
    }

    @UseGuards(LoggedInGraphQLGuard)
    @UseGuards(AdminRequiredGraphQLGuard)
    @Mutation(() => Problem, { name: 'createProblem' })
    async create(@Args('createProblemInput') createProblemInput: CreateProblemInput) {
        return this.problemsService.create(createProblemInput);
    }

    @UseGuards(LoggedInGraphQLGuard)
    @UseGuards(AdminRequiredGraphQLGuard)
    @Mutation(() => Problem, { name: 'updateProblem' })
    async update(
        @Args('id', { type: () => ID }, ParseObjectIdPipe) problemId: Types.ObjectId,
        @Args('updateProblemInput') updateProblemInput: UpdateProblemInput
    ) {
        return this.problemsService.update(problemId, updateProblemInput);
    }

    @UseGuards(LoggedInGraphQLGuard)
    @UseGuards(AdminRequiredGraphQLGuard)
    @Mutation(() => Problem, { name: 'deleteProblem' })
    async delete(@Args('id', { type: () => ID }, ParseObjectIdPipe) problemId: Types.ObjectId) {
        return this.problemsService.remove(problemId);
    }
}
