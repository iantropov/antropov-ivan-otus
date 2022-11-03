export interface User {
    _id: string;
    email: string;
    name: string;
    isAdmin: boolean;
    favorites: string[];
}

export interface Problem {
    _id: string;
    summary: string;
    description: string;
    solution: string;
}

export type ProblemData = Omit<Problem, '_id'>;

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
