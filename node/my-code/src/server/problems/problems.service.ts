import { Injectable, NotFoundException } from '@nestjs/common';
import { InjectModel } from '@nestjs/mongoose';
import { Model } from 'mongoose';

import { CreateProblemInput } from './input/create-problem.input';
import { UpdateProblemInput } from './input/update-problem.input';
import { Problem } from './entities/problem.entity';
import { UsersService } from '../users/users.service';
import { CategoriesService } from './categories.service';

@Injectable()
export class ProblemsService {
    constructor(
        @InjectModel(Problem.name) private readonly problemModel: Model<Problem>,
        private readonly usersService: UsersService,
        private readonly categoriesService: CategoriesService
    ) {}

    findAll() {
        return this.problemModel.find().exec();
    }

    async findOne(id: string) {
        const problem = await this.problemModel.findOne({ _id: id }).exec();
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
        const existingProblem = await this.problemModel
            .findOneAndUpdate({ _id: id }, { $set: updateProblemDto }, { new: true })
            .exec();

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
