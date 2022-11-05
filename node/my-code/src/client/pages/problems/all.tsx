import React from 'react';
import { NextPage } from 'next';
import { useMutation, useQuery } from '@apollo/client';
import classnames from 'classnames';
import Link from 'next/link';

import {
    DELETE_PROBLEM_MUTATION,
    LIKE_PROBLEM_MUTATION,
    SEARCH_PROBLEMS_QUERY,
    UNLIKE_PROBLEM_MUTATION,
    WHO_AM_I_QUERY
} from '../../lib/graphql';
import { useUser } from '../../hooks/use-user';
import { Problem, SearchProblemsResponse } from '../../lib/types';
import { Problems as ProblemsComponent } from '../../components/Problems';
import { ProblemsSearchFilter as ProblemsSearchFilterComponent } from '../../components/ProblemsSearchFilter';
import { Main } from '../../components/Main';

import styles from './all.module.scss';

const Problems: NextPage = () => {
    const { data, loading, refetch } = useQuery<SearchProblemsResponse>(SEARCH_PROBLEMS_QUERY, {
        fetchPolicy: 'network-only'
    });
    const [user, isUserLoading] = useUser();

    const [deleteProblem] = useMutation(DELETE_PROBLEM_MUTATION, {
        refetchQueries: [SEARCH_PROBLEMS_QUERY]
    });

    const [likeProblem] = useMutation(LIKE_PROBLEM_MUTATION, {
        refetchQueries: [{ query: WHO_AM_I_QUERY }]
    });
    const [unlikeProblem] = useMutation(UNLIKE_PROBLEM_MUTATION, {
        refetchQueries: [{ query: WHO_AM_I_QUERY }]
    });

    const onDelete = (problem: Problem) => {
        return deleteProblem({ variables: { problemId: problem._id } }).then(
            () => {
                console.log(`Removed problem #${problem._id} successfully!`);
            },
            error => {
                alert(error);
            }
        );
    };

    const onLike = (problem: Problem) => {
        return likeProblem({ variables: { problemId: problem._id } }).then(
            () => {
                console.log('Liked problem #${problem._id} successfully!');
            },
            error => {
                alert(error);
            }
        );
    };

    const onUnlike = (problem: Problem) => {
        return unlikeProblem({ variables: { problemId: problem._id } }).then(
            () => {
                console.log('Unliked problem #${problem._id} successfully!');
            },
            error => {
                alert(error);
            }
        );
    };

    if (loading || isUserLoading) return <p>Loading...</p>;
    if (!user) return null;

    return (
        <Main className={classnames(styles.problems)} align="top">
            <ProblemsSearchFilterComponent
                className={classnames(styles.problems__problemsFilterSearch)}
                onApply={refetch}
            />
            <ProblemsComponent
                className={classnames(styles.problems__problems)}
                problems={data.searchProblems}
                favorites={user.favorites}
                allowEdit={user.isAdmin}
                allowDelete={user.isAdmin}
                onDelete={onDelete}
                onLike={onLike}
                onUnlike={onUnlike}
            />
            {user.isAdmin && (
                <Link href="/problems/new">
                    <a className="btn btn-primary">Create Problem</a>
                </Link>
            )}
        </Main>
    );
};

export default Problems;
