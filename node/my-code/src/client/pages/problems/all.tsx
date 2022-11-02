import React from 'react';
import { NextPage } from 'next';
import { useQuery } from '@apollo/client';
import classnames from 'classnames';
import Link from 'next/link';

import { GET_PROBLEMS_QUERY } from '../../lib/graphql';
import { useUser } from '../../hooks/use-user';
import { ProblemsReponse } from '../../lib/types';
import { Problems as ProblemsComponent } from '../../components/Problems';

import styles from './all.module.scss';

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
                allowEdit={user.isAdmin}
                allowRemove={user.isAdmin}
            />
            {user.isAdmin && (
                <Link href="/problems/new">
                    <a className="btn btn-primary">Create Problem</a>
                </Link>
            )}
        </section>
    );
};

export default Problems;
