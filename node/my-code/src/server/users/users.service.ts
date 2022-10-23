import { Injectable, NotFoundException } from '@nestjs/common';
import { InjectModel } from '@nestjs/mongoose';
import { Model } from 'mongoose';

import { CreateUserInput } from './input/create-user.input';
import { UpdateUserInput } from './input/update-user.input';
import { User } from './entities/user.entity';

@Injectable()
export class UsersService {
    constructor(@InjectModel(User.name) private readonly userModel: Model<User>) {}

    findAll() {
        return this.userModel.find().exec();
    }

    async findOne(id: string) {
        const user = await this.userModel.findOne({ _id: id }).exec();
        if (!user) {
            throw new NotFoundException(`User #${id} not found`);
        }
        return user;
    }

    create(createUserDto: CreateUserInput) {
        const user = new this.userModel(createUserDto);
        return user.save();
    }

    async update(id: string, updateUserDto: UpdateUserInput) {
        const existingUser = await this.userModel
            .findOneAndUpdate({ _id: id }, { $set: updateUserDto }, { new: true })
            .exec();

        if (!existingUser) {
            throw new NotFoundException(`User #${id} not found`);
        }
        return existingUser;
    }

    async remove(id: string) {
        const user = await this.findOne(id);
        return user.remove();
    }
}
