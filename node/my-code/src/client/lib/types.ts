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

export interface ProblemsSearchFilter {
    text?: string;
    categoryIds?: string[];
    favorites?: boolean;
}

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

export interface SearchProblemsResponse {
    searchProblems: {
        edges: Problem[];
        pageInfo: {
            cursor?: string;
            hasNextPage: boolean;
        }
    }
}

export interface Message {
    id: string;
    type: 'error' | 'success',
    text: string;
}