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
    const [loginUser, { data, loading, error }] = useMutation(LOGIN_USER);

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
        <section className="my-content">
            <div className="my-header">
                <h1>Login page</h1>
            </div>
            <div className="my-body">
                <form onSubmit={onSubmit}>
                    <label>
                        Email:
                        <input name="email" type="email" value={email} onChange={onEmailChange} />
                    </label>
                    <label>
                        Password:
                        <input
                            name="password"
                            type="password"
                            value={password}
                            onChange={onPasswordChange}
                        />
                    </label>
                    <input type="submit" value="Login" />
                </form>
                <h2>Current data:</h2>
                {JSON.stringify(data)}
            </div>
        </section>
    );
};

export default Login;
