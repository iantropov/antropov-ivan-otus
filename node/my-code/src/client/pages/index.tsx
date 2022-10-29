import React from 'react';
import { NextPage } from 'next';
import { useQuery } from '@apollo/client';

import { GET_PROBLEMS_QUERY } from '../lib/graphql';
import { useUser } from '../hooks/use-user';
import { ProblemsReponse } from '../lib/types';
import { Problems } from '../components/Problems';

const Home: NextPage = () => {
    const { data, loading } = useQuery<ProblemsReponse>(GET_PROBLEMS_QUERY);
    const [user, isUserLoading] = useUser();

    if (loading || isUserLoading) return <p>Loading...</p>;
    if (!user) return null;

    return (
        <section className="my-content">
            <div className="my-header">
                <h1>Hello, {user.name}!</h1>
            </div>
            <Problems className="my-problems" problems={data.problems} />
        </section>
    );
};

export default Home;
