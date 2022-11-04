import React from 'react';
import { NextPage } from 'next';
import classnames from 'classnames';

import { useUser } from '../../hooks/use-user';

import styles from './styles.module.scss';

const Home: NextPage = () => {
    const [user, isUserLoading] = useUser();

    if (isUserLoading) return <p>Loading...</p>;
    if (!user) return null;

    return (
        <main className={classnames(styles.home)}>
            <h1 className={classnames(styles.home__header)}>Hello, {user.name}!</h1>
        </main>
    );
};

export default Home;
