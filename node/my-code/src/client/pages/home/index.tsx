import React from 'react';
import { NextPage } from 'next';
import { useQuery } from '@apollo/client';
import classnames from 'classnames';

import { GET_PROBLEMS_QUERY } from '../../lib/graphql';
import { useUser } from '../../hooks/use-user';
import { ProblemsReponse } from '../../lib/types';

import styles from './styles.module.scss';

const Home: NextPage = () => {
    const { data, loading } = useQuery<ProblemsReponse>(GET_PROBLEMS_QUERY);
    const [user, isUserLoading] = useUser();

    if (loading || isUserLoading) return <p>Loading...</p>;
    if (!user) return null;

    return (
        <section className="my-content">
            <h1 className={classnames(styles.home__header)}>Hello, {user.name}!</h1>
        </section>
    );
};

export default Home;
