import { BadRequestException, Injectable, NotFoundException } from '@nestjs/common';
import { InjectRepository } from '@nestjs/typeorm';
import { QueryFailedError, Repository } from 'typeorm';
import { CreateUserDto } from './dto/create-user.dto';
import { UpdateUserDto } from './dto/update-user.dto';
import { UNIQUE_EMAIL_CONTRAINT, User } from './user.enity';

const UNIQUE_CONSTRAINT_ERROR = '23505';

@Injectable()
export class UserService {
    constructor(
        @InjectRepository(User)
        private readonly userRepository: Repository<User>
    ) {}

    findAll() {
        return this.userRepository.find();
    }

    async findOne(id: string) {
        const user = await this.userRepository.findOne({ where: { id: +id } });
        if (!user) {
            throw new NotFoundException(`User #${id} not found!`);
        }
        return user;
    }

    async create(createUserDto: CreateUserDto) {
        const user = this.userRepository.create(createUserDto);
        try {
            return await this.userRepository.save(user);
        } catch (error) {
            if (
                error instanceof QueryFailedError &&
                error.driverError.code === UNIQUE_CONSTRAINT_ERROR &&
                error.driverError.constraint === UNIQUE_EMAIL_CONTRAINT
            ) {
                throw new BadRequestException('Please, choose another email');
            } else {
                throw error;
            }
        }
    }

    async update(id: string, updateUserDto: UpdateUserDto) {
        const user = await this.userRepository.preload({
            id: +id,
            ...updateUserDto
        });
        if (!user) {
            throw new NotFoundException(`User #${id} not found!`);
        }
        return this.userRepository.save(user);
    }

    async remove(id: string) {
        const user = await this.findOne(id);
        return this.userRepository.remove(user);
    }
}
