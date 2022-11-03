import { Injectable, NotFoundException } from '@nestjs/common';
import { InjectModel } from '@nestjs/mongoose';
import { Model } from 'mongoose';

import { Category } from './entities/category.entity';

@Injectable()
export class CategoriesService {
    constructor(@InjectModel(Category.name) private readonly categoryModel: Model<Category>) {}

    findAll() {
        return this.categoryModel.find().exec();
    }

    findByIds(categoryIds: string[]) {
        return this.categoryModel.find({ _id: categoryIds }).exec();
    }

    async findOne(id: string) {
        const category = await this.categoryModel.findOne({ _id: id }).exec();
        if (!category) {
            throw new NotFoundException(`Category #${id} not found`);
        }
        return category;
    }

    create(name: string) {
        const category = new this.categoryModel({ name });
        return category.save();
    }
}
