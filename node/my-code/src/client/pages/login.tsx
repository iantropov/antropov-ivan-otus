import React, { useState } from 'react';
import { NextPage } from 'next';

const Login: NextPage = () => {
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');

    const onEmailChange = (event) => {
        setEmail(event.value);
    }

    const onPasswordChange = (event) => {
        setPassword(event.value);
    }

    const onSubmit = (event) => {
        event.preventDefault();
        alert(JSON.stringify([...(new FormData(event.currentTarget)).entries()]));
    }

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
                        <input name="password" type="password" value={password} onChange={onPasswordChange} />
                    </label>
                    <input type="submit" value="Login"/>
                </form>
            </div>
        </section>
    );
};

export default Login;
