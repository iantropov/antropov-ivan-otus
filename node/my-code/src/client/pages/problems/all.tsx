import React from 'react';
import { NextPage } from 'next';
import { useQuery } from '@apollo/client';
import classnames from 'classnames';
import Link from 'next/link';

import { SEARCH_PROBLEMS_QUERY } from '../../lib/graphql';
import { useUser } from '../../hooks/use-user';
import { SearchProblemsResponse } from '../../lib/types';
import { Problems as ProblemsComponent } from '../../components/Problems';
import { ProblemsSearchFilter as ProblemsSearchFilterComponent } from '../../components/ProblemsSearchFilter';

import styles from './all.module.scss';

const Problems: NextPage = () => {
    const { data, loading, refetch } = useQuery<SearchProblemsResponse>(SEARCH_PROBLEMS_QUERY);
    const [user, isUserLoading] = useUser();

    if (loading || isUserLoading) return <p>Loading...</p>;
    if (!user) return null;

    return (
        <main className={classnames(styles.problems)}>
            <ProblemsSearchFilterComponent
                className={classnames(styles.problems__problemsFilterSearch)}
                onApply={refetch}
            />
            <ProblemsComponent
                className={classnames(styles.problems__problems)}
                problems={data.searchProblems}
                allowEdit={user.isAdmin}
                allowRemove={user.isAdmin}
            />
            {user.isAdmin && (
                <Link href="/problems/new">
                    <a className="btn btn-primary">Create Problem</a>
                </Link>
            )}
        </main>
    );
};

export default Problems;
