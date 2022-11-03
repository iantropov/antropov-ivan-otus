import { Module } from '@nestjs/common';
import { MongooseModule } from '@nestjs/mongoose';
import { UsersModule } from '../users/users.module';
import { Problem, ProblemSchema } from './entities/problem.entity';
import { ProblemsResolver } from './problems.resolver';
import { ProblemsService } from './problems.service';

@Module({
    imports: [
        MongooseModule.forFeature([
            {
                name: Problem.name,
                schema: ProblemSchema
            }
        ]),
        UsersModule
    ],
    providers: [ProblemsService, ProblemsResolver]
})
export class ProblemsModule {}
