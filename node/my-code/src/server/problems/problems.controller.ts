import { Body, Controller, Delete, Get, HttpCode, Param, Patch, Post } from '@nestjs/common';
import { CreateProblemInput } from './input/create-problem.input';
import { UpdateProblemInput } from './input/update-problem.input';
import { ProblemsService } from './problems.service';

@Controller('/problems')
export class ProblemsController {
    constructor(private readonly problemsService: ProblemsService) {}

    @Get()
    findAll() {
        return this.problemsService.findAll();
    }

    @Get(':id')
    findOne(@Param('id') id: string) {
        return this.problemsService.findOne(id);
    }

    @Post()
    create(@Body() createProblemDto: CreateProblemInput) {
        return this.problemsService.create(createProblemDto);
    }

    @Patch(':id')
    update(@Param('id') id: string, @Body() updateProblemDto: UpdateProblemInput) {
        return this.problemsService.update(id, updateProblemDto);
    }

    @Delete(':id')
    @HttpCode(204)
    remove(@Param('id') id: string) {
        this.problemsService.remove(id);
    }
}
