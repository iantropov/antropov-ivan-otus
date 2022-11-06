import React from 'react';
import { NextPage } from 'next';
import { useMutation, useQuery } from '@apollo/client';
import classnames from 'classnames';

import {
    GET_FAVORITE_PROBLEMS_QUERY,
    LIKE_PROBLEM_MUTATION,
    UNLIKE_PROBLEM_MUTATION,
    WHO_AM_I_QUERY
} from '../../lib/graphql';
import { useUser } from '../../hooks/use-user';
import { Problem, SearchProblemsResponse } from '../../lib/types';
import { Problems } from '../../components/Problems';
import { Main } from '../../components/Main';

import styles from './favorites.module.scss';

const Favorites: NextPage = () => {
    const { data, loading, fetchMore } = useQuery<SearchProblemsResponse>(GET_FAVORITE_PROBLEMS_QUERY, {
        fetchPolicy: 'network-only'
    });
    const [user, isUserLoading] = useUser();

    const [likeProblem] = useMutation(LIKE_PROBLEM_MUTATION, {
        refetchQueries: [{ query: WHO_AM_I_QUERY }, GET_FAVORITE_PROBLEMS_QUERY]
    });
    const [unlikeProblem] = useMutation(UNLIKE_PROBLEM_MUTATION, {
        refetchQueries: [{ query: WHO_AM_I_QUERY }, GET_FAVORITE_PROBLEMS_QUERY]
    });

    const onFetchMore = () => {
        fetchMore({
            variables: {
                cursor: data.searchProblems.pageInfo.cursor
            }
        });
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
        <Main className={classnames(styles.favorites)}>
            <h2 className={classnames(styles.favorites__header)}>You have favorite problems:</h2>
            <Problems
                className={classnames(styles.favorites__problems)}
                problems={data.searchProblems.edges}
                favorites={user.favorites}
                hasNextPage={data.searchProblems.pageInfo.hasNextPage}
                allowEdit={false}
                allowDelete={false}
                onFetchMore={onFetchMore}
                onLike={onLike}
                onUnlike={onUnlike}
            />
        </Main>
    );
};

export default Favorites;
