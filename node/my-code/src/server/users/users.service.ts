import { Injectable, NotFoundException } from '@nestjs/common';
import { InjectModel } from '@nestjs/mongoose';
import { Model } from 'mongoose';

import { CreateUserInput } from './input/create-user.input';
import { UpdateUserInput } from './input/update-user.input';
import { User } from './entities/user.entity';
import { GraphQLUser } from './entities/user-graphql.entity';

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

    async findByEmail(email: string) {
        const user = await this.userModel.findOne({ email }).exec();
        if (!user) {
            throw new NotFoundException(`User with email=${email} not found`);
        }
        return user;
    }

    create(createUserInput: CreateUserInput) {
        const user = new this.userModel(createUserInput);
        return user.save();
    }

    async update(id: string, updateUserDto: UpdateUserInput) {
        const existingUser = await this.userModel
            .findOneAndUpdate({ id }, { $set: updateUserDto }, { new: true })
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

    async likeProblem(user: GraphQLUser, problemId: string) {
        user.favorites.push(problemId);
        return this.userModel.findOneAndUpdate(
            { id: user._id },
            {
                $set: {
                    favorites: user.favorites
                }
            }
        );
    }

    async unlikeProblem(user: GraphQLUser, problemId: string) {
        const problemIdIdx = user.favorites.findIndex(id => problemId === id);
        if (problemIdIdx === -1) return Promise.resolve();
        user.favorites.splice(problemIdIdx, 1);
        return this.userModel.findOneAndUpdate(
            { id: user._id },
            {
                $set: {
                    favorites: user.favorites
                }
            }
        );
    }

    async unlinkeProblemForAllUsers(problemId: string) {
        const users = await this.findAll();
        return Promise.all(users.map(user => this.unlikeProblem(user, problemId)));
    }

    serialize(user: User): GraphQLUser {
        const { _id, email, name, isAdmin, favorites } = user;
        return { _id, email, name, isAdmin, favorites };
    }
}
