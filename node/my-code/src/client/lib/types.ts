export interface User {
    _id: string;
    email: string;
    name: string;
}

export interface Problem {
    _id: string;
    summary: string;
    description: string;
    solution: string;
}

export interface WhoAmIResponse {
    whoAmI: User;
}

export interface ProblemsReponse {
    problems: Problem[]
}
