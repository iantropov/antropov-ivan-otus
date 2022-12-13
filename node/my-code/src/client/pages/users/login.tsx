import React, { useState } from 'react';
import { NextPage } from 'next';
import { useMutation } from '@apollo/client';
import { useRouter } from 'next/router';
import classnames from 'classnames';

import { LOGIN_USER_MUTATION, WHO_AM_I_QUERY } from '../../lib/graphql';
import { Main } from '../../components/Main';
import { messageBroker } from '../../lib/message-broker';

import styles from './login.module.scss';

const Login: NextPage = () => {
    const router = useRouter();

    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const [isUserLoggingIn, setIsUserLoggingIn] = useState(false);

    const [loginUser, { client }] = useMutation(LOGIN_USER_MUTATION, {
        refetchQueries: [{ query: WHO_AM_I_QUERY }]
    });

    const onEmailChange = event => {
        setEmail(event.currentTarget.value);
    };

    const onPasswordChange = event => {
        setPassword(event.currentTarget.value);
    };

    const onSubmit = event => {
        event.preventDefault();
        setIsUserLoggingIn(true);

        loginUser({ variables: { email, password } }).then(
            () => {
                client.resetStore();
                router.push('/');
                messageBroker.addSuccessMessage('Logged in successfully!');
            },
            () => {
                setIsUserLoggingIn(false);
            }
        );
    };

    return (
        <Main className={styles.login}>
            <div className={styles.login__header}>
                <h1>Login page</h1>
            </div>
            <div className={classnames(styles.login__body, styles.loginBody)}>
                <form onSubmit={onSubmit}>
                    <div className={classnames(styles.loginBody__row, styles.loginRow)}>
                        <label className={classnames(styles.loginRow__label, 'form-label')}>
                            Email:
                            <input
                                className="form-control"
                                name="email"
                                type="email"
                                value={email}
                                onChange={onEmailChange}
                            />
                        </label>
                    </div>
                    <div className={classnames(styles.loginBody__row, styles.loginRow)}>
                        <label className={classnames(styles.loginRow__label, 'form-label')}>
                            Password:
                            <input
                                name="password"
                                type="password"
                                className="form-control"
                                value={password}
                                onChange={onPasswordChange}
                            />
                        </label>
                    </div>
                    <div className={styles.loginBody__row}>
                        <button
                            type="submit"
                            className={classnames(styles.loginRow__button, 'btn', 'btn-primary')}
                            disabled={isUserLoggingIn || !email || !password}
                        >
                            {isUserLoggingIn ? 'Logging in ...' : 'Login'}
                        </button>
                    </div>
                </form>
            </div>
        </Main>
    );
};

export default Login;
