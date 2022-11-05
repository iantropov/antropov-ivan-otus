import { Injectable, NotFoundException } from '@nestjs/common';
import { InjectModel } from '@nestjs/mongoose';
import { Model } from 'mongoose';

import { CreateProblemInput } from './dto/create-problem.input';
import { UpdateProblemInput } from './dto/update-problem.input';
import { Problem } from './entities/problem.entity';
import { UsersService } from '../users/users.service';
import { CategoriesService } from './categories.service';
import { SearchProblemsArgs } from './dto/search-problems.args';
import { GraphQLUser } from '../users/entities/user-graphql.entity';

const DEFAULT_LIMIT = 10;

@Injectable()
export class ProblemsService {
    constructor(
        @InjectModel(Problem.name) private readonly problemModel: Model<Problem>,
        private readonly usersService: UsersService,
        private readonly categoriesService: CategoriesService
    ) {}

    findAll() {
        return this.problemModel.find();
    }

    async search(user: GraphQLUser, args: SearchProblemsArgs) {
        let orSearchCriteria = [];
        if (args.text) {
            orSearchCriteria.push({
                summary: new RegExp(args.text, 'i')
            });
            orSearchCriteria.push({
                description: new RegExp(args.text, 'i')
            });
        }

        if (args.categoryIds && args.categoryIds.length > 0) {
            orSearchCriteria.push({
                'categories._id': { $in: args.categoryIds }
            });
        }

        let andSearchCriteria = [];
        if (args.favorites) {
            andSearchCriteria.push({
                _id: { $in: user.favorites }
            });
        }

        if (args.cursor) {
            andSearchCriteria.push({
                _id: { $gte: args.cursor }
            });
        }

        let filter = {};
        if (orSearchCriteria.length > 0) {
            filter = {
                ...filter,
                $or: orSearchCriteria
            };
        }
        if (andSearchCriteria.length > 0) {
            filter = {
                ...filter,
                $and: andSearchCriteria
            };
        }

        const limit = args.limit || DEFAULT_LIMIT;
        const problems = await this.problemModel
            .find(filter)
            .sort({ _id: 1 })
            .limit(limit + 1);

        return {
            edges: problems.length === limit + 1 ? problems.slice(0, -1) : problems,
            pageInfo: {
                cursor: problems[problems.length - 1]?._id,
                hasNextPage: problems.length === limit + 1
            }
        };
    }

    async findOne(id: string) {
        const problem = await this.problemModel.findOne({ _id: id });
        if (!problem) {
            throw new NotFoundException(`Problem #${id} not found`);
        }
        return problem;
    }

    async create(createProblemDto: CreateProblemInput) {
        const categories = await this.categoriesService.findByIds(createProblemDto.categoryIds);
        const problem = new this.problemModel({
            ...createProblemDto,
            categories
        });
        return problem.save();
    }

    async update(id: string, updateProblemDto: UpdateProblemInput) {
        const existingProblem = await this.problemModel.findOneAndUpdate(
            { _id: id },
            { $set: updateProblemDto },
            { new: true }
        );

        if (!existingProblem) {
            throw new NotFoundException(`Problem #${id} not found`);
        }
        return existingProblem;
    }

    async remove(id: string) {
        const problem = await this.findOne(id);
        await this.usersService.unlinkeProblemForAllUsers(id);
        return problem.remove();
    }
}
