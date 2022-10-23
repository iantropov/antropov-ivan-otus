import { Injectable, NotFoundException } from '@nestjs/common';
import { InjectModel } from '@nestjs/mongoose';
import { Model } from 'mongoose';

import { CreateProblemInput } from './input/create-problem.input';
import { UpdateProblemInput } from './input/update-problem.input';
import { Problem } from './entities/problem.entity';

@Injectable()
export class ProblemsService {
    constructor(@InjectModel(Problem.name) private readonly problemModel: Model<Problem>) {}

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

    create(createProblemDto: CreateProblemInput) {
        const problem = new this.problemModel(createProblemDto);
        return problem.save();
    }

    async update(id: string, updateProblemDto: UpdateProblemInput) {
        const existingProblem = await this.problemModel
            .findOneAndUpdate({ id }, { $set: updateProblemDto }, { new: true })
            .exec();

        if (!existingProblem) {
            throw new NotFoundException(`Problem #${id} not found`);
        }
        return existingProblem;
    }

    async remove(id: string) {
        const problem = await this.findOne(id);
        return problem.remove();
    }
}
