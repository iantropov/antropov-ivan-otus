import { Module } from '@nestjs/common';
import { MongooseModule } from '@nestjs/mongoose';
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
        ])
    ],
    providers: [ProblemsService, ProblemsResolver]
})
export class ProblemsModule {}
