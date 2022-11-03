export interface User {
    _id: string;
    email: string;
    name: string;
    isAdmin: boolean;
    favorites: string[];
}

export interface Category {
    _id: string;
    name: string;
}

export interface Problem {
    _id: string;
    summary: string;
    description: string;
    solution: string;
    categories: Category[];
    categoryIds: string[];
}

export type ProblemData = Omit<Problem, '_id' | 'categories'>;

export interface WhoAmIResponse {
    whoAmI: User;
}

export interface ProblemsReponse {
    problems: Problem[];
}

export interface ProblemReponse {
    problem: Problem;
}

export interface UsersResponse {
    users: User[];
}

export interface CategoriesReponse {
    categories: Category[];
}