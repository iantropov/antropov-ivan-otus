import React, { useState } from 'react';
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
import { Problem, ProblemsSearchFilter, SearchProblemsResponse } from '../../lib/types';
import { Problems as ProblemsComponent } from '../../components/Problems';
import { ProblemsSearchFilter as ProblemsSearchFilterComponent } from '../../components/ProblemsSearchFilter';
import { Main } from '../../components/Main';

import styles from './all.module.scss';

const PROBLEMS_LIMIT = 1;

const Problems: NextPage = () => {
    const { data, loading, refetch, fetchMore } = useQuery<SearchProblemsResponse>(
        SEARCH_PROBLEMS_QUERY,
        {
            fetchPolicy: 'network-only',
            variables: {
                limit: PROBLEMS_LIMIT
            }
        }
    );
    const [user, isUserLoading] = useUser();
    const [filter, setFilter] = useState({});

    const [deleteProblem] = useMutation(DELETE_PROBLEM_MUTATION, {
        refetchQueries: [SEARCH_PROBLEMS_QUERY]
    });

    const [likeProblem] = useMutation(LIKE_PROBLEM_MUTATION, {
        refetchQueries: [{ query: WHO_AM_I_QUERY }]
    });
    const [unlikeProblem] = useMutation(UNLIKE_PROBLEM_MUTATION, {
        refetchQueries: [{ query: WHO_AM_I_QUERY }]
    });

    const onApplyFilter = (filter: ProblemsSearchFilter) => {
        setFilter(filter);
        refetch({
            ...filter
        });
    };

    const onFetchMore = () => {
        fetchMore({
            variables: {
                ...filter,
                cursor: data.searchProblems.pageInfo.cursor
            }
        });
    };

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
                onApply={onApplyFilter}
            />
            {user.isAdmin && (
                <Link href="/problems/new">
                    <a className={classnames(styles.problems__createProblem, 'btn btn-primary')}>
                        Create Problem
                    </a>
                </Link>
            )}
            <ProblemsComponent
                className={classnames(styles.problems__problems)}
                problems={data.searchProblems.edges}
                favorites={user.favorites}
                allowEdit={user.isAdmin}
                allowDelete={user.isAdmin}
                hasNextPage={data.searchProblems.pageInfo.hasNextPage}
                onDelete={onDelete}
                onLike={onLike}
                onUnlike={onUnlike}
                onFetchMore={onFetchMore}
            />
        </Main>
    );
};

export default Problems;
