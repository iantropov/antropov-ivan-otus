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
