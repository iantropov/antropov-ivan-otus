import React from 'react';
import { NextPage } from 'next';
import { useQuery } from '@apollo/client';
import classnames from 'classnames';

import { GET_PROBLEMS_QUERY } from '../../lib/graphql';
import { useUser } from '../../hooks/use-user';
import { ProblemsReponse } from '../../lib/types';
import { Problems } from '../../components/Problems';

import styles from './favorites.module.scss';

const Favorites: NextPage = () => {
    const { data, loading } = useQuery<ProblemsReponse>(GET_PROBLEMS_QUERY);
    const [user, isUserLoading] = useUser();

    if (loading || isUserLoading) return <p>Loading...</p>;
    if (!user) return null;

    return (
        <section className={classnames(styles.favorites)}>
            <h2 className={classnames(styles.favorites__header)}>You have some problems:</h2>
            <Problems
                className={classnames(styles.favorites__problems)}
                problems={data.problems}
                allowEdit={false}
            />
        </section>
    );
};

export default Favorites;
