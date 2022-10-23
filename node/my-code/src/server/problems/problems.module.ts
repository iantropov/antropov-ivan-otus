import { Module } from '@nestjs/common';
import { MongooseModule } from '@nestjs/mongoose';
import { Problem, ProblemSchema } from './entities/problem.entity';
import { ProblemsController } from './problems.controller';
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
    providers: [ProblemsService, ProblemsResolver],
    controllers: [ProblemsController]
})
export class ProblemsModule {}
