import React from 'react';
import { NextPage } from 'next';
import Link from 'next/link';
import { useRouter } from 'next/router';
import { gql, useMutation, useQuery } from '@apollo/client';
import { ALL_USERS_QUERY, LOGOUT_USER_MUTATION, WHO_AM_I_QUERY } from '../lib/graphql-queries';

const Home: NextPage = () => {
    const router = useRouter();
    const { data, loading, error } = useQuery(ALL_USERS_QUERY);
    const { data: userData, loading: userLoading, error: userError } = useQuery(WHO_AM_I_QUERY);

    if (loading || userLoading) return <p>Loading...</p>;
    if (error) return <p>Oh no... {error.message}</p>;
    if (userError) return <p>Oh no... {userError.message}</p>;

    if (!userData?.whoAmI) {
        // router.replace('/login');
        return null;
    }

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
                        _id: {userData.whoAmI?._id}, email: {userData.whoAmI?.email}, name:{' '}
                        {userData.whoAmI?.name}
                    </p>
                )}
            </div>
        </section>
    );
};

export default Home;
