import { BadRequestException, Injectable, NotFoundException } from '@nestjs/common';
import { InjectModel } from '@nestjs/mongoose';
import { Model, Types } from 'mongoose';

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

    async findOne(id: Types.ObjectId) {
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

    async create(createUserInput: CreateUserInput) {
        const existingUser = await this.userModel.findOne({ email: createUserInput.email }).exec();
        if (existingUser) {
            throw new BadRequestException(
                `User with such email=${createUserInput.email} already exists`
            );
        }

        const user = new this.userModel(createUserInput);
        return user.save();
    }

    async update(id: Types.ObjectId, updateUserDto: UpdateUserInput) {
        const currentUser = await this.findOne(id);

        if (updateUserDto.email) {
            const existingUserByEmail = await this.userModel
                .findOne({ email: updateUserDto.email })
                .exec();
            if (existingUserByEmail && !existingUserByEmail._id.equals(currentUser._id)) {
                throw new BadRequestException(
                    `User with such email=${updateUserDto.email} already exists`
                );
            }
        }

        return this.userModel.findOneAndUpdate({ _id: id }, { $set: updateUserDto }, { new: true });
    }

    async remove(id: Types.ObjectId) {
        const user = await this.findOne(id);
        return user.remove();
    }

    likeProblem(user: User, problemId: Types.ObjectId) {
        user.favorites.push(problemId);
        return this.updateUserFavorites(user);
    }

    unlikeProblem(user: User, problemId: Types.ObjectId) {
        const problemIdIdx = user.favorites.findIndex(id => problemId.equals(id));
        if (problemIdIdx === -1) return Promise.resolve(user);
        user.favorites.splice(problemIdIdx, 1);
        return this.updateUserFavorites(user);
    }

    async unlinkProblemForAllUsers(problemId: Types.ObjectId) {
        const users = await this.findAll();
        return Promise.all(users.map(user => this.unlikeProblem(user, problemId)));
    }

    private updateUserFavorites(user: User) {
        return this.userModel.findOneAndUpdate(
            { _id: user._id },
            {
                $set: {
                    favorites: user.favorites
                }
            },
            { new: true }
        );
    }

    serialize(user: User): GraphQLUser {
        const { _id, email, name, isAdmin, favorites } = user;
        return { _id, email, name, isAdmin, favorites: favorites.map(String) };
    }
}
