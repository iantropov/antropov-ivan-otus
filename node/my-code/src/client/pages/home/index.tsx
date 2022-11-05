import React from 'react';
import { NextPage } from 'next';
import classnames from 'classnames';

import { useUser } from '../../hooks/use-user';
import { Main } from '../../components/Main';

import styles from './styles.module.scss';

const Home: NextPage = () => {
    const [user, isUserLoading] = useUser();

    if (isUserLoading) return <p>Loading...</p>;
    if (!user) return null;

    return (
        <Main className={classnames(styles.home)}>
            <h1 className={classnames(styles.home__header)}>Hello, {user.name}!</h1>
        </Main>
    );
};

export default Home;
