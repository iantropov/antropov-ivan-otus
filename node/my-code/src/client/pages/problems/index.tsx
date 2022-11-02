import React from 'react';
import { NextPage } from 'next';
import { useQuery } from '@apollo/client';
import classnames from 'classnames';

import { GET_PROBLEMS_QUERY } from '../../lib/graphql';
import { useUser } from '../../hooks/use-user';
import { ProblemsReponse } from '../../lib/types';
import { Problems as ProblemsComponent } from '../../components/Problems';

import styles from './styles.module.scss';

const Problems: NextPage = () => {
    const { data, loading } = useQuery<ProblemsReponse>(GET_PROBLEMS_QUERY);
    const [user, isUserLoading] = useUser();

    if (loading || isUserLoading) return <p>Loading...</p>;
    if (!user) return null;

    return (
        <section className={classnames(styles.problems)}>
            <h2 className={classnames(styles.problems__header)}>You have some problems:</h2>
            <ProblemsComponent
                className={classnames(styles.problems__problems)}
                problems={data.problems}
            />
        </section>
    );
};

export default Problems;
