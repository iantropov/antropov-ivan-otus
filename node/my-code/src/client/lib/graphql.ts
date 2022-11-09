import { gql } from '@apollo/client';

export const GET_USERS_QUERY = gql`
    query {
        users {
            _id
            email
            name
            isAdmin
        }
    }
`;

export const WHO_AM_I_QUERY = gql`
    query WHO_AM_I {
        whoAmI {
            _id
            name
            email
            isAdmin
            favorites
        }
    }
`;

export const GET_PROBLEMS_QUERY = gql`
    query {
        problems {
            summary
            description
            solution
            categories {
                _id
                name
            }
            _id
        }
    }
`;

export const GET_PROBLEM_QUERY = gql`
    query ($id: ID!) {
        problem(id: $id) {
            summary
            description
            solution
            categoryIds
            _id
        }
    }
`;

export const GET_CATEGORIES_QUERY = gql`
    query {
        categories {
            name
            _id
        }
    }
`;

export const SEARCH_PROBLEMS_QUERY = gql`
    query (
        $text: String
        $categoryIds: [String!]
        $favorites: Boolean
        $cursor: String
        $limit: Float
    ) {
        searchProblems(
            text: $text
            categoryIds: $categoryIds
            favorites: $favorites
            cursor: $cursor
            limit: $limit
        ) {
            edges {
                summary
                description
                solution
                categories {
                    _id
                    name
                }
                _id
            }
            pageInfo {
                hasNextPage
                cursor
            }
        }
    }
`;

export const GET_FAVORITE_PROBLEMS_QUERY = gql`
    query GetFavoriteProblems($cursor: String, $limit: Float) {
        searchProblems(favorites: true, cursor: $cursor, limit: $limit) {
            edges {
                summary
                description
                solution
                categories {
                    _id
                    name
                }
                _id
            }
            pageInfo {
                hasNextPage
                cursor
            }
        }
    }
`;

export const LOGOUT_USER_MUTATION = gql`
    mutation {
        logoutUser
    }
`;

export const LOGIN_USER_MUTATION = gql`
    mutation loginUser($email: String!, $password: String!) {
        loginUser(email: $email, password: $password) {
            accessToken
        }
    }
`;

export const REGISTER_USER_MUTATION = gql`
    mutation registerUser($name: String!, $email: String!, $password: String!) {
        registerUser(registerUserInput: { name: $name, email: $email, password: $password }) {
            _id
            name
            email
        }
    }
`;

export const DELETE_USER_MUTATION = gql`
    mutation deleteUser($userId: ID!) {
        deleteUser(id: $userId) {
            _id
        }
    }
`;

export const CREATE_PROBLEM_MUTATION = gql`
    mutation createProblem(
        $summary: String!
        $description: String!
        $solution: String
        $categoryIds: [String!]!
    ) {
        createProblem(
            createProblemInput: {
                summary: $summary
                description: $description
                solution: $solution
                categoryIds: $categoryIds
            }
        ) {
            summary
            description
            solution
            categories {
                _id
                name
            }
            categoryIds
            _id
        }
    }
`;

export const UPDATE_PROBLEM_MUTATION = gql`
    mutation updateProblem(
        $id: ID!
        $summary: String!
        $description: String!
        $solution: String
        $categoryIds: [String!]!
    ) {
        updateProblem(
            id: $id
            updateProblemInput: {
                summary: $summary
                description: $description
                solution: $solution
                categoryIds: $categoryIds
            }
        ) {
            summary
            description
            solution
            categories {
                _id
                name
            }
            categoryIds
            _id
        }
    }
`;

export const DELETE_PROBLEM_MUTATION = gql`
    mutation deleteProblem($problemId: ID!) {
        deleteProblem(id: $problemId) {
            _id
        }
    }
`;

export const LIKE_PROBLEM_MUTATION = gql`
    mutation likeProblem($problemId: ID!) {
        likeProblem(id: $problemId) {
            _id
        }
    }
`;

export const UNLIKE_PROBLEM_MUTATION = gql`
    mutation unlikeProblem($problemId: ID!) {
        unlikeProblem(id: $problemId) {
            _id
        }
    }
`;

export const CREATE_CATEGORY_MUTATION = gql`
    mutation createCategory($name: String!) {
        createCategory(createCategoryInput: { name: $name }) {
            name
            _id
        }
    }
`;
