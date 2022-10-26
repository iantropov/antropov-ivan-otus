import React, { useState } from 'react';
import { NextPage } from 'next';
import { gql, useMutation } from '@apollo/client';
import { useRouter } from 'next/router';

const LOGIN_USER = gql`
    mutation loginUser($email: String!, $password: String!) {
        loginUser(email: $email, password: $password) {
            accessToken
        }
    }
`;

const Login: NextPage = () => {
    const router = useRouter();

    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const [loginUser, { loading }] = useMutation(LOGIN_USER);

    const onEmailChange = event => {
        setEmail(event.currentTarget.value);
    };

    const onPasswordChange = event => {
        setPassword(event.currentTarget.value);
    };

    const onSubmit = event => {
        event.preventDefault();

        loginUser({ variables: { email, password } }).then(
            result => {
                console.log(result);
                router.push('/');
            },
            error => {
                console.log(error);
                alert(error);
            }
        );
        alert(JSON.stringify([...new FormData(event.currentTarget).entries()]));
    };

    if (loading) return <p>Loging in...</p>;

    return (
        <section className="login">
            <div className="login__header">
                <h1>Login page</h1>
            </div>
            <div className="login__body login-body">
                <form onSubmit={onSubmit}>
                    <div className="login-body__row login-row">
                        <label className="login-row__label form-label">
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
                    <div className="login-body__row login-row">
                        <label className="login-row__label form-label">
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
                    <div className="login-body__row">
                        <button type="submit" className="login-row__button btn btn-primary">
                            Login
                        </button>
                    </div>
                </form>
            </div>
        </section>
    );
};

export default Login;
