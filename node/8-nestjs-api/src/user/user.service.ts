import { Injectable, NotFoundException } from '@nestjs/common';
import { InjectRepository } from '@nestjs/typeorm';
import { Repository } from 'typeorm';
import { CreateUserDto } from './dto/create-user.dto';
import { UpdateUserDto } from './dto/update-user.dto';
import { User } from './user.entity';

type SerializedUser = Omit<User, 'password'>;

@Injectable()
export class UserService {
    constructor(
        @InjectRepository(User)
        private readonly userRepository: Repository<User>
    ) {}

    async findAll() {
        const users = await this.userRepository.find();
        return users.map(this.serializeUser);
    }

    async findOne(id: number) {
        const user = await this.userRepository.findOne({ where: { id } });
        if (!user) {
            throw new NotFoundException(`User #${id} not found!`);
        }
        return this.serializeUser(user);
    }

    async findById(id: number) {
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
        return this.serializeUser(await this.userRepository.save(user));
    }

    async update(id: number, updateUserDto: UpdateUserDto) {
        const user = await this.userRepository.preload({
            id,
            ...updateUserDto
        });
        if (!user) {
            throw new NotFoundException(`User #${id} not found!`);
        }
        return this.serializeUser(await this.userRepository.save(user));
    }

    async remove(id: number) {
        const user = await this.findById(id);
        return this.userRepository.remove(user);
    }

    private serializeUser(user: User) {
        const { password, ...attributes } = user;
        return attributes;
    }
}
