import { gql } from '@apollo/client';

export const ALL_USERS_QUERY = gql`
    query {
        users {
            _id
            email
            name
        }
    }
`;

export const WHO_AM_I_QUERY = gql`
    query WHO_AM_I {
        whoAmI {
            _id
            name
            email
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
        registerUser(registerUserInput: {
            name: $name,
            email: $email,
            password: $password
        }) {
            _id,
            name,
            email
        }
    }
`;
