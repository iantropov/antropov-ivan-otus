import React, { useState } from 'react';
import { NextPage } from 'next';
import { useMutation } from '@apollo/client';
import { useRouter } from 'next/router';
import classnames from 'classnames';

import { REGISTER_USER_MUTATION } from '../../lib/graphql';
import { Main } from '../../components/Main';

import styles from './register.module.scss';
import { messageBroker } from '../../lib/message-broker';

const Register: NextPage = () => {
    const router = useRouter();

    const [name, setName] = useState('');
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const [passwordRetype, setPasswordRetype] = useState('');
    const [isUserRegistering, setIsUserRegistering] = useState(false);

    const [registerUser] = useMutation(REGISTER_USER_MUTATION);

    const onNameChange = event => {
        setName(event.currentTarget.value);
    };

    const onEmailChange = event => {
        setEmail(event.currentTarget.value);
    };

    const onPasswordChange = event => {
        setPassword(event.currentTarget.value);
    };

    const onPasswordRetypeChange = event => {
        setPasswordRetype(event.currentTarget.value);
    };

    const onSubmit = event => {
        event.preventDefault();

        setIsUserRegistering(true);

        registerUser({ variables: { name, email, password } }).then(
            () => {
                router.push('/users/login');
                messageBroker.addSuccessMessage('Registered the user successfully!');
            },
            () => {
                setIsUserRegistering(false);
            }
        );
    };

    return (
        <Main className={styles.register}>
            <div className={styles.register__header}>
                <h1>Register page</h1>
            </div>
            <div className={classnames(styles.register__body, styles.registerBody)}>
                <form onSubmit={onSubmit}>
                    <div className={classnames(styles.registerBody__row, styles.registerRow)}>
                        <label className={classnames(styles.registerRow__label, 'form-label')}>
                            Name:
                            <input
                                className="form-control"
                                name="name"
                                type="text"
                                value={name}
                                onChange={onNameChange}
                            />
                        </label>
                    </div>
                    <div className={classnames(styles.registerBody__row, styles.registerRow)}>
                        <label className={classnames(styles.registerRow__label, 'form-label')}>
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
                    <div className={classnames(styles.registerBody__row, styles.registerRow)}>
                        <label className={classnames(styles.registerRow__label, 'form-label')}>
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
                    <div className={classnames(styles.registerBody__row, styles.registerRow)}>
                        <label className={classnames(styles.registerRow__label, 'form-label')}>
                            Re-type the Password:
                            <input
                                name="password-retype"
                                type="password"
                                className="form-control"
                                value={passwordRetype}
                                onChange={onPasswordRetypeChange}
                            />
                        </label>
                    </div>
                    <div className={styles.registerBody__row}>
                        <button
                            type="submit"
                            className={classnames(styles.registerRow__button, 'btn', 'btn-primary')}
                            disabled={
                                isUserRegistering ||
                                !name ||
                                !email ||
                                !password ||
                                password !== passwordRetype
                            }
                        >
                            {isUserRegistering ? 'Registering...' : 'Register'}
                        </button>
                    </div>
                </form>
            </div>
        </Main>
    );
};

export default Register;
