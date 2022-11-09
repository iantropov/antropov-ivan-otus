import { Test, TestingModule } from '@nestjs/testing';
import { getModelToken } from '@nestjs/mongoose';
import { Model, Types } from 'mongoose';

import { CategoriesService } from './categories.service';
import { Category } from './entities/category.entity';

const mockCategory = (name = 'category'): Partial<Category> => ({
    name,
    _id: new Types.ObjectId()
});

const category = mockCategory();
const category1 = mockCategory('category-1');
const category2 = mockCategory('category-2');

const categoriesArray = [category, category1, category2];

describe('CategoriesService', () => {
    let service: CategoriesService;
    let model: Model<Category>;

    beforeEach(async () => {
        const module: TestingModule = await Test.createTestingModule({
            providers: [
                CategoriesService,
                {
                    provide: getModelToken('Category'),
                    useValue: function () {}
                }
            ]
        }).compile();

        service = module.get<CategoriesService>(CategoriesService);
        model = module.get<Model<Category>>(getModelToken('Category'));
        model.find = jest.fn();
        model.findOne = jest.fn();
        model.prototype.save = jest.fn();
    });

    it('should be defined', () => {
        expect(service).toBeDefined();
    });

    afterEach(() => {
        jest.clearAllMocks();
    });

    it('should return all categories', async () => {
        jest.spyOn(model, 'find').mockReturnValue({
            exec: jest.fn().mockResolvedValueOnce(categoriesArray)
        } as any);

        const categories = await service.findAll();
        expect(categories).toEqual(categoriesArray);
    });

    it('should return all categories (by ids)', async () => {
        const spy = jest.spyOn(model, 'find').mockReturnValue({
            exec: jest.fn().mockResolvedValueOnce([category1, category2])
        } as any);

        const categories = await service.findByIds(['uuid-1', 'uuid-2']);
        expect(categories).toEqual([category1, category2]);
        expect(spy).toHaveBeenCalledWith({ _id: ['uuid-1', 'uuid-2'] });
    });

    it('should findOne by id', async () => {
        const spy = jest.spyOn(model, 'findOne').mockReturnValue({
            exec: jest.fn().mockResolvedValueOnce(category)
        } as any);

        const objectId = new Types.ObjectId();
        const foundCategory = await service.findOne(objectId);
        expect(foundCategory).toEqual(category);
        expect(spy).toHaveBeenCalledWith({ _id: objectId });
    });

    it('should throw error for duplicated name', async () => {
        jest.spyOn(model, 'findOne').mockReturnValue({
            exec: jest.fn().mockResolvedValueOnce(category)
        } as any);

        await expect(() => service.create('new-category')).rejects.toThrowError(
            'Category with such name=new-category already exists'
        );
    });

    it('should create new category', async () => {
        jest.spyOn(model, 'findOne').mockReturnValue({
            exec: jest.fn().mockResolvedValueOnce(null)
        } as any);

        jest.spyOn(model.prototype, 'save').mockResolvedValueOnce(category);

        const createdCategory = await service.create('new-category');
        expect(createdCategory).toEqual(category);
    });
});
