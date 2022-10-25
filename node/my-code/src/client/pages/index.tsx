import React from 'react';
import { NextPage } from 'next';
import Link from 'next/link';

import { gql, useQuery } from '@apollo/client'

const AllUsersQuery = gql`
  query {
    users {
      _id
      email
      name
    }
  }
`

const Home: NextPage = () => {
    const { data, loading, error } = useQuery(AllUsersQuery)

    if (loading) return <p>Loading...</p>
    if (error) return <p>Oh no... {error.message}</p>

    return (
        <section className="my-content">
            <div className="my-header">
                <h1>Hello, World 22!</h1>
            </div>
            <div className="my-links">
                <Link href="/login">Login</Link>
                <br/>
                <Link href="/register">Register</Link>
            </div>
            <div className="my-users">
                <ul className="my-users-list">
                    {data.users.map((user) => (
                        <li key={user._id}>
                            _id: {user._id}, email: {user.email}, name: {user.name}
                        </li>
                    ))}
                </ul>
            </div>
        </section>
    );
};

export default Home;
