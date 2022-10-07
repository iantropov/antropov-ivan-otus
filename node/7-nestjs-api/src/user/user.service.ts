import { Injectable, NotFoundException } from '@nestjs/common';
import { InjectRepository } from '@nestjs/typeorm';
import { Repository } from 'typeorm';
import { CreateUserDto } from './dto/create-user.dto';
import { UpdateUserDto } from './dto/update-user.dto';
import { User } from './user.enity';

@Injectable()
export class UserService {
    constructor(
        @InjectRepository(User)
        private readonly userRepository: Repository<User>
    ) {}

    findAll() {
        return this.userRepository.find();
    }

    async findOne(id: number) {
        const user = await this.userRepository.findOne({ where: { id } });
        if (!user) {
            throw new NotFoundException(`User #${id} not found!`);
        }
        return user;
    }

    async findByEmail(email: string) {
        const user = await this.userRepository.findOne({ where: { email } });
        if (!user) {
            throw new NotFoundException(`User with ${email} not found!`);
        }
        return user;
    }

    async create(createUserDto: CreateUserDto) {
        const user = this.userRepository.create(createUserDto);
        return await this.userRepository.save(user);
    }

    async update(id: number, updateUserDto: UpdateUserDto) {
        const user = await this.userRepository.preload({
            id,
            ...updateUserDto
        });
        if (!user) {
            throw new NotFoundException(`User #${id} not found!`);
        }
        return this.userRepository.save(user);
    }

    async remove(id: number) {
        const user = await this.findOne(id);
        return this.userRepository.remove(user);
    }
}
