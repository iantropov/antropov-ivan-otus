import React from 'react';
import { NextPage } from 'next';
import Link from 'next/link';
import { useRouter } from 'next/router';
import { gql, useMutation, useQuery } from '@apollo/client';

const AllUsersQuery = gql`
    query {
        users {
            _id
            email
            name
        }
    }
`;

const WhoAmIQuery = gql`
    query {
        whoAmI {
            _id
            name
            email
        }
    }
`;

const LOGOUT_USER = gql`
    mutation {
        logoutUser
    }
`;


const Home: NextPage = () => {
    const router = useRouter();
    const { data, loading, error } = useQuery(AllUsersQuery);
    const { data: userData, loading: userLoading, error: userError } = useQuery(WhoAmIQuery);
    const [logoutUser, { data: logoutData, loading: logoutLoading, error: logoutError }] = useMutation(LOGOUT_USER);

    const onLogoutUserClick = () => {
        logoutUser().then(() => {
            alert("Logged Out!");
            router.push('/login');
        }, (error) => {
            alert(error);
        });
    }

    if (loading || userLoading) return <p>Loading...</p>;
    if (error) return <p>Oh no... {userError.message}</p>;

    return (
        <section className="my-content">
            <div className="my-header">
                <h1>Hello, World 22!</h1>
            </div>
            <div className="my-links">
                <Link href="/login">Login</Link>
                <br />
                <Link href="/register">Register</Link>
            </div>
            <div className="my-users">
                <ul className="my-users-list">
                    {data.users.map(user => (
                        <li key={user._id}>
                            _id: {user._id}, email: {user.email}, name: {user.name}
                        </li>
                    ))}
                </ul>
            </div>
            <div className="my-user">
                <h2>Who Am I:</h2>
                {userError ? (
                    'No Data!'
                ) : (
                    <p>
                        _id: {userData.whoAmI._id}, email: {userData.whoAmI.email}, name: {userData.whoAmI.name}
                        <button type="button" onClick={onLogoutUserClick}>Log Out</button>
                    </p>
                )}
            </div>
        </section>
    );
};

export default Home;
