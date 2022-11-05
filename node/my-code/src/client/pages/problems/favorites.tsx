import React from 'react';
import { NextPage } from 'next';
import { useQuery } from '@apollo/client';
import classnames from 'classnames';

import { GET_FAVORITE_PROBLEMS_QUERY, SEARCH_PROBLEMS_QUERY } from '../../lib/graphql';
import { useUser } from '../../hooks/use-user';
import { SearchProblemsResponse } from '../../lib/types';
import { Problems } from '../../components/Problems';
import { Main } from '../../components/Main';

import styles from './favorites.module.scss';

const Favorites: NextPage = () => {
    const { data, loading } = useQuery<SearchProblemsResponse>(GET_FAVORITE_PROBLEMS_QUERY);
    const [user, isUserLoading] = useUser();

    if (loading || isUserLoading) return <p>Loading...</p>;
    if (!user) return null;

    return (
        <Main className={classnames(styles.favorites)}>
            <h2 className={classnames(styles.favorites__header)}>You have favorite problems:</h2>
            <Problems
                className={classnames(styles.favorites__problems)}
                problems={data.searchProblems}
                allowEdit={false}
                allowRemove={false}
            />
        </Main>
    );
};

export default Favorites;
